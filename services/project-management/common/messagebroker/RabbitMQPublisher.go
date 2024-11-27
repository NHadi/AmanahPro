package messagebroker

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)

type RabbitMQPublisher struct {
	service *RabbitMQService
}

// NewRabbitMQPublisher creates a new publisher
func NewRabbitMQPublisher(service *RabbitMQService) *RabbitMQPublisher {
	return &RabbitMQPublisher{service: service}
}

// Publish sends a message to the specified queue
func (p *RabbitMQPublisher) Publish(queueName string, message []byte) error {
	err := p.service.Channel.Publish(
		"",        // Default exchange
		queueName, // Queue name
		false,     // Mandatory
		false,     // Immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish message: %v", err)
	}
	return nil
}

// PublishEvent marshals an event and publishes it to the specified queue
func (p *RabbitMQPublisher) PublishEvent(queueName string, event interface{}) error {
	// Marshal the event to JSON
	message, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}

	// Publish the message
	if err := p.Publish(queueName, message); err != nil {
		return fmt.Errorf("failed to publish event: %w", err)
	}

	return nil
}
