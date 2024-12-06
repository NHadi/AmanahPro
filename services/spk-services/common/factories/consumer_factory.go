package factories

import (
	"github.com/NHadi/AmanahPro-common/services"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/streadway/amqp"
)

func CreateConsumers(esClient *elasticsearch.Client, rabbitChannel *amqp.Channel) map[string]*services.ConsumerService {
	return map[string]*services.ConsumerService{
		"spk_events": services.NewConsumerService(esClient, "spks", "even-stores", "spk_events"),
	}
}
