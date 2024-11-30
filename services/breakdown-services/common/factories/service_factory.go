package factories

import (
	"AmanahPro/services/breakdown-services/internal/application/services"
	"AmanahPro/services/breakdown-services/internal/domain/repositories"

	"AmanahPro/services/breakdown-services/common/messagebroker"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-redis/redis/v8"
)

func CreateServices(
	repos *repositories.Repositories,
	rabbitPublisher *messagebroker.RabbitMQPublisher,
	rabbitConsumer *messagebroker.RabbitMQConsumer,
	esClient *elasticsearch.Client,
	redisClient *redis.Client,
) *services.Services {
	return &services.Services{
		BreakdownService: services.NewBreakdownService(repos.BreakdownRepository, repos.BreakdownSectionRepository, repos.BreakdownItemRepository, rabbitPublisher, "breakdown_events"),
	}
}
