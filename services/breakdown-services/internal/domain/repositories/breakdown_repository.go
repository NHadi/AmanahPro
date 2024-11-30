package repositories

import (
	"AmanahPro/services/breakdown-services/internal/domain/models"
)

type BreakdownRepository interface {
	GetByID(breakdownID int) (*models.Breakdown, error)
	Create(breakdown *models.Breakdown) error
	Update(breakdown *models.Breakdown) error
	Delete(breakdownID int) error
	// Get a list of breakdowns filtered by organization ID (required) and optional filters
	FilterBreakdowns(organizationID int, breakdownID *int, projectID *int) ([]models.Breakdown, error)
}
