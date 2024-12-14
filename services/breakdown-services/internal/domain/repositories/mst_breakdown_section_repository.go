package repositories

import (
	"AmanahPro/services/breakdown-services/internal/domain/models"
)

type MstBreakdownSectionRepository interface {
	// Create inserts a new BreakdownSection record into the database
	Create(section *models.MstBreakdownSection) error

	// Update modifies an existing BreakdownSection record in the database
	Update(section *models.MstBreakdownSection) error

	// Delete removes a BreakdownSection record from the database
	Delete(sectionID int) error
	FilterBreakdowns(organizationID *int) ([]models.MstBreakdownSection, error)
	GetByID(id int) (*models.MstBreakdownSection, error)
}
