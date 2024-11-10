package services

import (
	"AmanahPro/services/user-management/internal/domain/repositories"

	"github.com/google/uuid"
)

type PermissionService struct {
	roleMenuRepo repositories.RoleMenuRepository
}

func NewPermissionService(roleMenuRepo repositories.RoleMenuRepository) *PermissionService {
	return &PermissionService{roleMenuRepo: roleMenuRepo}
}

func (s *PermissionService) AssignPermission(roleID uuid.UUID, menuID uuid.UUID, permission string) error {
	return s.roleMenuRepo.AssignPermission(roleID, menuID, permission)
}
