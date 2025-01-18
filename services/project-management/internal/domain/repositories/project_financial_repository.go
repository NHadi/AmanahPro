package repositories

import (
	"AmanahPro/services/project-management/internal/domain/models"
	"AmanahPro/services/project-management/internal/dto"
)

type ProjectFinancialRepository interface {
	Create(financial *models.ProjectFinancial) error
	Update(financial *models.ProjectFinancial) error
	Delete(id int) error
	GetByID(id int) (*models.ProjectFinancial, error)
	GetAllByProjectID(projectID int) ([]models.ProjectFinancial, error)
	GetProjectFinancialSummary(organizationID int) ([]dto.ProjectFinancialSummaryDTO, error)
	GetProjectFinancialSPVSummary(userID int) ([]dto.ProjectFinancialSPVSummaryDTO, error)
	GetProjectFinancialSPVDetails(userID int, projectID int) ([]dto.ProjectFinancialSPVDetailDTO, error)
}
