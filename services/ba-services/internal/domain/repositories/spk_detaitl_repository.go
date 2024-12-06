package repositories

import "AmanahPro/services/ba-services/internal/domain/models"

// SPKDetailRepository defines methods for accessing SPKDetail data
type SPKDetailRepository interface {
	Create(detail *models.SPKDetail) error
	Update(detail *models.SPKDetail) error
	Delete(detailId int) error
	GetByID(detailId int) (*models.SPKDetail, error)
}
