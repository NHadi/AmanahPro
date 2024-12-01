package repositories

import "AmanahPro/services/sph-services/internal/domain/models"

// SphSectionRepository defines the interface for SPH Section repository operations
type SphSectionRepository interface {
	// Create inserts a new SPH Section record into the database
	Create(section *models.SphSection) error

	// Update modifies an existing SPH Section record in the database
	Update(section *models.SphSection) error

	// Delete removes a SPH Section record from the database
	Delete(sectionID int) error

	// GetByID retrieves a SPH Section record by its ID
	GetByID(sectionID int) (*models.SphSection, error)
}
