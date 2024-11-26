package main

import (
	_ "AmanahPro/services/bank-services/docs" // Swagger docs
	"AmanahPro/services/bank-services/internal/application/services"
	domainRepositories "AmanahPro/services/bank-services/internal/domain/repositories"
	"AmanahPro/services/bank-services/internal/handlers"
	"AmanahPro/services/bank-services/internal/infrastructure/persistence"
	"AmanahPro/services/bank-services/internal/infrastructure/repositories"
	"log"
	"os"

	"github.com/NHadi/AmanahPro-common/messagebroker"
	"github.com/NHadi/AmanahPro-common/middleware"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	httpSwagger "github.com/swaggo/gin-swagger"
)

const defaultPort = "8082"

func main() {

	// Determine the runtime environment
	envFilePath := determineEnvFilePath("../../.env.local")

	err := godotenv.Load(envFilePath)
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Initialize DB
	db, err := persistence.InitializeDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Initialize RabbitMQ service
	rabbitMQURL := os.Getenv("RABBITMQ_URL")
	rabbitService, err := messagebroker.NewRabbitMQService(rabbitMQURL)
	if err != nil {
		log.Fatalf("Failed to initialize RabbitMQ service: %v", err)
	}
	defer rabbitService.Close()

	// Declare RabbitMQ queue
	rabbitQueue := "transactions_queue"
	err = rabbitService.DeclareQueue(rabbitQueue)
	if err != nil {
		log.Fatalf("Failed to declare RabbitMQ queue: %v", err)
	}

	// Initialize RabbitMQ publisher
	rabbitPublisher := messagebroker.NewRabbitMQPublisher(rabbitService)

	// Initialize Elasticsearch client
	esClient, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses:     []string{os.Getenv("ELASTICSEARCH_URL")},
		RetryOnStatus: []int{502, 503, 504},
		MaxRetries:    5,
	})
	if err != nil {
		log.Fatalf("Failed to initialize Elasticsearch client: %v", err)
	}

	// Check Elasticsearch connection
	res, err := esClient.Info()
	if err != nil || res.IsError() {
		log.Fatalf("Elasticsearch connection error: %v", err)
	}
	defer res.Body.Close()

	log.Println("Elasticsearch connection successful")

	// Initialize repositories
	var batchRepo domainRepositories.BatchRepository = repositories.NewBatchRepository(db)
	var transactionRepo domainRepositories.BankAccountTransactionRepository = repositories.NewBankAccountTransactionRepository(db, esClient, "bank-transactions")

	// Initialize application services
	uploadService := services.NewUploadService(transactionRepo, batchRepo, rabbitPublisher)
	transactionService := services.NewTransactionService(transactionRepo)

	// Initialize RabbitMQ consumer
	rabbitConsumer := messagebroker.NewRabbitMQConsumer(rabbitService)
	consumerService := services.NewConsumerService(esClient, "bank-transactions", rabbitConsumer, rabbitQueue)

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
	logger, err := middleware.InitializeLogger("bank-service", os.Getenv("ELASTICSEARCH_URL"), "bank-services-logs")
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

	// Start server
	port := os.Getenv("BANK_SERVICES_PORT")
	if port == "" {
		port = defaultPort
	}
	log.Printf("Server running at http://localhost:%s", port)
	log.Fatal(r.Run(":" + port))
}

// determineEnvFilePath determines the correct .env file path based on runtime environment
func determineEnvFilePath(localEnvPath string) string {
	// Check for Docker environment
	if isDockerEnvironment() {
		return "/app/.env" // Docker container path
	}

	if fileExists(localEnvPath) {
		return localEnvPath
	}

	// Fallback to production path
	return "/root/AmanahPro/.env" // Production path (e.g., VM)
}

// isDockerEnvironment checks if the application is running inside Docker
func isDockerEnvironment() bool {
	// Docker containers usually have a cgroup file
	if _, err := os.Stat("/.dockerenv"); err == nil {
		return true
	}
	// Alternatively, check for specific Docker files
	cgroupPath := "/proc/1/cgroup"
	if fileExists(cgroupPath) {
		return true
	}
	return false
}

// fileExists checks if a file or directory exists
func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
