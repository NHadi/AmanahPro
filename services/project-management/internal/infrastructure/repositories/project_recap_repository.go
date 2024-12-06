package repositories

import (
	"AmanahPro/services/project-management/internal/domain/models"
	"AmanahPro/services/project-management/internal/domain/repositories"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type projectRecapRepositoryImpl struct {
	db *gorm.DB
}

// NewProjectRecapRepository creates a new instance of ProjectRecapRepository
func NewProjectRecapRepository(db *gorm.DB) repositories.ProjectRecapRepository {
	return &projectRecapRepositoryImpl{
		db: db,
	}
}

// Create inserts a new ProjectRecap record into the database
func (r *projectRecapRepositoryImpl) Create(recap *models.ProjectRecap) error {
	log.Printf("Creating ProjectRecap: %+v", recap)

	if err := r.db.Create(recap).Error; err != nil {
		log.Printf("Failed to create ProjectRecap: %v", err)
		return fmt.Errorf("failed to create ProjectRecap: %w", err)
	}

	log.Printf("Successfully created ProjectRecap: %+v", recap)
	return nil
}

// Update modifies an existing ProjectRecap record in the database
func (r *projectRecapRepositoryImpl) Update(recap *models.ProjectRecap) error {
	log.Printf("Updating ProjectRecap ID: %d", recap.ID)

	if err := r.db.Save(recap).Error; err != nil {
		log.Printf("Failed to update ProjectRecap ID %d: %v", recap.ID, err)
		return fmt.Errorf("failed to update ProjectRecap: %w", err)
	}

	log.Printf("Successfully updated ProjectRecap ID: %d", recap.ID)
	return nil
}

// Delete removes a ProjectRecap record from the database
func (r *projectRecapRepositoryImpl) Delete(recapID int) error {
	log.Printf("Deleting ProjectRecap ID: %d", recapID)

	if err := r.db.Delete(&models.ProjectRecap{}, recapID).Error; err != nil {
		log.Printf("Failed to delete ProjectRecap ID %d: %v", recapID, err)
		return fmt.Errorf("failed to delete ProjectRecap: %w", err)
	}

	log.Printf("Successfully deleted ProjectRecap ID: %d", recapID)
	return nil
}

// GetByID retrieves a ProjectRecap record by its ID
func (r *projectRecapRepositoryImpl) GetByID(recapID int) (*models.ProjectRecap, error) {
	log.Printf("Retrieving ProjectRecap by ID: %d", recapID)

	var recap models.ProjectRecap
	if err := r.db.First(&recap, recapID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("ProjectRecap ID %d not found", recapID)
			return nil, nil
		}
		log.Printf("Failed to retrieve ProjectRecap ID %d: %v", recapID, err)
		return nil, fmt.Errorf("failed to retrieve ProjectRecap: %w", err)
	}

	log.Printf("Successfully retrieved ProjectRecap: %+v", recap)
	return &recap, nil
}
