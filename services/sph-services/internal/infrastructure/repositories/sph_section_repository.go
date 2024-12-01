package repositories

import (
	"AmanahPro/services/sph-services/internal/domain/models"
	"AmanahPro/services/sph-services/internal/domain/repositories"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type sphSectionRepositoryImpl struct {
	db *gorm.DB
}

// NewSphSectionRepository creates a new instance of SphSectionRepository
func NewSphSectionRepository(db *gorm.DB) repositories.SphSectionRepository {
	return &sphSectionRepositoryImpl{db: db}
}

// Create inserts a new SPH Section record into the database
func (r *sphSectionRepositoryImpl) Create(section *models.SphSection) error {
	log.Printf("Creating SPH Section: %+v", section)

	if err := r.db.Create(section).Error; err != nil {
		log.Printf("Failed to create SPH Section: %v", err)
		return fmt.Errorf("failed to create SPH Section: %w", err)
	}

	log.Printf("Successfully created SPH Section: %+v", section)
	return nil
}

// Update modifies an existing SPH Section record in the database
func (r *sphSectionRepositoryImpl) Update(section *models.SphSection) error {
	log.Printf("Updating SPH Section ID: %d", section.SphSectionId)

	if err := r.db.Save(section).Error; err != nil {
		log.Printf("Failed to update SPH Section ID %d: %v", section.SphSectionId, err)
		return fmt.Errorf("failed to update SPH Section: %w", err)
	}

	log.Printf("Successfully updated SPH Section ID: %d", section.SphSectionId)
	return nil
}

// Delete removes a SPH Section record from the database
func (r *sphSectionRepositoryImpl) Delete(sectionID int) error {
	log.Printf("Deleting SPH Section ID: %d", sectionID)

	if err := r.db.Delete(&models.SphSection{}, sectionID).Error; err != nil {
		log.Printf("Failed to delete SPH Section ID %d: %v", sectionID, err)
		return fmt.Errorf("failed to delete SPH Section: %w", err)
	}

	log.Printf("Successfully deleted SPH Section ID: %d", sectionID)
	return nil
}

// GetByID retrieves a SPH Section record by its ID
func (r *sphSectionRepositoryImpl) GetByID(sectionID int) (*models.SphSection, error) {
	log.Printf("Retrieving SPH Section by ID: %d", sectionID)

	var section models.SphSection
	if err := r.db.Preload("Details").First(&section, sectionID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("SPH Section ID %d not found", sectionID)
			return nil, nil
		}
		log.Printf("Failed to retrieve SPH Section ID %d: %v", sectionID, err)
		return nil, fmt.Errorf("failed to retrieve SPH Section: %w", err)
	}

	log.Printf("Successfully retrieved SPH Section: %+v", section)
	return &section, nil
}
