package factories

import (
	"AmanahPro/services/project-management/internal/application/services"
	"AmanahPro/services/project-management/internal/domain/repositories"

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
	return &services.Services{
		ProjectService:          services.NewProjectService(repos.ProjectRepository, repos.ProjectUserRepository, rabbitPublisher, "project_events"),
		ProjectFinancialService: services.NewProjectFinancialService(repos.ProjectFinancialRepository),
	}
}
