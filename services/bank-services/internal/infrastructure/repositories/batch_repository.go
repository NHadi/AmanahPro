package repositories

import (
	"AmanahPro/services/bank-services/internal/domain/models"
	"AmanahPro/services/bank-services/internal/domain/repositories"

	"gorm.io/gorm"
)

type BatchRepository struct {
	db *gorm.DB
}

func NewBatchRepository(db *gorm.DB) repositories.BatchRepository {
	return &BatchRepository{db: db}
}

func (r *BatchRepository) Create(batch *models.UploadBatch) error {
	if err := r.db.Create(batch).Error; err != nil {
		return err
	}
	return nil
}

// BatchExists checks if a batch already exists for the given account and period
func (r *BatchRepository) BatchExists(accountID, year, month uint) (bool, error) {
	var count int64
	err := r.db.Model(&models.UploadBatch{}).
		Where("AccountID = ? AND Year = ? AND Month = ?", accountID, year, month).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
