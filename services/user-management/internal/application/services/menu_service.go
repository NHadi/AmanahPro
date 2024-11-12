package services

import (
	"AmanahPro/services/user-management/internal/domain/models"
	"AmanahPro/services/user-management/internal/domain/repositories"
)

type MenuService struct {
	menuRepo repositories.MenuRepository
}

func NewMenuService(menuRepo repositories.MenuRepository) *MenuService {
	return &MenuService{menuRepo: menuRepo}
}

func (s *MenuService) GetAccessibleMenus(roleID int) ([]models.Menu, error) {
	return s.menuRepo.FindByRole(roleID)
}

func (s *MenuService) CreateMenu(menuName, path, icon string, order int) (*models.Menu, error) {
	menu := &models.Menu{
		MenuName: menuName,
		Path:     path,
		Icon:     icon,
		Order:    order,
	}
	err := s.menuRepo.Create(menu)
	return menu, err
}
