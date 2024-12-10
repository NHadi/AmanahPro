package factories

import (
	domainRepo "AmanahPro/services/ba-services/internal/domain/repositories"
	"AmanahPro/services/ba-services/internal/infrastructure/repositories"

	"github.com/elastic/go-elasticsearch/v8"
	"gorm.io/gorm"
)

// CreateRepositories initializes and returns the repositories for Spk
func CreateRepositories(db *gorm.DB, esClient *elasticsearch.Client) *domainRepo.Repositories {
	return &domainRepo.Repositories{
		BARepository:         repositories.NewBARepository(db, esClient, "bas"),
		BAProgressRepository: repositories.NewBAProgressRepository(db),
		BASectionRepository:  repositories.NewBASectionRepository(db),
		BADetailRepository:   repositories.NewBADetailRepository(db),
	}
}
