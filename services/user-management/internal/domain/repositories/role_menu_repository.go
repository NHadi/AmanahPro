package repositories

import (
	"AmanahPro/services/user-management/internal/domain/models"
)

type RoleMenuRepository interface {
	AssignMenu(roleMenu *models.RoleMenu) error
	FindMenusByRoleID(roleID int) ([]models.Menu, error)
	RemoveMenu(roleID, menuID int) error
	HasPermission(roleID, menuID int, permission string) (bool, error)
	AssignPermission(roleID, menuID int, permission string) error
}
