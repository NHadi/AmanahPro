package factories

import (
	"github.com/NHadi/AmanahPro-common/services"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/streadway/amqp"
)

func CreateConsumers(esClient *elasticsearch.Client, rabbitChannel *amqp.Channel) map[string]*services.ConsumerService {
	return map[string]*services.ConsumerService{
		"project_events": services.NewConsumerService(esClient, "projects", "even-stores", "project_events"),
	}
}
