package repositories

import (
	"AmanahPro/services/breakdown-services/internal/domain/models"
	"AmanahPro/services/breakdown-services/internal/domain/repositories"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type breakdownItemRepositoryImpl struct {
	db *gorm.DB
}

// GetBreakdownItemsBySectionId retrieves all items for a given section ID
func (r *breakdownItemRepositoryImpl) GetBreakdownItemsBySectionId(sectionId int) ([]models.BreakdownItem, error) {
	var items []models.BreakdownItem
	if err := r.db.Where("SectionId = ?", sectionId).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

// GetByID retrieves a BreakdownItem by its ID
func (r *breakdownItemRepositoryImpl) GetByID(id int) (*models.BreakdownItem, error) {
	log.Printf("Retrieving BreakdownItem with ID: %d", id)

	var item models.BreakdownItem
	if err := r.db.First(&item, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("BreakdownItem with ID %d not found", id)
			return nil, fmt.Errorf("breakdown item not found")
		}
		log.Printf("Failed to retrieve BreakdownItem with ID %d: %v", id, err)
		return nil, fmt.Errorf("failed to retrieve breakdown item: %w", err)
	}

	log.Printf("Successfully retrieved BreakdownItem: %+v", item)
	return &item, nil
}

// NewBreakdownItemRepository creates a new instance of BreakdownItemRepository
func NewBreakdownItemRepository(db *gorm.DB) repositories.BreakdownItemRepository {
	return &breakdownItemRepositoryImpl{db: db}
}

// Create inserts a new BreakdownItem record into the database
func (r *breakdownItemRepositoryImpl) Create(item *models.BreakdownItem) error {
	log.Printf("Creating BreakdownItem: %+v", item)

	if err := r.db.Create(item).Error; err != nil {
		log.Printf("Failed to create BreakdownItem: %v", err)
		return fmt.Errorf("failed to create BreakdownItem: %w", err)
	}

	log.Printf("Successfully created BreakdownItem: %+v", item)
	return nil
}

// Update modifies an existing BreakdownItem record in the database
func (r *breakdownItemRepositoryImpl) Update(item *models.BreakdownItem) error {
	log.Printf("Updating BreakdownItem ID: %d", item.BreakdownItemId)

	if err := r.db.Save(item).Error; err != nil {
		log.Printf("Failed to update BreakdownItem ID %d: %v", item.BreakdownItemId, err)
		return fmt.Errorf("failed to update BreakdownItem: %w", err)
	}

	log.Printf("Successfully updated BreakdownItem ID: %d", item.BreakdownItemId)
	return nil
}

// Delete removes a BreakdownItem record from the database
func (r *breakdownItemRepositoryImpl) Delete(itemID int) error {
	log.Printf("Deleting BreakdownItem ID: %d", itemID)

	if err := r.db.Delete(&models.BreakdownItem{}, itemID).Error; err != nil {
		log.Printf("Failed to delete BreakdownItem ID %d: %v", itemID, err)
		return fmt.Errorf("failed to delete BreakdownItem: %w", err)
	}

	log.Printf("Successfully deleted BreakdownItem ID: %d", itemID)
	return nil
}
