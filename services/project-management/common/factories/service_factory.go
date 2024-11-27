package factories

import (
	"AmanahPro/services/project-management/internal/application/services"
	"AmanahPro/services/project-management/internal/domain/repositories"

	"AmanahPro/services/project-management/common/messagebroker"

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
		ProjectService:      services.NewProjectService(repos.ProjectRepository, rabbitPublisher, "project_events"),
		ProjectUserService:  services.NewProjectUserService(repos.ProjectUserRepository, rabbitPublisher, "project_user_events"),
		ProjectRecapService: services.NewProjectRecapService(repos.ProjectRecapRepository, rabbitPublisher, "project_recap_events"),
	}
}
