package services

import (
	"AmanahPro/services/bank-services/internal/application/dto"
	"AmanahPro/services/bank-services/internal/domain/models"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/streadway/amqp"
)

type ConsumerService struct {
	esClient *elasticsearch.Client
	esIndex  string
	rabbitCh *amqp.Channel
	rabbitQ  string
}

// NewConsumerService creates a new instance of the ConsumerService
func NewConsumerService(esClient *elasticsearch.Client, esIndex string, rabbitCh *amqp.Channel, rabbitQ string) *ConsumerService {
	return &ConsumerService{
		esClient: esClient,
		esIndex:  esIndex,
		rabbitCh: rabbitCh,
		rabbitQ:  rabbitQ,
	}
}

// StartConsumer starts listening to RabbitMQ and indexing messages into Elasticsearch
func (c *ConsumerService) StartConsumer() error {
	// Declare the queue in case it doesn't already exist
	_, err := c.rabbitCh.QueueDeclare(
		c.rabbitQ, // queue name
		true,      // durable
		false,     // auto-delete
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare RabbitMQ queue: %v", err)
	}

	// Consume messages from the queue
	msgs, err := c.rabbitCh.Consume(
		c.rabbitQ, // queue name
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to register a consumer: %v", err)
	}

	// Start a goroutine to process incoming messages
	go func() {
		for d := range msgs {
			var transactions []models.BankAccountTransactions
			err := json.Unmarshal(d.Body, &transactions)
			if err != nil {
				log.Printf("Error parsing RabbitMQ message: %v", err)
				continue
			}

			// Index transactions into Elasticsearch
			for _, transaction := range transactions {
				err := c.indexTransaction(transaction)
				if err != nil {
					log.Printf("Error indexing transaction: %v", err)
				}
			}
		}
	}()

	log.Printf("Consumer is now listening to queue: %s", c.rabbitQ)
	return nil
}

// indexTransaction indexes a single transaction into Elasticsearch
func (c *ConsumerService) indexTransaction(transaction models.BankAccountTransactions) error {
	// Map the model to the DTO
	dtoTransaction := dto.BankAccountTransactionDTO{
		ID:         transaction.ID,
		AccountID:  transaction.AccountID,
		BatchID:    transaction.BatchID,
		Tanggal:    transaction.Tanggal.Format("2006-01-02"),  // Format the date as ISO string
		Keterangan: strings.Trim(transaction.Keterangan, `"`), // Remove quotes if necessary
		Cabang:     transaction.Cabang,
		Credit:     transaction.Credit,
		Debit:      transaction.Debit,
		Saldo:      transaction.Saldo,
	}

	// Marshal the DTO to JSON
	data, err := json.Marshal(dtoTransaction)
	if err != nil {
		return fmt.Errorf("failed to marshal DTO transaction: %v", err)
	}

	docID := fmt.Sprintf("%d", transaction.ID) // Use the transaction ID as the document ID

	// Index the document in Elasticsearch
	res, err := c.esClient.Index(
		c.esIndex,                              // Index name
		strings.NewReader(string(data)),        // Document body
		c.esClient.Index.WithDocumentID(docID), // Document ID
		c.esClient.Index.WithRefresh("true"),   // Refresh the index
	)
	if err != nil {
		return fmt.Errorf("failed to index transaction in Elasticsearch: %v", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error indexing transaction: %s", res.String())
	}

	log.Printf("Indexed transaction ID: %d to Elasticsearch", transaction.ID)
	return nil
}
