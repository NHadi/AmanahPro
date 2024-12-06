package repositories

import "AmanahPro/services/ba-services/internal/domain/models"

// SPKSectionRepository defines methods for accessing SPKSection data
type SPKSectionRepository interface {
	Create(section *models.SPKSection) error
	Update(section *models.SPKSection) error
	Delete(sectionId int) error
	GetByID(sectionId int) (*models.SPKSection, error)
}
