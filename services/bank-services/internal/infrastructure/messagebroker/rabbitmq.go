package messagebroker

import (
	"fmt"

	"github.com/streadway/amqp"
)

type RabbitPublisher struct {
	channel *amqp.Channel
	queue   string
}

// NewRabbitPublisher initializes a RabbitMQ publisher
func NewRabbitPublisher(conn *amqp.Connection, queue string) (*RabbitPublisher, error) {
	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to create channel: %v", err)
	}

	// Declare a queue to ensure it exists
	_, err = ch.QueueDeclare(
		queue, // Name of the queue
		true,  // Durable
		false, // Auto-deleted
		false, // Exclusive
		false, // No-wait
		nil,   // Arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare queue: %v", err)
	}

	return &RabbitPublisher{channel: ch, queue: queue}, nil
}

// Publish sends a message to the RabbitMQ queue
func (p *RabbitPublisher) Publish(message []byte) error {
	if p.channel == nil {
		return fmt.Errorf("channel is not initialized")
	}

	err := p.channel.Publish(
		"",      // Exchange (default exchange)
		p.queue, // Routing key (queue name)
		false,   // Mandatory
		false,   // Immediate
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

// Close closes the RabbitMQ channel
func (p *RabbitPublisher) Close() error {
	if p.channel != nil {
		return p.channel.Close()
	}
	return nil
}
