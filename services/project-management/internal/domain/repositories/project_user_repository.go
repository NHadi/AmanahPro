package repositories

import (
	"AmanahPro/services/project-management/internal/domain/models"
	"AmanahPro/services/project-management/internal/dto"
)

type ProjectUserRepository interface {
	// Create inserts a new project user into the database
	Create(user *models.ProjectUser) error

	// Update modifies an existing project user in the database
	Update(user *models.ProjectUser) error

	// Delete removes a project user from the database by ID
	Delete(id int) error

	// FindByProjectID retrieves project users associated with a specific project ID from Elasticsearch
	FindByProjectID(projectID int, organizationID *int) ([]dto.ProjectUserDTO, error)

	FindByUserAndProject(userID, projectID int, organizationID *int) (*dto.ProjectUserDTO, error)
}
