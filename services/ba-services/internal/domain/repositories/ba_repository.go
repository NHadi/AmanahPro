package repositories

import "AmanahPro/services/ba-services/internal/domain/models"

// BARepository defines methods for accessing BA data
type BARepository interface {
	Create(ba *models.BA) error
	Update(ba *models.BA) error
	Delete(baId int) error
	GetByID(baId int, loadSections bool) (*models.BA, error)
	Filter(organizationID int, baID *int, projectID *int) ([]models.BA, error)
}
