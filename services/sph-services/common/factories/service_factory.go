package factories

import (
	"AmanahPro/services/sph-services/internal/application/services"
	"AmanahPro/services/sph-services/internal/domain/repositories"

	"github.com/NHadi/AmanahPro-common/messagebroker"

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
	// Create the SphService instance
	sphService := services.NewSphService(
		repos.SphRepository,
		repos.SphSectionRepository,
		repos.SphDetailRepository,
		rabbitPublisher,
		"sph_events",
	)

	// Create the gRPC wrapper for SphService
	grpcSphService := services.NewGrpcSphService(sphService)

	// Return the services including the gRPC service
	return &services.Services{
		SphService:     sphService,
		GrpcSphService: grpcSphService,
	}
}
