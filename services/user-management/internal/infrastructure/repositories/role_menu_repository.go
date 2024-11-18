package repositories

import (
	"AmanahPro/services/user-management/internal/domain/models"
	"AmanahPro/services/user-management/internal/domain/repositories"
	"time"

	"gorm.io/gorm"
)

type roleMenuRepository struct {
	db *gorm.DB
}

func NewRoleMenuRepository(db *gorm.DB) repositories.RoleMenuRepository {
	return &roleMenuRepository{db: db}
}

// HasPermission checks if a specific role has a specific permission on a specific menu.
func (r *roleMenuRepository) HasPermission(roleID, menuID int, permission string) (bool, error) {
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

func (r *roleMenuRepository) FindMenusByRoleID(roleID int) ([]models.Menu, error) {
	var menus []models.Menu
	err := r.db.Joins("JOIN RoleMenus ON menus.menu_id = RoleMenus.menu_id").
		Where("RoleMenus.role_id = ?", roleID).
		Find(&menus).Error
	return menus, err
}

func (r *roleMenuRepository) RemoveMenu(roleID, menuID int) error {
	return r.db.Where("role_id = ? AND menu_id = ?", roleID, menuID).Delete(&models.RoleMenu{}).Error
}

// AssignPermission assigns a combined permission string (e.g., C, CR, CRUD) to a role for a menu item.
func (r *roleMenuRepository) AssignPermission(roleID, menuID int, permission string) error {
	var existingRoleMenu models.RoleMenu
	err := r.db.Where("role_id = ? AND menu_id = ?", roleID, menuID).First(&existingRoleMenu).Error

	if err == nil {
		// Update existing permission
		existingRoleMenu.Permission = permission
		existingRoleMenu.AssignedAt = time.Now() // Update timestamp
		return r.db.Save(&existingRoleMenu).Error
	} else if err != gorm.ErrRecordNotFound {
		// Unexpected error
		return err
	}

	// Insert new permission
	newRoleMenu := models.RoleMenu{
		RoleID:     roleID,
		MenuID:     menuID,
		Permission: permission,
		AssignedAt: time.Now(), // Explicitly set timestamp
	}
	return r.db.Create(&newRoleMenu).Error
}

// GetPermissionByRoleAndMenu fetches the permission for a given role and menu.
func (r *roleMenuRepository) GetPermissionByRoleAndMenu(roleID, menuID int) (models.RoleMenu, error) {
	var roleMenu models.RoleMenu
	err := r.db.
		Where("role_id = ? AND menu_id = ?", roleID, menuID).
		First(&roleMenu).Error
	return roleMenu, err
}
