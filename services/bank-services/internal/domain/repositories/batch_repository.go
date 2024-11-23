package repositories

import (
	"AmanahPro/services/bank-services/internal/domain/models"
)

type BatchRepository interface {
	Create(batch *models.UploadBatch) error
	BatchExists(accountID, year, month uint) (bool, error)
}
