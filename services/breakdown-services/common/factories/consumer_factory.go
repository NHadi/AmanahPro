package factories

import (
	"AmanahPro/services/breakdown-services/internal/application/services"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/streadway/amqp"
)

func CreateConsumers(esClient *elasticsearch.Client, rabbitChannel *amqp.Channel) map[string]services.Consumer {
	return map[string]services.Consumer{
		"breakdown_events":         services.NewConsumerService(esClient, "breakdowns", "breakdown_events"),
		"project_events_breakdown": services.NewProjectEventConsumer(esClient, "project_events_breakdown"),
	}
}
