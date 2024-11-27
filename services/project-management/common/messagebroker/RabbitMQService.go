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

func InitializeRabbitMQService(connURL string, queueNames []string) (*RabbitMQService, error) {
	// Initialize RabbitMQ connection and channel
	rabbitService, err := NewRabbitMQService(connURL)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize RabbitMQ service: %w", err)
	}

	// Declare all required queues
	if err := rabbitService.DeclareQueues(queueNames); err != nil {
		rabbitService.Close() // Ensure resources are cleaned up
		return nil, fmt.Errorf("failed to declare queues: %w", err)
	}

	return rabbitService, nil
}

func (s *RabbitMQService) DeclareQueues(queueNames []string) error {
	for _, queueName := range queueNames {
		if err := s.DeclareQueue(queueName); err != nil {
			return fmt.Errorf("failed to declare queue '%s': %w", queueName, err)
		}
	}
	return nil
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
