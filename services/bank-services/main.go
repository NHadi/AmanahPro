package main

import (
	_ "AmanahPro/services/bank-services/docs" // Swagger docs
	"AmanahPro/services/bank-services/internal/application/services"
	domainRepositories "AmanahPro/services/bank-services/internal/domain/repositories"
	"AmanahPro/services/bank-services/internal/handlers"
	"AmanahPro/services/bank-services/internal/infrastructure/messagebroker"
	"AmanahPro/services/bank-services/internal/infrastructure/persistence"
	"AmanahPro/services/bank-services/internal/infrastructure/repositories"
	"log"
	"os"

	"github.com/NHadi/AmanahPro-common/middleware"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
	swaggerFiles "github.com/swaggo/files"
	httpSwagger "github.com/swaggo/gin-swagger"
)

const defaultPort = "8082"

// @title Bank Services API
// @version 1.0
// @description This is the Bank Services API for managing account transactions and upload batches.
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Provide your JWT token with "Bearer " prefix, e.g., "Bearer <token>"
// @host localhost:8082
// @BasePath /
func main() {
	// Check if running in Docker (using an environment variable)
	envFilePath := ".env" // Default path
	if _, isInDocker := os.LookupEnv("DOCKER_ENV"); isInDocker {
		envFilePath = "/app/.env" // Path for Docker container
	}

	// Load environment variables
	err := godotenv.Load(envFilePath)
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Initialize DB
	db, err := persistence.InitializeDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Initialize RabbitMQ connection
	rabbitConn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer rabbitConn.Close()

	rabbitCh, err := rabbitConn.Channel()
	if err != nil {
		log.Fatalf("Failed to create RabbitMQ channel: %v", err)
	}
	defer rabbitCh.Close()

	// Initialize RabbitMQ publisher
	rabbitQueue := "transactions_queue"
	rabbitPublisher, err := messagebroker.NewRabbitPublisher(rabbitConn, rabbitQueue)
	if err != nil {
		log.Fatalf("Failed to initialize RabbitMQ publisher: %v", err)
	}

	// Initialize Elasticsearch client
	esClient, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://elasticsearch:9200"},
	})
	if err != nil {
		log.Fatalf("Failed to initialize Elasticsearch client: %v", err)
	}
	// Initialize repositories
	var batchRepo domainRepositories.BatchRepository = repositories.NewBatchRepository(db)
	var transactionRepo domainRepositories.BankAccountTransactionRepository = repositories.NewBankAccountTransactionRepository(db, esClient, "bank-transactions")

	// Initialize application services
	uploadService := services.NewUploadService(transactionRepo, batchRepo, rabbitPublisher)
	transactionService := services.NewTransactionService(transactionRepo)
	consumerService := services.NewConsumerService(esClient, "bank-transactions", rabbitCh, rabbitQueue)

	// Initialize handlers
	uploadHandler := handlers.NewUploadHandler(uploadService, transactionRepo, batchRepo)
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	// Start RabbitMQ consumer
	go func() {
		err := consumerService.StartConsumer()
		if err != nil {
			log.Fatalf("Failed to start RabbitMQ consumer: %v", err)
		}
	}()

	// Initialize Gin router
	r := gin.Default()

	// Initialize common logger
	logger, err := middleware.InitializeLogger("bank-services", "http://elasticsearch:9200", "bank-services-logs")
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	// Attach common logging middleware
	r.Use(middleware.GinLoggingMiddleware(logger))

	// Middleware to log requests
	r.Use(func(c *gin.Context) {
		log.Printf("Incoming request: %s %s", c.Request.Method, c.Request.URL.Path)
		c.Next()
	})

	// Group for protected routes
	api := r.Group("/api")
	api.Use(middleware.JWTAuthMiddleware(os.Getenv("JWT_SECRET")))

	// Upload Routes
	api.POST("/upload", uploadHandler.UploadBatch)

	// Transaction Routes
	api.GET("/transactions", transactionHandler.GetTransactionsByBankAndPeriod)

	// Swagger documentation route
	r.GET("/swagger/*any", httpSwagger.WrapHandler(swaggerFiles.Handler))

	// Get port from environment or default to 8082
	port := os.Getenv("BANK_SERVICES_PORT")
	if port == "" {
		port = defaultPort
	}

	log.Printf("Server running at http://localhost:%s", port)
	log.Fatal(r.Run(":" + port))
}
