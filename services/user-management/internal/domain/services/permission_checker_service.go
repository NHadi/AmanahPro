package services

import (
	"AmanahPro/services/user-management/internal/domain/repositories"
)

type PermissionCheckerService struct {
	roleMenuRepo repositories.RoleMenuRepository
}

func NewPermissionCheckerService(roleMenuRepo repositories.RoleMenuRepository) *PermissionCheckerService {
	return &PermissionCheckerService{roleMenuRepo: roleMenuRepo}
}

func (s *PermissionCheckerService) HasPermission(roleID, menuID int, permission string) (bool, error) {
	return s.roleMenuRepo.HasPermission(roleID, menuID, permission)
}
