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

// NewBADetailRepository creates a new instance of BADetailRepository
func NewBADetailRepository(db *gorm.DB) repositories.BADetailRepository {
	return &baDetailRepositoryImpl{db: db}
}

func (r *baDetailRepositoryImpl) Create(detail *models.BADetail) error {
	log.Printf("Creating BADetail: %+v", detail)
	if err := r.db.Create(detail).Error; err != nil {
		log.Printf("Failed to create BADetail: %v", err)
		return fmt.Errorf("failed to create BADetail: %w", err)
	}
	return nil
}

func (r *baDetailRepositoryImpl) Update(detail *models.BADetail) error {
	log.Printf("Updating BADetail ID: %d", detail.DetailId)
	if err := r.db.Save(detail).Error; err != nil {
		log.Printf("Failed to update BADetail ID %d: %v", detail.DetailId, err)
		return fmt.Errorf("failed to update BADetail: %w", err)
	}
	return nil
}

func (r *baDetailRepositoryImpl) Delete(detailId int) error {
	log.Printf("Deleting BADetail ID: %d", detailId)
	if err := r.db.Delete(&models.BADetail{}, detailId).Error; err != nil {
		log.Printf("Failed to delete BADetail ID %d: %v", detailId, err)
		return fmt.Errorf("failed to delete BADetail: %w", err)
	}
	return nil
}

func (r *baDetailRepositoryImpl) GetByID(detailId int) (*models.BADetail, error) {
	log.Printf("Retrieving BADetail by ID: %d", detailId)
	var detail models.BADetail
	if err := r.db.First(&detail, detailId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("BADetail ID %d not found", detailId)
			return nil, nil
		}
		log.Printf("Failed to retrieve BADetail ID %d: %v", detailId, err)
		return nil, fmt.Errorf("failed to retrieve BADetail: %w", err)
	}
	return &detail, nil
}
