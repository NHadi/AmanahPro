package factories

import (
	"AmanahPro/services/sph-services/internal/application/services"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/streadway/amqp"
)

func CreateConsumers(esClient *elasticsearch.Client, rabbitChannel *amqp.Channel) map[string]*services.ConsumerService {
	return map[string]*services.ConsumerService{
		"sph_events": services.NewConsumerService(esClient, "sphs", "sph_events"),
	}
}
