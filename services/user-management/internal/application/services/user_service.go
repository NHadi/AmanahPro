package services

import (
	"errors"

	"AmanahPro/services/user-management/internal/domain/models"
	"AmanahPro/services/user-management/internal/domain/repositories"
	"AmanahPro/services/user-management/internal/domain/services"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo          repositories.UserRepository
	roleAssignmentSvc *services.RoleAssignmentService
}

func NewUserService(userRepo repositories.UserRepository, roleAssignmentSvc *services.RoleAssignmentService) *UserService {
	return &UserService{
		userRepo:          userRepo,
		roleAssignmentSvc: roleAssignmentSvc,
	}
}

func (s *UserService) CreateUser(username, email, password string, organizationId *int, roleId int) (*models.User, error) {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	// Create the user model
	user := &models.User{
		Username:       username,
		Email:          email,
		Password:       string(hashedPassword),
		Status:         true,
		OrganizationID: organizationId,
	}

	// Save the user in the database
	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	// Assign the user to a role
	s.roleAssignmentSvc.AssignRole(user.UserID, roleId)

	return user, nil
}

func (s *UserService) AssignRoleToUser(userID, roleID int) error {
	return s.roleAssignmentSvc.AssignRole(userID, roleID)
}

func (s *UserService) Authenticate(email, password string) (*models.User, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid password")
	}

	return user, nil
}

func (s *UserService) LoadByOrganizationID(organizationId int) ([]models.User, error) {
	users, err := s.userRepo.FindByOrganizationID(organizationId)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// UpdateUser updates the user details and re-assigns roles if necessary (without transaction commit)
func (s *UserService) UpdateUser(userID int, username, email, password string, organizationId *int, roleID int) (*models.User, error) {
	// Fetch the existing user
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Update the user's details
	user.Username = username
	user.Email = email
	if password != "" { // If a new password is provided, hash it and update
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return nil, errors.New("failed to hash password")
		}
		user.Password = string(hashedPassword)
	}
	user.OrganizationID = organizationId

	// Save updated user details
	if err := s.userRepo.Update(user); err != nil {
		return nil, errors.New("failed to update user")
	}

	// Re-assign the role (it may or may not have changed)
	if err := s.roleAssignmentSvc.AssignRole(user.UserID, roleID); err != nil {
		return nil, err
	}

	return user, nil
}
