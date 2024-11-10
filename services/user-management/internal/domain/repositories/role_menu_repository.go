package repositories

import (
	"AmanahPro/services/user-management/internal/domain/models"

	"github.com/google/uuid"
)

type RoleMenuRepository interface {
	AssignMenu(roleMenu *models.RoleMenu) error
	FindMenusByRoleID(roleID uuid.UUID) ([]models.Menu, error)
	RemoveMenu(roleID, menuID uuid.UUID) error
	HasPermission(roleID, menuID uuid.UUID, permission string) (bool, error)
	AssignPermission(roleID, menuID uuid.UUID, permission string) error
}
