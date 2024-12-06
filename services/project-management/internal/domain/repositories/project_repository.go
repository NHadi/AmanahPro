package repositories

import (
	"AmanahPro/services/project-management/internal/domain/models"
	"AmanahPro/services/project-management/internal/dto"
)

type ProjectRepository interface {
	Create(project *models.Project) error
	Update(project *models.Project) error
	Delete(id int) error
	SearchProjectsByOrganization(organizationID int, query string) ([]dto.ProjectDTO, error)
	GetByID(id int, loadRelations bool) (*models.Project, error)
}
