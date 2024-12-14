package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/streadway/amqp"
)

type ProjectEventConsumer struct {
	esClient  *elasticsearch.Client
	queueName string
	handlers  map[string]func(map[string]interface{}, map[string]interface{}) error // Event-specific handlers
}

// NewProjectEventConsumer initializes the consumer for project-related events
func NewProjectEventConsumer(esClient *elasticsearch.Client, queueName string) *ProjectEventConsumer {
	service := &ProjectEventConsumer{
		esClient:  esClient,
		queueName: queueName,
		handlers:  make(map[string]func(map[string]interface{}, map[string]interface{}) error),
	}

	service.handlers["Updated"] = service.handleProjectUpdated

	return service
}

// StartConsumer starts consuming project-related events
func (c *ProjectEventConsumer) StartConsumer(channel *amqp.Channel, concurrency int) error {
	msgs, err := channel.Consume(
		c.queueName,
		"",
		false, // Manual acknowledgment
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to start consuming messages: %w", err)
	}

	workerChan := make(chan bool, concurrency)

	go func() {
		for msg := range msgs {
			workerChan <- true
			go func(m amqp.Delivery) {
				defer func() { <-workerChan }()
				if err := c.processMessage(m.Body); err != nil {
					log.Printf("Error processing message: %v", err)
					m.Nack(false, true) // Requeue on failure
				} else {
					m.Ack(false)
				}
			}(msg)
		}
	}()

	log.Printf("ProjectEventConsumer is now actively listening to queue: %s", c.queueName)
	select {}
}

// processMessage routes messages to appropriate handlers
func (c *ProjectEventConsumer) processMessage(msg []byte) error {
	log.Printf("Processing message from queue %s: %s", c.queueName, string(msg))

	var event struct {
		Event   string                 `json:"event"`
		Payload map[string]interface{} `json:"payload"`
		Meta    map[string]interface{} `json:"meta"` // Optional metadata
	}

	// Parse the message
	if err := json.Unmarshal(msg, &event); err != nil {
		return fmt.Errorf("failed to parse message: %w", err)
	}

	// Route to the appropriate handler
	handler, exists := c.handlers[event.Event]
	if !exists {
		log.Printf("Unhandled event type: %s", event.Event)
		return nil // Acknowledge unknown event types to prevent re-delivery
	}

	// Pass both payload and meta to the handler
	return handler(event.Payload, event.Meta)
}

// handleProjectUpdated updates breakdown records with the new project name
func (c *ProjectEventConsumer) handleProjectUpdated(payload map[string]interface{}, meta map[string]interface{}) error {
	// Extract projectId from the payload
	projectID, ok := payload["ProjectID"].(float64)
	if !ok {
		return fmt.Errorf("missing or invalid ProjectID in payload")
	}

	// Extract projectName from the payload
	projectName, ok := payload["ProjectName"].(string) // Corrected key
	if !ok {
		return fmt.Errorf("missing or invalid ProjectName in payload")
	}

	projectIDStr := fmt.Sprintf("%.0f", projectID)
	log.Printf("Updating project name for ProjectID: %s to %s", projectIDStr, projectName)

	// Update breakdown records in Elasticsearch
	return c.updateBreakdownProjectName(projectIDStr, projectName)
}

func (c *ProjectEventConsumer) updateBreakdownProjectName(projectID, projectName string) error {
	query := map[string]interface{}{
		"script": map[string]interface{}{
			"source": "ctx._source.projectName = params.projectName",
			"params": map[string]interface{}{
				"projectName": projectName,
			},
		},
		"query": map[string]interface{}{
			"term": map[string]interface{}{
				"ProjectId": projectID,
			},
		},
	}

	data, err := json.Marshal(query)
	if err != nil {
		return fmt.Errorf("failed to marshal update query: %w", err)
	}

	log.Printf("Executing UpdateByQuery with body: %s", string(data))

	req := esapi.UpdateByQueryRequest{
		Index: []string{"breakdowns"},
		Body:  bytes.NewReader(data),
	}

	res, err := req.Do(context.Background(), c.esClient)
	if err != nil {
		return fmt.Errorf("failed to execute update by query: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		body, _ := io.ReadAll(res.Body)
		log.Printf("Elasticsearch response error: %s", string(body))
		return fmt.Errorf("update by query returned an error: %s", res.String())
	}

	log.Printf("Elasticsearch UpdateByQuery successful for projectID: %s", projectID)
	return nil
}
