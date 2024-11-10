package services

import (
	"AmanahPro/services/user-management/internal/domain/repositories"

	"github.com/google/uuid"
)

type PermissionCheckerService struct {
	roleMenuRepo repositories.RoleMenuRepository
}

func NewPermissionCheckerService(roleMenuRepo repositories.RoleMenuRepository) *PermissionCheckerService {
	return &PermissionCheckerService{roleMenuRepo: roleMenuRepo}
}

func (s *PermissionCheckerService) HasPermission(roleID uuid.UUID, menuID uuid.UUID, permission string) (bool, error) {
	return s.roleMenuRepo.HasPermission(roleID, menuID, permission)
}
