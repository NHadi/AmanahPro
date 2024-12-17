package repositories

import (
	"AmanahPro/services/project-management/internal/domain/models"
)

type ProjectFinancialRepository interface {
	Create(financial *models.ProjectFinancial) error
	Update(financial *models.ProjectFinancial) error
	Delete(id int) error
	GetByID(id int) (*models.ProjectFinancial, error)
	GetAllByProjectID(projectID int) ([]models.ProjectFinancial, error)
}
