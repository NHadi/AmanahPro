package repositories

import "AmanahPro/services/spk-services/internal/domain/models"

// SPKRepository defines methods for accessing SPK data
type SPKRepository interface {
	Create(spk *models.SPK) error
	Update(spk *models.SPK) error
	Delete(spkId int) error
	GetByID(spkId int, loadDetails bool) (*models.SPK, error)
	Filter(organizationID int, spkID *int, projectID *int) ([]models.SPK, error)
}
