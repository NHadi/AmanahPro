package repositories

import (
	"AmanahPro/services/user-management/internal/domain/models"
)

type UserRoleRepository interface {
	AssignRole(userID, roleID string) error
	FindRolesByUserID(userID string) ([]models.Role, error)
	UserHasRole(userID, roleID string) (bool, error)
	RemoveRole(userID, roleID string) error
}
