package factories

import (
	domainRepo "AmanahPro/services/bank-services/internal/domain/repositories"
	"AmanahPro/services/bank-services/internal/infrastructure/repositories"

	"github.com/elastic/go-elasticsearch/v8"
	"gorm.io/gorm"
)

func CreateRepositories(db *gorm.DB, esClient *elasticsearch.Client) *domainRepo.Repositories {
	return &domainRepo.Repositories{
		BatchRepo:       repositories.NewBatchRepository(db),
		TransactionRepo: repositories.NewBankAccountTransactionRepository(db, esClient, "bank-transactions"),
	}
}
