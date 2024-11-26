package main

import (
	"AmanahPro/services/bank-services/bootstrap"
	"AmanahPro/services/bank-services/config"
	"AmanahPro/services/bank-services/factories"
	"AmanahPro/services/bank-services/internal/application/services"
	"AmanahPro/services/bank-services/internal/domain/repositories"
	"AmanahPro/services/bank-services/internal/handlers"
	"AmanahPro/services/bank-services/routes"
	"log"
	"os"
	"time"

	_ "AmanahPro/services/bank-services/docs"

	"github.com/NHadi/AmanahPro-common/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Bank Services API
// @version 1.0
// @description This is the Bank Services API documentation for managing transactions, uploads, and reconciliations.
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Provide your JWT token with "Bearer " prefix, e.g., "Bearer <token>"
// @host localhost:8082
// @BasePath /
func main() {

	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("Application panic recovered: %v", r)
		}
	}()

	// Configure daily log file
	configureDailyLogFile()

	// Log startup message
	log.Println("Starting Bank Services API...")

	// Load configuration
	cfg := loadConfig()

	// Initialize dependencies
	deps := initializeDependencies(cfg)

	// Declare RabbitMQ queue
	initializeRabbitMQ(deps)

	// Initialize services and handlers
	repos := factories.CreateRepositories(deps.DB, deps.ElasticsearchClient)
	services := factories.CreateServices(repos, deps.RabbitMQPublisher, deps.RabbitMQConsumer, deps.ElasticsearchClient, deps.RedisClient)
	handlerInstances := initializeHandlers(services, repos)

	// Configure and start background tasks
	configureReconciliationScheduler(deps, services)
	startRabbitMQConsumer(services)

	// Start Gin server
	startServer(cfg, deps, handlerInstances)
}

// configureDailyLogFile sets up daily logging into a file.
func configureDailyLogFile() {
	// Ensure the log folder exists
	logDir := "log"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err := os.Mkdir(logDir, os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create log directory: %v", err)
		}
	}

	// Create log file with the current date
	logFileName := logDir + "/Log" + time.Now().Format("20060102") + ".log"
	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	// Set log output to the file and also log to stdout
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("Logging initialized to", logFileName)
}

// loadConfig loads the application configuration.
func loadConfig() *config.Config {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	return cfg
}

// initializeDependencies sets up all external dependencies like DB, RabbitMQ, Redis, Elasticsearch, etc.
func initializeDependencies(cfg *config.Config) *bootstrap.Dependencies {
	deps, err := bootstrap.InitDependencies(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize dependencies: %v", err)
	}
	return deps
}

// initializeRabbitMQ declares the RabbitMQ queue for the application.
func initializeRabbitMQ(deps *bootstrap.Dependencies) {
	const rabbitQueue = "transactions_queue"
	err := deps.RabbitMQService.DeclareQueue(rabbitQueue)
	if err != nil {
		log.Fatalf("Failed to declare RabbitMQ queue: %v", err)
	}
}

// initializeHandlers initializes all handlers used by the application.
func initializeHandlers(services *services.Services, repos *repositories.Repositories) *handlers.Handlers {
	return handlers.NewHandlers(
		handlers.NewUploadHandler(services.UploadService, repos.TransactionRepo, repos.BatchRepo),
		handlers.NewTransactionHandler(services.TransactionService),
		handlers.NewReconciliationHandler(services.ReconciliationService),
	)
}

// configureReconciliationScheduler sets up a periodic task for reconciliation.
func configureReconciliationScheduler(deps *bootstrap.Dependencies, services *services.Services) {
	_, err := deps.Scheduler.AddFunc("@every 10m", func() {
		log.Println("Starting periodic reconciliation...")
		if err := services.ReconciliationService.PerformReconciliation(); err != nil {
			log.Printf("Reconciliation failed: %v", err)
		} else {
			log.Println("Reconciliation completed successfully")
		}
	})
	if err != nil {
		log.Fatalf("Failed to configure reconciliation scheduler: %v", err)
	}

	// Start the scheduler
	go deps.Scheduler.Start()
}

// startRabbitMQConsumer starts the RabbitMQ consumer.
func startRabbitMQConsumer(services *services.Services) {
	go func() {
		err := services.ConsumerService.StartConsumer()
		if err != nil {
			log.Fatalf("Failed to start RabbitMQ consumer: %v", err)
		}
	}()
}

// startServer starts the Gin server and registers all routes.
func startServer(cfg *config.Config, deps *bootstrap.Dependencies, handlerInstances *handlers.Handlers) {
	router := gin.Default()

	// Apply middleware
	router.Use(deps.LoggerMiddleware)
	router.Use(func(c *gin.Context) {
		log.Printf("Incoming request: %s %s", c.Request.Method, c.Request.URL.Path)
		c.Next()
	})

	// Register API routes
	api := router.Group("/api")
	api.Use(middleware.JWTAuthMiddleware(cfg.JWTSecret))
	routes.RegisterAPIRoutes(api, handlerInstances)

	// Register Swagger routes
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start the server
	port := cfg.Port
	if port == "" {
		port = "8082"
	}
	log.Printf("Server running at http://localhost:%s", port)
	log.Fatal(router.Run(":" + port))
}
