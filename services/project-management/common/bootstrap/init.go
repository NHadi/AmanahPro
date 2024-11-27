package bootstrap

import (
	"AmanahPro/services/project-management/common/config"
	"AmanahPro/services/project-management/common/messagebroker"
	"AmanahPro/services/project-management/internal/infrastructure/persistence"
	"log"

	"github.com/NHadi/AmanahPro-common/middleware"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

type Dependencies struct {
	DB                  *gorm.DB
	RabbitMQService     *messagebroker.RabbitMQService
	RabbitMQPublisher   *messagebroker.RabbitMQPublisher
	RabbitMQConsumer    *messagebroker.RabbitMQConsumer
	ElasticsearchClient *elasticsearch.Client
	RedisClient         *redis.Client
	Scheduler           *cron.Cron
	LoggerMiddleware    gin.HandlerFunc
}

func InitDependencies(cfg *config.Config) (*Dependencies, error) {
	// Initialize Database
	db, err := persistence.InitializeDB(cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}

	// List of queues to declare
	queueNames := []string{
		"project_events",
		"project_user_events",
		"project_recap_events",
	}

	// Initialize RabbitMQ service
	rabbitService, err := messagebroker.InitializeRabbitMQService(cfg.RabbitMQURL, queueNames)
	if err != nil {
		log.Fatalf("Error initializing RabbitMQ: %v", err)
	}

	// Application logic here
	log.Println("RabbitMQ service initialized and queues declared successfully.")

	rabbitPublisher := messagebroker.NewRabbitMQPublisher(rabbitService)
	rabbitConsumer := messagebroker.NewRabbitMQConsumer(rabbitService)

	// Initialize Elasticsearch
	esClient, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{cfg.ElasticsearchURL},
	})
	if err != nil {
		return nil, err
	}

	// Initialize Redis
	redisClient, err := persistence.InitializeRedis(cfg.RedisURL, cfg.RedisPassword, 0)
	if err != nil {
		return nil, err
	}

	// Initialize Scheduler
	scheduler := cron.New()

	logger, err := middleware.InitializeLogger("project-management", cfg.ElasticsearchURL, "project-management-logs")
	if err != nil {
		return nil, err
	}
	loggerMiddleware := middleware.GinLoggingMiddleware(logger)

	return &Dependencies{
		DB:                  db,
		RabbitMQService:     rabbitService,
		RabbitMQPublisher:   rabbitPublisher,
		RabbitMQConsumer:    rabbitConsumer,
		ElasticsearchClient: esClient,
		RedisClient:         redisClient,
		Scheduler:           scheduler,
		LoggerMiddleware:    loggerMiddleware,
	}, nil
}
