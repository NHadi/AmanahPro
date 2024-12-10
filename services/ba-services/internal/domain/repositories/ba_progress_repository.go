package repositories

import "AmanahPro/services/ba-services/internal/domain/models"

// BAProgressRepository defines methods for accessing BAProgress data
type BAProgressRepository interface {
	Create(progress *models.BAProgress) error
	Update(progress *models.BAProgress) error
	Delete(progressId int) error
	GetByID(progressId int) (*models.BAProgress, error)
}
