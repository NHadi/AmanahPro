package repositories

import (
	"AmanahPro/services/project-management/internal/domain/models"
	"AmanahPro/services/project-management/internal/dto"
)

type ProjectRecapRepository interface {
	Create(recap *models.ProjectRecap) error
	Update(recap *models.ProjectRecap) error
	Delete(id int) error
	FindByProjectID(projectID int) (*dto.ProjectRecapDTO, error)
}
