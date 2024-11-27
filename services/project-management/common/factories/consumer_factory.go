package factories

import (
	"AmanahPro/services/project-management/internal/application/services"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/streadway/amqp"
)

func CreateConsumers(esClient *elasticsearch.Client, rabbitChannel *amqp.Channel) map[string]*services.ConsumerService {
	return map[string]*services.ConsumerService{
		"project_events":       services.NewConsumerService(esClient, "projects", "project_events"),
		"project_user_events":  services.NewConsumerService(esClient, "project_user", "project_user_events"),
		"project_recap_events": services.NewConsumerService(esClient, "project_recap", "project_recap_events"),
	}
}
