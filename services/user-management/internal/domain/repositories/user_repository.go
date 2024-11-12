package repositories

import (
	"AmanahPro/services/user-management/internal/domain/models"
)

type UserRepository interface {
	Create(user *models.User) error
	FindByID(id int) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindAll() ([]models.User, error)
	DeleteByID(id int) error
}
