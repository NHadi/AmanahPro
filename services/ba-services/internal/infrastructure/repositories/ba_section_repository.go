package repositories

import (
	"AmanahPro/services/ba-services/internal/domain/models"
	"AmanahPro/services/ba-services/internal/domain/repositories"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type baSectionRepositoryImpl struct {
	db *gorm.DB
}

// NewBASectionRepository creates a new instance of BASectionRepository
func NewBASectionRepository(db *gorm.DB) repositories.BASectionRepository {
	return &baSectionRepositoryImpl{db: db}
}

func (r *baSectionRepositoryImpl) Create(section *models.BASection) error {
	log.Printf("Creating BASection: %+v", section)
	if err := r.db.Create(section).Error; err != nil {
		log.Printf("Failed to create BASection: %v", err)
		return fmt.Errorf("failed to create BASection: %w", err)
	}
	return nil
}

func (r *baSectionRepositoryImpl) Update(section *models.BASection) error {
	log.Printf("Updating BASection ID: %d", section.BASectionId)
	if err := r.db.Save(section).Error; err != nil {
		log.Printf("Failed to update BASection ID %d: %v", section.BASectionId, err)
		return fmt.Errorf("failed to update BASection: %w", err)
	}
	return nil
}

func (r *baSectionRepositoryImpl) Delete(sectionId int) error {
	log.Printf("Deleting BASection ID: %d", sectionId)
	if err := r.db.Delete(&models.BASection{}, sectionId).Error; err != nil {
		log.Printf("Failed to delete BASection ID %d: %v", sectionId, err)
		return fmt.Errorf("failed to delete BASection: %w", err)
	}
	return nil
}

func (r *baSectionRepositoryImpl) GetByID(sectionId int) (*models.BASection, error) {
	log.Printf("Retrieving BASection by ID: %d", sectionId)
	var section models.BASection
	if err := r.db.First(&section, sectionId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("BASection ID %d not found", sectionId)
			return nil, nil
		}
		log.Printf("Failed to retrieve BASection ID %d: %v", sectionId, err)
		return nil, fmt.Errorf("failed to retrieve BASection: %w", err)
	}
	return &section, nil
}
