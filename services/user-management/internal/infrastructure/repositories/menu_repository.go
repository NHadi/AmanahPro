package repositories

import (
	"AmanahPro/services/user-management/internal/domain/models"
	"AmanahPro/services/user-management/internal/domain/repositories"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type menuRepository struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) repositories.MenuRepository {
	return &menuRepository{db: db}
}

func (r *menuRepository) Create(menu *models.Menu) error {
	return r.db.Create(menu).Error
}

func (r *menuRepository) FindByID(id uuid.UUID) (*models.Menu, error) {
	var menu models.Menu
	err := r.db.First(&menu, "menu_id = ?", id).Error
	return &menu, err
}

// FindByRole fetches accessible menus based on the roleID.
func (r *menuRepository) FindByRole(roleID uuid.UUID) ([]models.Menu, error) {
	var menus []models.Menu
	err := r.db.
		Table("Menus").
		Joins("JOIN RoleMenus ON RoleMenus.menu_id = Menus.menu_id").
		Where("RoleMenus.role_id = ?", roleID).
		Select("DISTINCT Menus.*").
		Find(&menus).Error
	return menus, err
}

func (r *menuRepository) FindAll() ([]models.Menu, error) {
	var menus []models.Menu
	err := r.db.Find(&menus).Error
	return menus, err
}

func (r *menuRepository) DeleteByID(id uuid.UUID) error {
	return r.db.Delete(&models.Menu{}, "menu_id = ?", id).Error
}
