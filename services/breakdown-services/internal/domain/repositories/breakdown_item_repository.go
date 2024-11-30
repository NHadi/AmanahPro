package repositories

import (
	"AmanahPro/services/breakdown-services/internal/domain/models"
)

type BreakdownItemRepository interface {
	// Create inserts a new BreakdownItem record into the database
	Create(item *models.BreakdownItem) error

	// Update modifies an existing BreakdownItem record in the database
	Update(item *models.BreakdownItem) error

	// Delete removes a BreakdownItem record from the database
	Delete(itemID int) error

	GetByID(id int) (*models.BreakdownItem, error)
}
