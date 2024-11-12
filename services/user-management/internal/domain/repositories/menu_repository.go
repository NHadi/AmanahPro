package repositories

import (
	"AmanahPro/services/user-management/internal/domain/models"
)

type MenuRepository interface {
	Create(menu *models.Menu) error
	FindByID(id int) (*models.Menu, error)
	FindByRole(roleID int) ([]models.Menu, error)
	FindAll() ([]models.Menu, error)
	DeleteByID(id int) error
}
