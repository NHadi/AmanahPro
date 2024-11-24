package repositories

import (
	"AmanahPro/services/bank-services/internal/domain/models"
)

type BatchRepository interface {
	Create(batch *models.UploadBatch) error
	BatchExists(organizationID, year, month uint) (bool, error)
}
