package repositories

import (
	"AmanahPro/services/breakdown-services/internal/domain/models"
)

type BreakdownSectionRepository interface {
	// Create inserts a new BreakdownSection record into the database
	Create(section *models.BreakdownSection) error

	// Update modifies an existing BreakdownSection record in the database
	Update(section *models.BreakdownSection) error

	// Delete removes a BreakdownSection record from the database
	Delete(sectionID int) error
}
