package repositories

import (
	"AmanahPro/services/user-management/internal/domain/models"
)

type RoleRepository interface {
	Create(role *models.Role) error
	FindByID(id string) (*models.Role, error)
	FindAll() ([]models.Role, error)
	DeleteByID(id string) error
}
