package services

import (
	"errors"

	"AmanahPro/services/user-management/internal/domain/models"
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
	// Remove any existing role assignments for this user
	if err := s.userRoleRepo.DeleteAllRolesByUserID(userID); err != nil {
		return errors.New("failed to remove existing role assignments")
	}

	// Create a new user role assignment with the new role
	userRole := models.UserRole{
		UserID: userID,
		RoleID: roleID,
	}

	// Insert the new user role into the database
	if err := s.userRoleRepo.CreateRoleAssignment(&userRole); err != nil {
		return errors.New("failed to assign role to user")
	}

	return nil
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
