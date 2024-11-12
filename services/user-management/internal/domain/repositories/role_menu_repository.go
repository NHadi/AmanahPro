package repositories

import (
	"AmanahPro/services/user-management/internal/domain/models"
)

type RoleMenuRepository interface {
	AssignMenu(roleMenu *models.RoleMenu) error
	FindMenusByRoleID(roleID string) ([]models.Menu, error)
	RemoveMenu(roleID, menuID string) error
	HasPermission(roleID, menuID string, permission string) (bool, error)
	AssignPermission(roleID, menuID string, permission string) error
}
