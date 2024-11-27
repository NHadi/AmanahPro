package factories

import (
	domainRepo "AmanahPro/services/project-management/internal/domain/repositories"
	"AmanahPro/services/project-management/internal/infrastructure/repositories"

	"github.com/elastic/go-elasticsearch/v8"
	"gorm.io/gorm"
)

func CreateRepositories(db *gorm.DB, esClient *elasticsearch.Client) *domainRepo.Repositories {
	return &domainRepo.Repositories{
		ProjectRecapRepository: repositories.NewProjectRecapRepository(db, esClient, "projects"),
		ProjectRepository:      repositories.NewProjectRepository(db, esClient, "project_recap"),
		ProjectUserRepository:  repositories.NewProjectUserRepository(db, esClient, "project_user"),
	}
}
