package repositories

import (
	"AmanahPro/services/ba-services/internal/domain/models"
	"AmanahPro/services/ba-services/internal/domain/repositories"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type baProgressRepositoryImpl struct {
	db *gorm.DB
}

// NewBAProgressRepository creates a new instance of BAProgressRepository
func NewBAProgressRepository(db *gorm.DB) repositories.BAProgressRepository {
	return &baProgressRepositoryImpl{db: db}
}

func (r *baProgressRepositoryImpl) Create(progress *models.BAProgress) error {
	log.Printf("Creating BAProgress: %+v", progress)
	if err := r.db.Create(progress).Error; err != nil {
		log.Printf("Failed to create BAProgress: %v", err)
		return fmt.Errorf("failed to create BAProgress: %w", err)
	}
	return nil
}

func (r *baProgressRepositoryImpl) Update(progress *models.BAProgress) error {
	log.Printf("Updating BAProgress ID: %d", progress.BAProgressId)
	if err := r.db.Save(progress).Error; err != nil {
		log.Printf("Failed to update BAProgress ID %d: %v", progress.BAProgressId, err)
		return fmt.Errorf("failed to update BAProgress: %w", err)
	}
	return nil
}

func (r *baProgressRepositoryImpl) Delete(progressId int) error {
	log.Printf("Deleting BAProgress ID: %d", progressId)
	if err := r.db.Delete(&models.BAProgress{}, progressId).Error; err != nil {
		log.Printf("Failed to delete BAProgress ID %d: %v", progressId, err)
		return fmt.Errorf("failed to delete BAProgress: %w", err)
	}
	return nil
}

func (r *baProgressRepositoryImpl) GetByID(progressId int) (*models.BAProgress, error) {
	log.Printf("Retrieving BAProgress by ID: %d", progressId)
	var progress models.BAProgress
	if err := r.db.First(&progress, progressId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("BAProgress ID %d not found", progressId)
			return nil, nil
		}
		log.Printf("Failed to retrieve BAProgress ID %d: %v", progressId, err)
		return nil, fmt.Errorf("failed to retrieve BAProgress: %w", err)
	}
	return &progress, nil
}
