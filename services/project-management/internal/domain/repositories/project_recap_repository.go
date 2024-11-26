package repositories

import (
	"AmanahPro/services/project-management/internal/domain/models"
)

type ProjectRecapRepository interface {
	Create(projectRecap *models.ProjectRecap) error
	Update(projectRecap *models.ProjectRecap) error
	Delete(id int) error
	FindByID(id int) (*models.ProjectRecap, error)
	FindAllByOrganizationID(organizationID int) ([]models.ProjectRecap, error)
}
