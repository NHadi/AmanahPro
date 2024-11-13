package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"AmanahPro/services/user-management/internal/application/services"

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
// @Security BearerAuth
// @Tags Menu
// @Param roleID path int true "Role ID"
// @Success 200 {array} models.Menu
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /api/menus/{roleID} [get]
func (h *MenuHandler) GetAccessibleMenus(w http.ResponseWriter, r *http.Request) {
	// Parse roleID as an integer from the URL path
	roleIDStr := mux.Vars(r)["roleID"]
	roleID, err := strconv.Atoi(roleIDStr)
	if err != nil {
		http.Error(w, "Invalid role ID", http.StatusBadRequest)
		return
	}

	// Fetch accessible menus by roleID
	menus, err := h.menuService.GetAccessibleMenus(roleID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode the response in JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(menus)
}

// CreateMenu godoc
// @Summary Create a new menu
// @Description Create a new menu entry
// @Tags Menus
// @Accept json
// @Produce json
// @Param menu body models.Menu true "Menu Data"
// @Success 201 {object} models.Menu
// @Failure 400 {object} map[string]string
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
