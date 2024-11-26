package bootstrap

import (
	"AmanahPro/services/bank-services/common/messagebroker"
	"AmanahPro/services/bank-services/config"
	"AmanahPro/services/bank-services/internal/infrastructure/persistence"

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

	// Initialize RabbitMQ
	rabbitService, err := messagebroker.NewRabbitMQService(cfg.RabbitMQURL)
	if err != nil {
		return nil, err
	}
	rabbitPublisher := messagebroker.NewRabbitMQPublisher(rabbitService)
	rabbitConsumer := messagebroker.NewRabbitMQConsumer(rabbitService)

	// Declare RabbitMQ Queue
	err = rabbitService.DeclareQueue("transactions_queue")
	if err != nil {
		return nil, err
	}

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

	logger, err := middleware.InitializeLogger("bank-service", cfg.ElasticsearchURL, "bank-services-logs")
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
