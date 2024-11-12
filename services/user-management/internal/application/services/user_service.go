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

func (s *UserService) CreateUser(username, email, password string) (*models.User, error) {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
		Status:   true,
	}

	err = s.userRepo.Create(user)
	return user, err
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
