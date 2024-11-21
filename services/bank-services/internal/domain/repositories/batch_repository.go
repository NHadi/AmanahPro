package repositories

import (
	"AmanahPro/services/bank-services/internal/domain/models"
	"time"
)

type BatchRepository interface {
	Create(batch *models.UploadBatch) error
	BatchExists(accountID uint, periodeStart, periodeEnd time.Time) (bool, error)
}
