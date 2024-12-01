package factories

import (
	domainRepo "AmanahPro/services/sph-services/internal/domain/repositories"
	"AmanahPro/services/sph-services/internal/infrastructure/repositories"

	"github.com/elastic/go-elasticsearch/v8"
	"gorm.io/gorm"
)

// CreateRepositories initializes and returns the repositories for SPH
func CreateRepositories(db *gorm.DB, esClient *elasticsearch.Client) *domainRepo.Repositories {
	return &domainRepo.Repositories{
		SphRepository:        repositories.NewSphRepository(db, esClient, "sphs"),
		SphSectionRepository: repositories.NewSphSectionRepository(db),
		SphDetailRepository:  repositories.NewSphDetailRepository(db),
	}
}
