package repositories

import (
	"AmanahPro/services/user-management/internal/domain/models"

	"github.com/google/uuid"
)

type RoleRepository interface {
	Create(role *models.Role) error
	FindByID(id uuid.UUID) (*models.Role, error)
	FindAll() ([]models.Role, error)
	DeleteByID(id uuid.UUID) error
}
