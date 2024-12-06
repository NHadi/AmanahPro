package repositories

import (
	"AmanahPro/services/ba-services/internal/domain/models"
	"AmanahPro/services/ba-services/internal/domain/repositories"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type baDetailRepositoryImpl struct {
	db *gorm.DB
}

// NewSPKDetailRepository creates a new instance of SPKDetailRepository
func NewSPKDetailRepository(db *gorm.DB) repositories.SPKDetailRepository {
	return &baDetailRepositoryImpl{
		db: db,
	}
}

// Create inserts a new SPKDetail record into the database
func (r *baDetailRepositoryImpl) Create(detail *models.SPKDetail) error {
	log.Printf("Creating SPKDetail: %+v", detail)

	if err := r.db.Create(detail).Error; err != nil {
		log.Printf("Failed to create SPKDetail: %v", err)
		return fmt.Errorf("failed to create SPKDetail: %w", err)
	}

	log.Printf("Successfully created SPKDetail: %+v", detail)
	return nil
}

// Update modifies an existing SPKDetail record in the database
func (r *baDetailRepositoryImpl) Update(detail *models.SPKDetail) error {
	log.Printf("Updating SPKDetail ID: %d", detail.DetailId)

	if err := r.db.Save(detail).Error; err != nil {
		log.Printf("Failed to update SPKDetail ID %d: %v", detail.DetailId, err)
		return fmt.Errorf("failed to update SPKDetail: %w", err)
	}

	log.Printf("Successfully updated SPKDetail ID: %d", detail.DetailId)
	return nil
}

// Delete removes a SPKDetail record from the database
func (r *baDetailRepositoryImpl) Delete(detailId int) error {
	log.Printf("Deleting SPKDetail ID: %d", detailId)

	if err := r.db.Delete(&models.SPKDetail{}, detailId).Error; err != nil {
		log.Printf("Failed to delete SPKDetail ID %d: %v", detailId, err)
		return fmt.Errorf("failed to delete SPKDetail: %w", err)
	}

	log.Printf("Successfully deleted SPKDetail ID: %d", detailId)
	return nil
}

// GetByID retrieves a SPKDetail record by its ID
func (r *baDetailRepositoryImpl) GetByID(detailId int) (*models.SPKDetail, error) {
	log.Printf("Retrieving SPKDetail by ID: %d", detailId)

	var detail models.SPKDetail
	if err := r.db.First(&detail, detailId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("SPKDetail ID %d not found", detailId)
			return nil, nil
		}
		log.Printf("Failed to retrieve SPKDetail ID %d: %v", detailId, err)
		return nil, fmt.Errorf("failed to retrieve SPKDetail: %w", err)
	}

	log.Printf("Successfully retrieved SPKDetail: %+v", detail)
	return &detail, nil
}
