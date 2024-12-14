package factories

import (
	domainRepo "AmanahPro/services/breakdown-services/internal/domain/repositories"
	"AmanahPro/services/breakdown-services/internal/infrastructure/repositories"

	"github.com/elastic/go-elasticsearch/v8"
	"gorm.io/gorm"
)

func CreateRepositories(db *gorm.DB, esClient *elasticsearch.Client) *domainRepo.Repositories {
	return &domainRepo.Repositories{
		BreakdownItemRepository:       repositories.NewBreakdownItemRepository(db),
		BreakdownSectionRepository:    repositories.NewBreakdownSectionRepository(db),
		MstBreakdownItemRepository:    repositories.NewMstBreakdownItemRepository(db),
		MstBreakdownSectionRepository: repositories.NewMstBreakdownSectionRepository(db),
		BreakdownRepository:           repositories.NewBreakdownRepository(db, esClient, "breakdowns"),
	}
}
