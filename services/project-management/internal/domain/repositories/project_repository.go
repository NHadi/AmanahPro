package repositories

import (
	"AmanahPro/services/project-management/internal/domain/models"
)

type ProjectRepository interface {
	Create(project *models.Project) error
	Update(project *models.Project) error
	Delete(id int) error
	FindByID(id int) (*models.Project, error)
	FindAllByOrganizationID(organizationID int) ([]models.Project, error)
}
