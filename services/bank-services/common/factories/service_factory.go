package factories

import (
	"AmanahPro/services/bank-services/internal/application/services"
	"AmanahPro/services/bank-services/internal/domain/repositories"

	"AmanahPro/services/bank-services/common/messagebroker"

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
		UploadService:         services.NewUploadService(repos.TransactionRepo, repos.BatchRepo, rabbitPublisher),
		TransactionService:    services.NewTransactionService(repos.TransactionRepo),
		ReconciliationService: services.NewReconciliationService(esClient, "bank-transactions", redisClient, repos.TransactionRepo),
		ConsumerService:       services.NewConsumerService(esClient, "bank-transactions", rabbitConsumer, "transactions_queue", services.NewReconciliationService(esClient, "bank-transactions", redisClient, repos.TransactionRepo)),
	}
}
