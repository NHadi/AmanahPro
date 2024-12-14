package repositories

import (
	"AmanahPro/services/breakdown-services/internal/domain/models"
	"AmanahPro/services/breakdown-services/internal/domain/repositories"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type mstBreakdownSectionRepositoryImpl struct {
	db *gorm.DB
}

// NewMstBreakdownSectionRepository creates a new instance of MstBreakdownSectionRepository
func NewMstBreakdownSectionRepository(db *gorm.DB) repositories.MstBreakdownSectionRepository {
	return &mstBreakdownSectionRepositoryImpl{db: db}
}

// Create inserts a new MstBreakdownSection record into the database
func (r *mstBreakdownSectionRepositoryImpl) Create(section *models.MstBreakdownSection) error {
	log.Printf("Creating MstBreakdownSection: %+v", section)

	if err := r.db.Create(section).Error; err != nil {
		log.Printf("Failed to create MstBreakdownSection: %v", err)
		return fmt.Errorf("failed to create MstBreakdownSection: %w", err)
	}

	log.Printf("Successfully created MstBreakdownSection: %+v", section)
	return nil
}

// Update modifies an existing MstBreakdownSection record in the database
func (r *mstBreakdownSectionRepositoryImpl) Update(section *models.MstBreakdownSection) error {
	log.Printf("Updating MstBreakdownSection ID: %d", section.MstBreakdownSectionId)

	if err := r.db.Save(section).Error; err != nil {
		log.Printf("Failed to update MstBreakdownSection ID %d: %v", section.MstBreakdownSectionId, err)
		return fmt.Errorf("failed to update MstBreakdownSection: %w", err)
	}

	log.Printf("Successfully updated MstBreakdownSection ID: %d", section.MstBreakdownSectionId)
	return nil
}

// Delete removes a MstBreakdownSection record from the database
func (r *mstBreakdownSectionRepositoryImpl) Delete(sectionID int) error {
	log.Printf("Deleting MstBreakdownSection ID: %d", sectionID)

	if err := r.db.Delete(&models.MstBreakdownSection{}, sectionID).Error; err != nil {
		log.Printf("Failed to delete MstBreakdownSection ID %d: %v", sectionID, err)
		return fmt.Errorf("failed to delete MstBreakdownSection: %w", err)
	}

	log.Printf("Successfully deleted MstBreakdownSection ID: %d", sectionID)
	return nil
}

// FilterBreakdowns retrieves all MstBreakdownSections for a specific organization ID, including the associated items ordered by Sort Asc
func (r *mstBreakdownSectionRepositoryImpl) FilterBreakdowns(organizationID *int) ([]models.MstBreakdownSection, error) {
	log.Printf("Filtering MstBreakdownSections for OrganizationID: %v", organizationID)

	var sections []models.MstBreakdownSection

	// Use Preload to include Items and order both sections and items by Sort ascending
	query := r.db.Preload("Items", func(db *gorm.DB) *gorm.DB {
		return db.Order("Sort ASC") // Order items by Sort ascending
	}).Order("Sort ASC") // Order sections by Sort ascending

	// Apply organization ID filter if provided
	if organizationID != nil {
		query = query.Where("OrganizationId = ?", *organizationID)
	}

	// Execute the query
	if err := query.Find(&sections).Error; err != nil {
		log.Printf("Failed to filter MstBreakdownSections: %v", err)
		return nil, fmt.Errorf("failed to filter MstBreakdownSections: %w", err)
	}

	log.Printf("Successfully retrieved MstBreakdownSections: %+v", sections)
	return sections, nil
}

func (r *mstBreakdownSectionRepositoryImpl) GetByID(id int) (*models.MstBreakdownSection, error) {
	log.Printf("Fetching MstBreakdownSection by ID: %d", id)

	var item models.MstBreakdownSection
	if err := r.db.Preload("Items").First(&item, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("MstBreakdownSection not found for ID: %d", id)
			return nil, fmt.Errorf("MstBreakdownSection not found")
		}
		log.Printf("Error fetching MstBreakdownSection by ID %d: %v", id, err)
		return nil, fmt.Errorf("error fetching MstBreakdownSection: %w", err)
	}

	log.Printf("Successfully fetched MstBreakdownSection: %+v", item)
	return &item, nil
}
