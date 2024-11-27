package services

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/streadway/amqp"
)

func InitializeConsumerServices(
	esClient *elasticsearch.Client,
	rabbitChannel *amqp.Channel,
) (*Services, error) {
	// Service initialization logic
	return &Services{
		ProjectConsumer:      NewConsumerService(esClient, "projects", "project_events"),
		ProjectRecapConsumer: NewConsumerService(esClient, "project_recap", "project_recap_events"),
		ProjectUserConsumer:  NewConsumerService(esClient, "project_user", "project_user_events"),
	}, nil
}
