package repositories

import (
	"AmanahPro/services/sph-services/internal/domain/models"
	"AmanahPro/services/sph-services/internal/domain/repositories"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type sphDetailRepositoryImpl struct {
	db *gorm.DB
}

// NewSphDetailRepository creates a new instance of SphDetailRepository
func NewSphDetailRepository(db *gorm.DB) repositories.SphDetailRepository {
	return &sphDetailRepositoryImpl{db: db}
}

// Create inserts a new SPH Detail record into the database
func (r *sphDetailRepositoryImpl) Create(detail *models.SphDetail) error {
	log.Printf("Creating SPH Detail: %+v", detail)

	if err := r.db.Create(detail).Error; err != nil {
		log.Printf("Failed to create SPH Detail: %v", err)
		return fmt.Errorf("failed to create SPH Detail: %w", err)
	}

	log.Printf("Successfully created SPH Detail: %+v", detail)
	return nil
}

// Update modifies an existing SPH Detail record in the database
func (r *sphDetailRepositoryImpl) Update(detail *models.SphDetail) error {
	log.Printf("Updating SPH Detail ID: %d", detail.SphDetailId)

	if err := r.db.Save(detail).Error; err != nil {
		log.Printf("Failed to update SPH Detail ID %d: %v", detail.SphDetailId, err)
		return fmt.Errorf("failed to update SPH Detail: %w", err)
	}

	log.Printf("Successfully updated SPH Detail ID: %d", detail.SphDetailId)
	return nil
}

// Delete removes a SPH Detail record from the database
func (r *sphDetailRepositoryImpl) Delete(detailID int) error {
	log.Printf("Deleting SPH Detail ID: %d", detailID)

	if err := r.db.Delete(&models.SphDetail{}, detailID).Error; err != nil {
		log.Printf("Failed to delete SPH Detail ID %d: %v", detailID, err)
		return fmt.Errorf("failed to delete SPH Detail: %w", err)
	}

	log.Printf("Successfully deleted SPH Detail ID: %d", detailID)
	return nil
}

// GetByID retrieves a SPH Detail record by its ID
func (r *sphDetailRepositoryImpl) GetByID(detailID int) (*models.SphDetail, error) {
	log.Printf("Retrieving SPH Detail by ID: %d", detailID)

	var detail models.SphDetail
	if err := r.db.First(&detail, detailID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("SPH Detail ID %d not found", detailID)
			return nil, nil
		}
		log.Printf("Failed to retrieve SPH Detail ID %d: %v", detailID, err)
		return nil, fmt.Errorf("failed to retrieve SPH Detail: %w", err)
	}

	log.Printf("Successfully retrieved SPH Detail: %+v", detail)
	return &detail, nil
}
