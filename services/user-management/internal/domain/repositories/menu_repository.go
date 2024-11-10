package repositories

import (
	"AmanahPro/services/user-management/internal/domain/models"

	"github.com/google/uuid"
)

type MenuRepository interface {
	Create(menu *models.Menu) error
	FindByID(id uuid.UUID) (*models.Menu, error)
	FindByRole(roleID uuid.UUID) ([]models.Menu, error)
	FindAll() ([]models.Menu, error)
	DeleteByID(id uuid.UUID) error
}
