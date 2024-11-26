package services

import (
	"AmanahPro/services/bank-services/internal/domain/models"
	"encoding/json"
	"fmt"
	"log"

	"github.com/NHadi/AmanahPro-common/messagebroker"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/streadway/amqp"
)

type ConsumerService struct {
	esClient              *elasticsearch.Client
	esIndex               string
	rabbitConsumer        *messagebroker.RabbitMQConsumer
	queueName             string
	reconciliationService *ReconciliationService
}

// NewConsumerService creates a new instance of the ConsumerService
func NewConsumerService(esClient *elasticsearch.Client, esIndex string, rabbitConsumer *messagebroker.RabbitMQConsumer, queueName string, reconciliationService *ReconciliationService) *ConsumerService {
	return &ConsumerService{
		esClient:              esClient,
		esIndex:               esIndex,
		rabbitConsumer:        rabbitConsumer,
		queueName:             queueName,
		reconciliationService: reconciliationService,
	}
}

// StartConsumer starts listening to RabbitMQ and delegating message processing
func (c *ConsumerService) StartConsumer() error {
	err := c.rabbitConsumer.Consume(c.queueName, func(msg amqp.Delivery) error {
		// Decode the JSON message into the target struct
		var transactions []models.BankAccountTransactions
		err := json.Unmarshal(msg.Body, &transactions)
		if err != nil {
			log.Printf("Error parsing message to struct: %v", err)
			return err
		}

		// Delegate to the ReconciliationService
		err = c.reconciliationService.BulkIndexTransactions(transactions)
		if err != nil {
			log.Printf("Bulk indexing error: %v", err)
			return err
		}

		return nil // Successfully processed the message
	})
	if err != nil {
		return fmt.Errorf("failed to start RabbitMQ consumer: %v", err)
	}

	log.Printf("Consumer is now listening to queue: %s", c.queueName)
	return nil
}
