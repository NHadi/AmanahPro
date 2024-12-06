package repositories

import (
	"AmanahPro/services/project-management/internal/domain/models"
)

type ProjectRecapRepository interface {
	Create(recap *models.ProjectRecap) error
	Update(recap *models.ProjectRecap) error
	Delete(id int) error
	GetByID(id int) (*models.ProjectRecap, error)
}
