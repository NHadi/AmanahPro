package repositories

import (
	"AmanahPro/services/spk-services/internal/domain/models"
	"AmanahPro/services/spk-services/internal/domain/repositories"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type spkSectionRepositoryImpl struct {
	db *gorm.DB
}

// NewSPKSectionRepository creates a new instance of SPKSectionRepository
func NewSPKSectionRepository(db *gorm.DB) repositories.SPKSectionRepository {
	return &spkSectionRepositoryImpl{
		db: db,
	}
}

// Create inserts a new SPKSection record into the database
func (r *spkSectionRepositoryImpl) Create(section *models.SPKSection) error {
	log.Printf("Creating SPKSection: %+v", section)

	if err := r.db.Create(section).Error; err != nil {
		log.Printf("Failed to create SPKSection: %v", err)
		return fmt.Errorf("failed to create SPKSection: %w", err)
	}

	log.Printf("Successfully created SPKSection: %+v", section)
	return nil
}

// Update modifies an existing SPKSection record in the database
func (r *spkSectionRepositoryImpl) Update(section *models.SPKSection) error {
	log.Printf("Updating SPKSection ID: %d", section.SectionId)

	if err := r.db.Save(section).Error; err != nil {
		log.Printf("Failed to update SPKSection ID %d: %v", section.SectionId, err)
		return fmt.Errorf("failed to update SPKSection: %w", err)
	}

	log.Printf("Successfully updated SPKSection ID: %d", section.SectionId)
	return nil
}

// Delete removes an SPKSection record from the database
func (r *spkSectionRepositoryImpl) Delete(sectionID int) error {
	log.Printf("Deleting SPKSection ID: %d", sectionID)

	if err := r.db.Delete(&models.SPKSection{}, sectionID).Error; err != nil {
		log.Printf("Failed to delete SPKSection ID %d: %v", sectionID, err)
		return fmt.Errorf("failed to delete SPKSection: %w", err)
	}

	log.Printf("Successfully deleted SPKSection ID: %d", sectionID)
	return nil
}

// GetByID retrieves an SPKSection record by its ID
func (r *spkSectionRepositoryImpl) GetByID(sectionID int) (*models.SPKSection, error) {
	log.Printf("Retrieving SPKSection by ID: %d", sectionID)

	var section models.SPKSection
	if err := r.db.First(&section, sectionID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("SPKSection ID %d not found", sectionID)
			return nil, nil
		}
		log.Printf("Failed to retrieve SPKSection ID %d: %v", sectionID, err)
		return nil, fmt.Errorf("failed to retrieve SPKSection: %w", err)
	}

	log.Printf("Successfully retrieved SPKSection: %+v", section)
	return &section, nil
}
