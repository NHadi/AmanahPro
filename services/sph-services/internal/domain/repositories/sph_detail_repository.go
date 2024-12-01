package repositories

import "AmanahPro/services/sph-services/internal/domain/models"

// SphDetailRepository defines the interface for SPH Detail repository operations
type SphDetailRepository interface {
	// Create inserts a new SPH Detail record into the database
	Create(detail *models.SphDetail) error

	// Update modifies an existing SPH Detail record in the database
	Update(detail *models.SphDetail) error

	// Delete removes a SPH Detail record from the database
	Delete(detailID int) error

	// GetByID retrieves a SPH Detail record by its ID
	GetByID(detailID int) (*models.SphDetail, error)
}
