package factories

import (
	"github.com/NHadi/AmanahPro-common/services"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/streadway/amqp"
)

func CreateConsumers(esClient *elasticsearch.Client, rabbitChannel *amqp.Channel) map[string]*services.ConsumerService {
	return map[string]*services.ConsumerService{
		"ba_events": services.NewConsumerService(esClient, "bas", "even-stores", "ba_events"),
	}
}
