package repositories

import (
	"AmanahPro/services/breakdown-services/internal/domain/models"
)

type MstBreakdownItemRepository interface {
	// Create inserts a new BreakdownItem record into the database
	Create(item *models.MstBreakdownItem) error

	// Update modifies an existing BreakdownItem record in the database
	Update(item *models.MstBreakdownItem) error

	// Delete removes a BreakdownItem record from the database
	Delete(itemID int) error

	GetByID(id int) (*models.MstBreakdownItem, error)
	GetMstBreakdownItemsBySectionId(sectionId int) ([]models.MstBreakdownItem, error)
}
