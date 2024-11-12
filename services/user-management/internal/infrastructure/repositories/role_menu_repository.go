package repositories

import (
	"errors"

	"AmanahPro/services/user-management/internal/domain/models"
	"AmanahPro/services/user-management/internal/domain/repositories"

	"gorm.io/gorm"
)

type roleMenuRepository struct {
	db *gorm.DB
}

func NewRoleMenuRepository(db *gorm.DB) repositories.RoleMenuRepository {
	return &roleMenuRepository{db: db}
}

// HasPermission checks if a specific role has a specific permission on a specific menu.
func (r *roleMenuRepository) HasPermission(roleID, menuID string, permission string) (bool, error) {
	var count int64
	err := r.db.Model(&models.RoleMenu{}).
		Where("role_id = ? AND menu_id = ? AND permission = ?", roleID, menuID, permission).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *roleMenuRepository) AssignMenu(roleMenu *models.RoleMenu) error {
	return r.db.Create(roleMenu).Error
}

func (r *roleMenuRepository) FindMenusByRoleID(roleID string) ([]models.Menu, error) {
	var menus []models.Menu
	err := r.db.Joins("JOIN role_menus ON menus.menu_id = role_menus.menu_id").
		Where("role_menus.role_id = ?", roleID).
		Find(&menus).Error
	return menus, err
}

func (r *roleMenuRepository) RemoveMenu(roleID, menuID string) error {
	return r.db.Where("role_id = ? AND menu_id = ?", roleID, menuID).Delete(&models.RoleMenu{}).Error
}

// AssignPermission assigns a specific permission to a role for a menu item.
func (r *roleMenuRepository) AssignPermission(roleID, menuID string, permission string) error {
	// Validate permission value
	validPermissions := map[string]bool{"view": true, "edit": true, "delete": true}
	if !validPermissions[permission] {
		return errors.New("invalid permission")
	}

	// Check if the permission already exists
	var existingRoleMenu models.RoleMenu
	err := r.db.Where("role_id = ? AND menu_id = ? AND permission = ?", roleID, menuID, permission).First(&existingRoleMenu).Error
	if err == nil {
		// Permission already exists, no need to add it again
		return nil
	} else if err != gorm.ErrRecordNotFound {
		// An error occurred other than "record not found"
		return err
	}

	// Insert new permission
	newRoleMenu := models.RoleMenu{
		RoleID:     roleID,
		MenuID:     menuID,
		Permission: permission,
	}
	return r.db.Create(&newRoleMenu).Error
}
