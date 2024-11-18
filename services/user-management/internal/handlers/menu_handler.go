package handlers

import (
	"net/http"
	"strconv"

	"AmanahPro/services/user-management/internal/application/services"

	"github.com/gin-gonic/gin"
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
// @Success 200 {array} dto.MenuWithPermissionDTO
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /api/menus/{roleID} [get]
func (h *MenuHandler) GetAccessibleMenus(c *gin.Context) {
	// Parse roleID as an integer from the URL path
	roleIDStr := c.Param("roleID")
	roleID, err := strconv.Atoi(roleIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	// Fetch accessible menus by roleID
	menus, err := h.menuService.GetMenusWithPermissions(roleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Encode the response in JSON
	c.JSON(http.StatusOK, menus)
}

// CreateMenu godoc
// @Summary Create a new menu
// @Description Create a new menu entry
// @Security BearerAuth
// @Tags Menus
// @Accept json
// @Produce json
// @Param menu body models.Menu true "Menu Data"
// @Success 201 {object} models.Menu
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 400 {object} map[string]string
// @Router /api/menus [post]
func (h *MenuHandler) CreateMenu(c *gin.Context) {
	var menuData struct {
		MenuName string `json:"menu_name"`
		Path     string `json:"path"`
		Icon     string `json:"icon"`
		Order    int    `json:"order"`
	}
	if err := c.ShouldBindJSON(&menuData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	menu, err := h.menuService.CreateMenu(menuData.MenuName, menuData.Path, menuData.Icon, menuData.Order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, menu)
}
