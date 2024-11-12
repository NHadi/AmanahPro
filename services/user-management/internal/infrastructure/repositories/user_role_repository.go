package repositories

import (
	"AmanahPro/services/user-management/internal/domain/models"
	"AmanahPro/services/user-management/internal/domain/repositories"

	"gorm.io/gorm"
)

type userRoleRepository struct {
	db *gorm.DB
}

func NewUserRoleRepository(db *gorm.DB) repositories.UserRoleRepository {
	return &userRoleRepository{db: db}
}

// AssignRole assigns a role to a user by adding an entry to the UserRoles table.
func (r *userRoleRepository) AssignRole(userID, roleID int) error {
	userRole := &models.UserRole{
		UserID: userID,
		RoleID: roleID,
	}
	return r.db.Create(userRole).Error
}

func (r *userRoleRepository) FindRolesByUserID(userID int) ([]models.Role, error) {
	var roles []models.Role
	err := r.db.Joins("JOIN user_roles ON roles.role_id = user_roles.role_id").
		Where("user_roles.user_id = ?", userID).
		Find(&roles).Error
	return roles, err
}

func (r *userRoleRepository) RemoveRole(userID, roleID int) error {
	return r.db.Where("user_id = ? AND role_id = ?", userID, roleID).Delete(&models.UserRole{}).Error
}

func (r *userRoleRepository) UserHasRole(userID, roleID int) (bool, error) {
	var count int64
	err := r.db.Model(&models.UserRole{}).
		Where("user_id = ? AND role_id = ?", userID, roleID).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
