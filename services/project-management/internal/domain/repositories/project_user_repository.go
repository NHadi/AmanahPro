package repositories

import (
	"AmanahPro/services/project-management/internal/domain/models"
)

type ProjectUserRepository interface {
	Create(projectUser *models.ProjectUser) error
	Update(projectUser *models.ProjectUser) error
	Delete(id int) error
	FindByID(id int) (*models.ProjectUser, error)
	FindAllByProjectID(projectID int) ([]models.ProjectUser, error)
}
