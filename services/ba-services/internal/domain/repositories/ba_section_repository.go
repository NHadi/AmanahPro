package repositories

import "AmanahPro/services/ba-services/internal/domain/models"

// BASectionRepository defines methods for accessing BASection data
type BASectionRepository interface {
	Create(section *models.BASection) error
	Update(section *models.BASection) error
	Delete(sectionId int) error
	GetByID(sectionId int) (*models.BASection, error)
}
