package repositories

import "AmanahPro/services/ba-services/internal/domain/models"

// SPKRepository defines methods for accessing SPK data
type SPKRepository interface {
	Create(ba *models.SPK) error
	Update(ba *models.SPK) error
	Delete(baId int) error
	GetByID(baId int) (*models.SPK, error)
	Filter(organizationID int, baID *int, projectID *int) ([]models.SPK, error)
}
