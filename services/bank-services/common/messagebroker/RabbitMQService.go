package messagebroker

import (
	"fmt"

	"github.com/streadway/amqp"
)

type RabbitMQService struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

// NewRabbitMQService initializes RabbitMQ connection and channel
func NewRabbitMQService(connURL string) (*RabbitMQService, error) {
	conn, err := amqp.Dial(connURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %v", err)
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to create RabbitMQ channel: %v", err)
	}

	return &RabbitMQService{Conn: conn, Channel: channel}, nil
}

// DeclareQueue declares a queue if it doesn't exist
func (s *RabbitMQService) DeclareQueue(queueName string) error {
	_, err := s.Channel.QueueDeclare(
		queueName,
		true,  // Durable
		false, // Auto-delete
		false, // Exclusive
		false, // No-wait
		nil,   // Arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare queue: %v", err)
	}
	return nil
}

// Close closes RabbitMQ connection and channel
func (s *RabbitMQService) Close() {
	if s.Channel != nil {
		_ = s.Channel.Close()
	}
	if s.Conn != nil {
		_ = s.Conn.Close()
	}
}
