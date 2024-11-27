package main

import (
	"AmanahPro/services/project-management/common/bootstrap"
	"AmanahPro/services/project-management/common/config"
	"AmanahPro/services/project-management/common/factories"
	"AmanahPro/services/project-management/common/messagebroker"
	commonMiddleware "AmanahPro/services/project-management/common/middleware"
	"AmanahPro/services/project-management/common/routes"
	"AmanahPro/services/project-management/internal/application/services"
	"AmanahPro/services/project-management/internal/handlers"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"strconv"
	"time"

	_ "AmanahPro/services/project-management/docs"

	"github.com/NHadi/AmanahPro-common/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Project Management Services API
// @version 1.0
// @description This is the Project Management Services API documentation for managing project, and reconciliations.
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Provide your JWT token with "Bearer " prefix, e.g., "Bearer <token>"
// @host localhost:8083
// @BasePath /
const (
	defaultPort = "8083"
	logDir      = "log"
)

func main() {
	defer recoverFromPanic()

	configureLogging()
	log.Println("Starting Project Management Services API...")

	cfg := loadConfig()
	deps := initializeDependencies(cfg)

	repos := factories.CreateRepositories(deps.DB, deps.ElasticsearchClient)
	services := factories.CreateServices(repos, deps.RabbitMQPublisher, deps.RabbitMQConsumer, deps.ElasticsearchClient, deps.RedisClient)
	consumers := factories.CreateConsumers(deps.ElasticsearchClient, deps.RabbitMQService.Channel)

	// Start RabbitMQ consumers
	startRabbitMQConsumers(cfg, consumers, deps.RabbitMQService)

	router := setupRouter(cfg, deps, handlers.NewHandlers(
		handlers.NewProjectHandler(services.ProjectService),
	))
	startServerWithGracefulShutdown(deps, cfg, router)
}

// configureLogging sets up daily logging into a file.
func configureLogging() {
	// Check if the log directory exists; if not, create it
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		if err := os.Mkdir(logDir, os.ModePerm); err != nil {
			log.Fatalf("Failed to create log directory: %v", err)
		}
	}

	// Define the log file name based on the current date
	logFileName := fmt.Sprintf("%s/Log%s.log", logDir, time.Now().Format("20060102"))
	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	// Check the APP_ENV environment variable
	appEnv := os.Getenv("APP_ENV")
	if appEnv != "PRODUCTION" {
		// If not production, write logs to both the terminal and the log file
		log.SetOutput(io.MultiWriter(os.Stdout, logFile))
	} else {
		// If production, write logs only to the log file
		log.SetOutput(logFile)
	}

	// Set log flags for consistent log formatting
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Printf("Logging initialized: %s", logFileName)
	log.Printf("Environment: %s", appEnv)
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

func startRabbitMQConsumers(cfg *config.Config, consumers map[string]*services.ConsumerService, rabbitService *messagebroker.RabbitMQService) {
	concurrency, _ := strconv.Atoi(cfg.CONCURRENCY)

	if concurrency == 0 {
		concurrency = 5
	}

	// Process each consumer
	for queueName, consumer := range consumers {
		go func(c *services.ConsumerService, q string) {
			log.Printf("Starting consumer for queue: %s", q)

			// Ensure RabbitMQ connection is active
			if rabbitService.Conn == nil || rabbitService.Conn.IsClosed() {
				log.Printf("RabbitMQ connection is not open for queue: %s", q)
				return
			}

			// Create a channel for the consumer
			channel, err := rabbitService.Conn.Channel()
			if err != nil {
				log.Printf("Failed to create channel for consumer %s: %v", q, err)
				return
			}

			// Do not close the channel here; let RabbitMQ manage it
			if err := c.StartConsumer(channel, concurrency); err != nil {
				log.Printf("Failed to start consumer for queue '%s': %v", q, err)
			}
		}(consumer, queueName)
	}

	log.Printf("RabbitMQ Consumers started with concurrency: %d", concurrency)
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
func startServerWithGracefulShutdown(deps *bootstrap.Dependencies, cfg *config.Config, router *gin.Engine) {
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
	shutdownRabbitMQ(deps.RabbitMQService)
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	log.Println("Server shutdown complete")
}

func shutdownRabbitMQ(rabbitService *messagebroker.RabbitMQService) {
	log.Println("Closing RabbitMQ connection...")
	if rabbitService != nil {
		rabbitService.Close()
		log.Println("RabbitMQ connection closed.")
	}
}
