package services

import (
	"AmanahPro/services/user-management/internal/domain/models"
	"AmanahPro/services/user-management/internal/domain/repositories"
)

type RoleService struct {
	roleRepo repositories.RoleRepository
}

func NewRoleService(roleRepo repositories.RoleRepository) *RoleService {
	return &RoleService{roleRepo: roleRepo}
}

func (s *RoleService) CreateRole(roleName, description string) (*models.Role, error) {
	role := &models.Role{
		Name:        roleName,
		Description: description,
	}
	err := s.roleRepo.Create(role)
	return role, err
}

func (s *RoleService) GetRoles() ([]models.Role, error) {
	return s.roleRepo.FindAll()
}
