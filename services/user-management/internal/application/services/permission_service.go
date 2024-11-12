package services

import (
	"AmanahPro/services/user-management/internal/domain/repositories"
)

type PermissionService struct {
	roleMenuRepo repositories.RoleMenuRepository
}

func NewPermissionService(roleMenuRepo repositories.RoleMenuRepository) *PermissionService {
	return &PermissionService{roleMenuRepo: roleMenuRepo}
}

func (s *PermissionService) AssignPermission(roleID, menuID, permission string) error {
	return s.roleMenuRepo.AssignPermission(roleID, menuID, permission)
}
