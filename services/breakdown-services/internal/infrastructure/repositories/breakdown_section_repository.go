package repositories

import (
	"AmanahPro/services/breakdown-services/internal/domain/models"
	"AmanahPro/services/breakdown-services/internal/domain/repositories"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type breakdownSectionRepositoryImpl struct {
	db *gorm.DB
}

// NewBreakdownSectionRepository creates a new instance of BreakdownSectionRepository
func NewBreakdownSectionRepository(db *gorm.DB) repositories.BreakdownSectionRepository {
	return &breakdownSectionRepositoryImpl{db: db}
}

// GetBreakdownSectionsByBreakdownId retrieves all breakdown sections for a specific breakdown ID
func (r *breakdownSectionRepositoryImpl) GetBreakdownSectionsByBreakdownId(breakdownId int) ([]models.BreakdownSection, error) {
	var sections []models.BreakdownSection
	if err := r.db.Preload("Items").Where("BreakdownId = ?", breakdownId).Find(&sections).Error; err != nil {
		return nil, err
	}
	return sections, nil
}

// GetByIDAndBreakdownID retrieves a BreakdownSection by its ID and associated Breakdown ID
func (r *breakdownSectionRepositoryImpl) GetByIDAndBreakdownID(sectionID int, breakdownID int) (*models.BreakdownSection, error) {
	log.Printf("Querying BreakdownSection with SectionID: %d and BreakdownID: %d", sectionID, breakdownID)

	var section models.BreakdownSection

	// Query the database for the BreakdownSection
	err := r.db.Where("[BreakdownSectionId] = ? AND [BreakdownId] = ?", sectionID, breakdownID).First(&section).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("BreakdownSection not found for SectionID: %d and BreakdownID: %d", sectionID, breakdownID)
			return nil, fmt.Errorf("breakdown section not found")
		}
		log.Printf("Error querying BreakdownSection: %v", err)
		return nil, fmt.Errorf("error querying BreakdownSection: %w", err)
	}

	log.Printf("Successfully retrieved BreakdownSection: %+v", section)
	return &section, nil
}

// Create inserts a new BreakdownSection record into the database
func (r *breakdownSectionRepositoryImpl) Create(section *models.BreakdownSection) error {
	log.Printf("Creating BreakdownSection: %+v", section)

	if err := r.db.Create(section).Error; err != nil {
		log.Printf("Failed to create BreakdownSection: %v", err)
		return fmt.Errorf("failed to create BreakdownSection: %w", err)
	}

	log.Printf("Successfully created BreakdownSection: %+v", section)
	return nil
}

// Update modifies an existing BreakdownSection record in the database
func (r *breakdownSectionRepositoryImpl) Update(section *models.BreakdownSection) error {
	log.Printf("Updating BreakdownSection ID: %d", section.BreakdownSectionId)

	if err := r.db.Save(section).Error; err != nil {
		log.Printf("Failed to update BreakdownSection ID %d: %v", section.BreakdownSectionId, err)
		return fmt.Errorf("failed to update BreakdownSection: %w", err)
	}

	log.Printf("Successfully updated BreakdownSection ID: %d", section.BreakdownSectionId)
	return nil
}

// Delete removes a BreakdownSection record from the database
func (r *breakdownSectionRepositoryImpl) Delete(sectionID int) error {
	log.Printf("Deleting BreakdownSection ID: %d", sectionID)

	if err := r.db.Delete(&models.BreakdownSection{}, sectionID).Error; err != nil {
		log.Printf("Failed to delete BreakdownSection ID %d: %v", sectionID, err)
		return fmt.Errorf("failed to delete BreakdownSection: %w", err)
	}

	log.Printf("Successfully deleted BreakdownSection ID: %d", sectionID)
	return nil
}
