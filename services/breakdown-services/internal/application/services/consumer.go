package services

import (
	"github.com/streadway/amqp"
)

// Consumer defines the interface for all consumers
type Consumer interface {
	StartConsumer(channel *amqp.Channel, concurrency int) error
}
