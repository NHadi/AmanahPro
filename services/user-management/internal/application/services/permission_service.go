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

type MenuWithPermissionDTO struct {
	MenuID     int    `json:"menu_id"`
	Name       string `json:"name"`
	Path       string `json:"path"`
	Permission string `json:"permission"`
}

func (s *PermissionService) AssignPermission(roleID, menuID int, permission string) error {
	return s.roleMenuRepo.AssignPermission(roleID, menuID, permission)
}
