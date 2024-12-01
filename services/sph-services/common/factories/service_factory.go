package factories

import (
	"AmanahPro/services/sph-services/internal/application/services"
	"AmanahPro/services/sph-services/internal/domain/repositories"

	"AmanahPro/services/sph-services/common/messagebroker"

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
		SphService: services.NewSphService(repos.SphRepository, repos.SphSectionRepository, repos.SphDetailRepository, rabbitPublisher, "sph_events"),
	}
}
