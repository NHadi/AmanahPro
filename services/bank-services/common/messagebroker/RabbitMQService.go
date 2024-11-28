package messagebroker

import (
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

type RabbitMQService struct {
	Conn          *amqp.Connection
	Channel       *amqp.Channel
	connURL       string
	notifyClose   chan *amqp.Error
	reconnectWait time.Duration
	queueNames    []string
}

// NewRabbitMQService initializes RabbitMQ with auto-reconnection and queue declarations
func NewRabbitMQService(connURL string, queueNames []string) (*RabbitMQService, error) {
	service := &RabbitMQService{
		connURL:       connURL,
		reconnectWait: 5 * time.Second, // Wait time before reconnection attempts
		queueNames:    queueNames,
	}
	if err := service.connect(); err != nil {
		return nil, err
	}
	go service.handleReconnect()
	return service, nil
}

// connect establishes a new RabbitMQ connection and channel, and declares queues
func (s *RabbitMQService) connect() error {
	conn, err := amqp.Dial(s.connURL)
	if err != nil {
		return fmt.Errorf("failed to connect to RabbitMQ: %v", err)
	}
	channel, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to create RabbitMQ channel: %v", err)
	}

	s.Conn = conn
	s.Channel = channel
	s.notifyClose = make(chan *amqp.Error)
	s.Conn.NotifyClose(s.notifyClose)

	log.Println("RabbitMQ connected successfully")

	// Declare queues upon connection
	if err := s.DeclareQueues(); err != nil {
		return fmt.Errorf("failed to declare queues: %v", err)
	}
	return nil
}

// handleReconnect listens for connection closure and attempts to reconnect
func (s *RabbitMQService) handleReconnect() {
	for {
		err := <-s.notifyClose
		if err != nil {
			log.Printf("RabbitMQ connection lost: %v. Attempting to reconnect...", err)
		}

		// Attempt reconnection
		for {
			if err := s.connect(); err != nil {
				log.Printf("RabbitMQ reconnection failed: %v. Retrying in %s...", err, s.reconnectWait)
				time.Sleep(s.reconnectWait)
			} else {
				log.Println("RabbitMQ reconnected successfully")
				break
			}
		}
	}
}

// DeclareQueues declares all required queues
func (s *RabbitMQService) DeclareQueues() error {
	for _, queueName := range s.queueNames {
		if err := s.DeclareQueue(queueName); err != nil {
			return fmt.Errorf("failed to declare queue '%s': %w", queueName, err)
		}
	}
	log.Println("RabbitMQ queues declared successfully")
	return nil
}

// DeclareQueue declares a single queue
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

// NewChannel creates and returns a new channel for use by publishers or consumers
func (s *RabbitMQService) NewChannel() (*amqp.Channel, error) {
	if s.Conn == nil || s.Conn.IsClosed() {
		return nil, amqp.ErrClosed
	}
	return s.Conn.Channel()
}

// Close cleans up the RabbitMQ connection and channel
func (s *RabbitMQService) Close() {
	if s.Channel != nil {
		_ = s.Channel.Close()
	}
	if s.Conn != nil {
		_ = s.Conn.Close()
	}
}
