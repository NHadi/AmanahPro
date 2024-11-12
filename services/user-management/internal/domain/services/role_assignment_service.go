package services

import (
	"errors"

	"AmanahPro/services/user-management/internal/domain/repositories"
)

type RoleAssignmentService struct {
	userRoleRepo repositories.UserRoleRepository
}

func NewRoleAssignmentService(userRoleRepo repositories.UserRoleRepository) *RoleAssignmentService {
	return &RoleAssignmentService{
		userRoleRepo: userRoleRepo,
	}
}

// AssignRole assigns a role to a user
func (s *RoleAssignmentService) AssignRole(userID, roleID int) error {
	// Check if the user already has the role
	exists, err := s.userRoleRepo.UserHasRole(userID, roleID)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("user already has the role")
	}

	// Assign the role
	return s.userRoleRepo.AssignRole(userID, roleID)
}

// RemoveRole removes a role from a user
func (s *RoleAssignmentService) RemoveRole(userID, roleID int) error {
	// Check if the user has the role
	exists, err := s.userRoleRepo.UserHasRole(userID, roleID)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("user does not have the role")
	}

	// Remove the role
	return s.userRoleRepo.RemoveRole(userID, roleID)
}

// UserHasRole checks if a user has a specific role
func (s *RoleAssignmentService) UserHasRole(userID, roleID int) (bool, error) {
	return s.userRoleRepo.UserHasRole(userID, roleID)
}
