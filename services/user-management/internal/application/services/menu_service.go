package services

import (
	"AmanahPro/services/user-management/internal/application/dto"
	"AmanahPro/services/user-management/internal/domain/models"
	"AmanahPro/services/user-management/internal/domain/repositories"
)

type MenuService struct {
	menuRepository     repositories.MenuRepository
	roleMenuRepository repositories.RoleMenuRepository
}

func NewMenuService(menuRepo repositories.MenuRepository, roleMenuRepo repositories.RoleMenuRepository) *MenuService {
	return &MenuService{
		menuRepository:     menuRepo,
		roleMenuRepository: roleMenuRepo,
	}
}

func (s *MenuService) GetAccessibleMenus(roleID int) ([]models.Menu, error) {
	return s.menuRepository.FindByRole(roleID)
}

func (s *MenuService) GetMenusWithPermissions(roleID int) ([]dto.MenuWithPermissionDTO, error) {
	// Fetch all menus assigned to the role
	menus, err := s.menuRepository.FindByRole(roleID)
	if err != nil {
		return nil, err
	}

	// If no menus are found, return an empty slice
	if len(menus) == 0 {
		return []dto.MenuWithPermissionDTO{}, nil
	}

	// Map domain models to DTOs
	var menuWithPermissionDTOs []dto.MenuWithPermissionDTO
	for _, menu := range menus {
		permission, err := s.roleMenuRepository.GetPermissionByRoleAndMenu(roleID, menu.MenuID)
		if err != nil {
			return nil, err
		}

		menuWithPermissionDTOs = append(menuWithPermissionDTOs, dto.MenuWithPermissionDTO{
			MenuID:     menu.MenuID,
			MenuName:   menu.Name,
			Path:       menu.Path,
			Permission: permission.Permission,
		})
	}

	return menuWithPermissionDTOs, nil
}

func (s *MenuService) CreateMenu(menuName, path, icon string, order int) (*models.Menu, error) {
	menu := &models.Menu{
		Name:  menuName,
		Path:  path,
		Icon:  icon,
		Order: &order,
	}
	err := s.menuRepository.Create(menu)
	return menu, err
}
