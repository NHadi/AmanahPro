package repositories

import (
	"AmanahPro/services/user-management/internal/domain/models"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(user *models.User) error
	FindByID(id uuid.UUID) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindAll() ([]models.User, error)
	DeleteByID(id uuid.UUID) error
}
