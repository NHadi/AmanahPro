package repositories

import "AmanahPro/services/sph-services/internal/domain/models"

// SphRepository defines the interface for SPH repository operations
type SphRepository interface {
	// Create inserts a new SPH record into the database
	Create(sph *models.Sph) error

	// Update modifies an existing SPH record in the database
	Update(sph *models.Sph) error

	// Delete removes a SPH record from the database
	Delete(sphID int) error

	// GetByID retrieves a SPH record by its ID
	GetByID(sphID int) (*models.Sph, error)

	Filter(organizationID int, sphID *int, projectID *int) ([]models.Sph, error)

	// Transaction management
	Begin() (SphRepository, error)
	Commit() error
	Rollback() error
}
