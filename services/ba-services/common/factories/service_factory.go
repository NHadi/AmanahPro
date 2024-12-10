package factories

import (
	"AmanahPro/services/ba-services/internal/application/services"
	"AmanahPro/services/ba-services/internal/domain/repositories"
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

	// Create the baService instance
	baService := services.NewBAService(
		repos.BARepository,
		repos.BASectionRepository,
		repos.BADetailRepository,
		repos.BAProgressRepository,
		rabbitPublisher,
		"ba_events",
		sphGrpcClient,
	)

	return &services.Services{
		BAService: baService,
	}
}
