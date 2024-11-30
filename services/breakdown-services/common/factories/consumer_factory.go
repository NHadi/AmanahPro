package factories

import (
	"AmanahPro/services/breakdown-services/internal/application/services"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/streadway/amqp"
)

func CreateConsumers(esClient *elasticsearch.Client, rabbitChannel *amqp.Channel) map[string]*services.ConsumerService {
	return map[string]*services.ConsumerService{
		"breakdown_events": services.NewConsumerService(esClient, "breakdowns", "breakdown_events"),
	}
}
