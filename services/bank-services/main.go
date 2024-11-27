package main

import (
	"AmanahPro/services/bank-services/common/bootstrap"
	"AmanahPro/services/bank-services/common/config"
	"AmanahPro/services/bank-services/common/factories"
	commonMiddleware "AmanahPro/services/bank-services/common/middleware"
	"AmanahPro/services/bank-services/common/routes"
	"AmanahPro/services/bank-services/internal/application/services"
	"AmanahPro/services/bank-services/internal/domain/repositories"
	"AmanahPro/services/bank-services/internal/handlers"
	"AmanahPro/services/bank-services/schedulers"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"time"

	_ "AmanahPro/services/bank-services/docs"

	"github.com/NHadi/AmanahPro-common/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	defaultPort       = "8082"
	rabbitMQQueueName = "transactions_queue"
	logDir            = "log"
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
	defer recoverFromPanic()

	configureLogging()
	log.Println("Starting Bank Services API...")

	cfg := loadConfig()
	deps := initializeDependencies(cfg)
	initializeRabbitMQ(deps)

	repos := factories.CreateRepositories(deps.DB, deps.ElasticsearchClient)
	services := factories.CreateServices(repos, deps.RabbitMQPublisher, deps.RabbitMQConsumer, deps.ElasticsearchClient, deps.RedisClient)
	handlers := initializeHandlers(services, repos)

	//scheduler setup
	schedulers.ConfigureReconciliationScheduler(deps.Scheduler, services.ReconciliationService.PerformReconciliation)

	startRabbitMQConsumer(services)

	router := setupRouter(cfg, deps, handlers)
	startServerWithGracefulShutdown(cfg, router)
}

// configureLogging sets up daily logging into a file.
func configureLogging() {
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		if err := os.Mkdir(logDir, os.ModePerm); err != nil {
			log.Fatalf("Failed to create log directory: %v", err)
		}
	}

	logFileName := fmt.Sprintf("%s/Log%s.log", logDir, time.Now().Format("20060102"))
	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Printf("Logging initialized: %s", logFileName)
}

// recoverFromPanic recovers from panics and logs the stack trace.
func recoverFromPanic() {
	if r := recover(); r != nil {
		log.Printf("Application panic recovered: %v", r)
		log.Printf("Stack trace: %s", debug.Stack())
	}
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
	err := deps.RabbitMQService.DeclareQueue(rabbitMQQueueName)
	if err != nil {
		log.Fatalf("Failed to declare RabbitMQ queue [%s]: %v", rabbitMQQueueName, err)
	}
	log.Printf("RabbitMQ queue [%s] declared successfully", rabbitMQQueueName)
}

// initializeHandlers initializes all handlers used by the application.
func initializeHandlers(services *services.Services, repos *repositories.Repositories) *handlers.Handlers {
	return handlers.NewHandlers(
		handlers.NewUploadHandler(services.UploadService, repos.TransactionRepo, repos.BatchRepo),
		handlers.NewTransactionHandler(services.TransactionService),
		handlers.NewReconciliationHandler(services.ReconciliationService),
	)
}

// startRabbitMQConsumer starts the RabbitMQ consumer.
func startRabbitMQConsumer(services *services.Services) {
	go func() {
		err := services.ConsumerService.StartConsumer()
		if err != nil {
			log.Fatalf("Failed to start RabbitMQ consumer: %v", err)
		}
		log.Println("RabbitMQ consumer started successfully")
	}()
}

// setupRouter sets up the Gin router and middleware.
func setupRouter(cfg *config.Config, deps *bootstrap.Dependencies, handlers *handlers.Handlers) *gin.Engine {
	router := gin.Default()

	router.Use(deps.LoggerMiddleware)
	router.Use(commonMiddleware.RequestLoggingMiddleware())

	api := router.Group("/api")
	api.Use(middleware.JWTAuthMiddleware(cfg.JWTSecret))
	routes.RegisterAPIRoutes(api, handlers)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}

// startServerWithGracefulShutdown starts the Gin server and handles graceful shutdown.
func startServerWithGracefulShutdown(cfg *config.Config, router *gin.Engine) {
	port := cfg.Port
	if port == "" {
		port = defaultPort
	}

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	log.Printf("Server running at http://localhost:%s", port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutting down server...")
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	log.Println("Server shutdown complete")
}
