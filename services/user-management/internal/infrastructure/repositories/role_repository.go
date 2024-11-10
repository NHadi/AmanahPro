package repositories

import (
	"AmanahPro/services/user-management/internal/domain/models"
	"AmanahPro/services/user-management/internal/domain/repositories"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) repositories.RoleRepository {
	return &roleRepository{db: db}
}

func (r *roleRepository) Create(role *models.Role) error {
	return r.db.Create(role).Error
}

func (r *roleRepository) FindByID(id uuid.UUID) (*models.Role, error) {
	var role models.Role
	err := r.db.First(&role, "role_id = ?", id).Error
	return &role, err
}

func (r *roleRepository) FindAll() ([]models.Role, error) {
	var roles []models.Role
	err := r.db.Find(&roles).Error
	return roles, err
}

func (r *roleRepository) DeleteByID(id uuid.UUID) error {
	return r.db.Delete(&models.Role{}, "role_id = ?", id).Error
}
