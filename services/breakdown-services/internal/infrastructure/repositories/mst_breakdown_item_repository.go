package repositories

import (
	"AmanahPro/services/breakdown-services/internal/domain/models"
	"AmanahPro/services/breakdown-services/internal/domain/repositories"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type mstBreakdownItemRepositoryImpl struct {
	db *gorm.DB
}

// NewMstBreakdownItemRepository creates a new instance of MstBreakdownItemRepository
func NewMstBreakdownItemRepository(db *gorm.DB) repositories.MstBreakdownItemRepository {
	return &mstBreakdownItemRepositoryImpl{db: db}
}

// GetMstBreakdownItemsBySectionId retrieves all master breakdown items for a given section ID
func (r *mstBreakdownItemRepositoryImpl) GetMstBreakdownItemsBySectionId(sectionId int) ([]models.MstBreakdownItem, error) {
	var items []models.MstBreakdownItem
	if err := r.db.Where("MstBreakdownSectionId = ?", sectionId).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

// Create inserts a new MstBreakdownItem record into the database
func (r *mstBreakdownItemRepositoryImpl) Create(item *models.MstBreakdownItem) error {
	log.Printf("Creating MstBreakdownItem: %+v", item)

	if err := r.db.Create(item).Error; err != nil {
		log.Printf("Failed to create MstBreakdownItem: %v", err)
		return fmt.Errorf("failed to create MstBreakdownItem: %w", err)
	}

	log.Printf("Successfully created MstBreakdownItem: %+v", item)
	return nil
}

// Update modifies an existing MstBreakdownItem record in the database
func (r *mstBreakdownItemRepositoryImpl) Update(item *models.MstBreakdownItem) error {
	log.Printf("Updating MstBreakdownItem ID: %d", item.MstBreakdownItemId)

	if err := r.db.Save(item).Error; err != nil {
		log.Printf("Failed to update MstBreakdownItem ID %d: %v", item.MstBreakdownItemId, err)
		return fmt.Errorf("failed to update MstBreakdownItem: %w", err)
	}

	log.Printf("Successfully updated MstBreakdownItem ID: %d", item.MstBreakdownItemId)
	return nil
}

// Delete removes a MstBreakdownItem record from the database
func (r *mstBreakdownItemRepositoryImpl) Delete(itemID int) error {
	log.Printf("Deleting MstBreakdownItem ID: %d", itemID)

	if err := r.db.Delete(&models.MstBreakdownItem{}, itemID).Error; err != nil {
		log.Printf("Failed to delete MstBreakdownItem ID %d: %v", itemID, err)
		return fmt.Errorf("failed to delete MstBreakdownItem: %w", err)
	}

	log.Printf("Successfully deleted MstBreakdownItem ID: %d", itemID)
	return nil
}

// GetByID retrieves a MstBreakdownItem by its ID
func (r *mstBreakdownItemRepositoryImpl) GetByID(id int) (*models.MstBreakdownItem, error) {
	log.Printf("Fetching MstBreakdownItem by ID: %d", id)

	var item models.MstBreakdownItem
	if err := r.db.First(&item, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("MstBreakdownItem not found for ID: %d", id)
			return nil, fmt.Errorf("MstBreakdownItem not found")
		}
		log.Printf("Error fetching MstBreakdownItem by ID %d: %v", id, err)
		return nil, fmt.Errorf("error fetching MstBreakdownItem: %w", err)
	}

	log.Printf("Successfully fetched MstBreakdownItem: %+v", item)
	return &item, nil
}
