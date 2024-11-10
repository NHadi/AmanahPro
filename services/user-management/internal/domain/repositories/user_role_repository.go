package repositories

import (
	"AmanahPro/services/user-management/internal/domain/models"

	"github.com/google/uuid"
)

type UserRoleRepository interface {
	AssignRole(userID, roleID uuid.UUID) error
	FindRolesByUserID(userID uuid.UUID) ([]models.Role, error)
	UserHasRole(userID, roleID uuid.UUID) (bool, error)
	RemoveRole(userID, roleID uuid.UUID) error
}
