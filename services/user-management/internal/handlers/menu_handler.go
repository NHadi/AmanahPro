package handlers

import (
	"encoding/json"
	"net/http"

	"AmanahPro/services/user-management/internal/application/services"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// MenuHandler represents the HTTP handler for menu operations
type MenuHandler struct {
	menuService *services.MenuService
}

func NewMenuHandler(menuService *services.MenuService) *MenuHandler {
	return &MenuHandler{menuService: menuService}
}

// GetAccessibleMenus godoc
// @Summary Get accessible menus by role ID
// @Param roleID path string true "Role ID"
// @Success 200 {array} models.Menu
// @Router /menus/{roleID} [get]
func (h *MenuHandler) GetAccessibleMenus(w http.ResponseWriter, r *http.Request) {
	roleID, err := uuid.Parse(mux.Vars(r)["roleID"])
	if err != nil {
		http.Error(w, "Invalid role ID", http.StatusBadRequest)
		return
	}
	menus, err := h.menuService.GetAccessibleMenus(roleID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(menus)
}

// CreateMenu godoc
// @Summary Create a new menu
// @Param menu body models.Menu true "Menu"
// @Success 201 {object} models.Menu
// @Router /menus [post]
func (h *MenuHandler) CreateMenu(w http.ResponseWriter, r *http.Request) {
	var menuData struct {
		MenuName string `json:"menu_name"`
		Path     string `json:"path"`
		Icon     string `json:"icon"`
		Order    int    `json:"order"`
	}
	if err := json.NewDecoder(r.Body).Decode(&menuData); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	menu, err := h.menuService.CreateMenu(menuData.MenuName, menuData.Path, menuData.Icon, menuData.Order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(menu)
}
