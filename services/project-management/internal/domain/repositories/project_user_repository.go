package repositories

import (
	"AmanahPro/services/project-management/internal/domain/models"
)

type ProjectUserRepository interface {
	// Create inserts a new project user into the database
	Create(user *models.ProjectUser) error

	// Update modifies an existing project user in the database
	Update(user *models.ProjectUser) error

	// Delete removes a project user from the database by ID
	Delete(id int) error
	GetByID(id int) (*models.ProjectUser, error)
	GetByProjectID(projectID int) ([]models.ProjectUser, error)
}
