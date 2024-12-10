package repositories

import "AmanahPro/services/ba-services/internal/domain/models"

// BADetailRepository defines methods for accessing BADetail data
type BADetailRepository interface {
	Create(detail *models.BADetail) error
	Update(detail *models.BADetail) error
	Delete(detailId int) error
	GetByID(detailId int) (*models.BADetail, error)
}
