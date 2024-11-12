package repositories

import (
	"AmanahPro/services/user-management/internal/domain/models"
)

type UserRoleRepository interface {
	AssignRole(userID, roleID int) error
	FindRolesByUserID(userID int) ([]models.Role, error)
	UserHasRole(userID, roleID int) (bool, error)
	RemoveRole(userID, roleID int) error
}
