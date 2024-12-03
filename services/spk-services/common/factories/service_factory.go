package factories

import (
	"AmanahPro/services/spk-services/internal/application/services"
	"AmanahPro/services/spk-services/internal/domain/repositories"
	"fmt"

	"github.com/NHadi/AmanahPro-common/messagebroker"
	"github.com/NHadi/AmanahPro-common/protos"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CreateServices(
	repos *repositories.Repositories,
	rabbitPublisher *messagebroker.RabbitMQPublisher,
	rabbitConsumer *messagebroker.RabbitMQConsumer,
	esClient *elasticsearch.Client,
	redisClient *redis.Client,
) *services.Services {
	// Initialize gRPC client
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to SPH service: %v", err))
	}
	sphGrpcClient := protos.NewSphServiceClient(conn)

	// Create the SpkService instance
	spkService := services.NewSpkService(
		repos.SpkRepository,
		repos.SpkSectionRepository,
		repos.SpkDetailRepository,
		rabbitPublisher,
		"spk_events",
		sphGrpcClient, // Inject SPH gRPC client
	)

	return &services.Services{
		SPKService: spkService,
	}
}
