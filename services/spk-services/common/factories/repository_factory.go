package factories

import (
	domainRepo "AmanahPro/services/spk-services/internal/domain/repositories"
	"AmanahPro/services/spk-services/internal/infrastructure/repositories"

	"github.com/elastic/go-elasticsearch/v8"
	"gorm.io/gorm"
)

// CreateRepositories initializes and returns the repositories for Spk
func CreateRepositories(db *gorm.DB, esClient *elasticsearch.Client) *domainRepo.Repositories {
	return &domainRepo.Repositories{
		SpkRepository:        repositories.NewSPKRepository(db, esClient, "spks"),
		SpkSectionRepository: repositories.NewSPKSectionRepository(db),
		SpkDetailRepository:  repositories.NewSPKDetailRepository(db),
	}
}
