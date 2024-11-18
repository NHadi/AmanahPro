package dto

// MenuWithPermissionDTO represents a data structure for menu details with permissions.
type MenuWithPermissionDTO struct {
	MenuID     int    `json:"menu_id"`    // The ID of the menu
	MenuName   string `json:"menu_name"`  // The name of the menu
	Path       string `json:"path"`       // The path of the menu
	Permission string `json:"permission"` // The permission string (e.g., CRUD, R, D)
}
