package bootstrap

import (
	"AmanahPro/services/spk-services/common/config"
	"AmanahPro/services/spk-services/common/factories"
	"log"
	"strconv"
	"time"

	"github.com/NHadi/AmanahPro-common/infrastructure/persistence"
	"github.com/NHadi/AmanahPro-common/messagebroker"
	commonServices "github.com/NHadi/AmanahPro-common/services"

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

	// List of queues to declare
	queueNames := []string{
		"spk_events",
	}

	// Initialize RabbitMQ service
	rabbitService, err := messagebroker.NewRabbitMQService(cfg.RabbitMQURL, queueNames)
	if err != nil {
		log.Fatalf("Error initializing RabbitMQ: %v", err)
	}

	// Initialize RabbitMQ publisher
	rabbitPublisher := messagebroker.NewRabbitMQPublisher(rabbitService)

	// Initialize RabbitMQ consumer
	rabbitConsumer := messagebroker.NewRabbitMQConsumer(rabbitService)
	// Initialize consumers and start them
	consumers := factories.CreateConsumers(esClient, rabbitService.Channel)
	startConsumers(cfg, consumers, rabbitService)

	// Setup RabbitMQ reconnection handling
	rabbitService.SetOnReconnect(func() {
		log.Println("RabbitMQ reconnected. Reinitializing consumers...")
		for queueName, consumer := range consumers {
			go func(c *commonServices.ConsumerService, q string) {
				for {
					channel, err := rabbitService.NewChannel()
					if err != nil {
						log.Printf("Failed to create channel for consumer %s: %v. Retrying in 5s...", q, err)
						time.Sleep(5 * time.Second)
						continue
					}

					if err := c.StartConsumer(channel, 5); err != nil {
						log.Printf("Consumer for queue '%s' exited with error: %v. Restarting...", q, err)
						time.Sleep(5 * time.Second)
					} else {
						break
					}
				}
			}(consumer, queueName)
		}
	})

	// Initialize Scheduler
	scheduler := cron.New()

	// Initialize Logger
	logger, err := middleware.InitializeLogger("spk-services", cfg.ElasticsearchURL, "spk-services-logs")
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

// startConsumers starts RabbitMQ consumers for the provided queues.
func startConsumers(cfg *config.Config, consumers map[string]*commonServices.ConsumerService, rabbitService *messagebroker.RabbitMQService) {
	concurrency, err := strconv.Atoi(cfg.CONCURRENCY)
	if err != nil || concurrency == 0 {
		concurrency = 5
	}

	for queueName, consumer := range consumers {
		go func(c *commonServices.ConsumerService, q string) {
			for {
				log.Printf("Starting consumer for queue: %s", q)

				// Create a channel for the consumer
				channel, err := rabbitService.NewChannel()
				if err != nil {
					log.Printf("Failed to create channel for consumer %s: %v. Retrying in 5s...", q, err)
					time.Sleep(5 * time.Second)
					continue
				}

				// Set QoS (optional, based on your use case)
				err = channel.Qos(
					10,    // Prefetch count
					0,     // Prefetch size
					false, // Apply per channel
				)
				if err != nil {
					log.Printf("Failed to set QoS for consumer %s: %v", q, err)
					time.Sleep(5 * time.Second)
					continue
				}

				// Start consuming
				if err := c.StartConsumer(channel, concurrency); err != nil {
					log.Printf("Consumer for queue '%s' exited with error: %v. Restarting...", q, err)
					time.Sleep(5 * time.Second)
				}
			}
		}(consumer, queueName)
	}

	log.Printf("RabbitMQ Consumers started with concurrency: %d", concurrency)
}
