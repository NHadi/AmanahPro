package handlers

import (
	"fmt"
	"net/http"

	"AmanahPro/services/user-management/internal/application/services"

	"github.com/gin-gonic/gin"
)

// PermissionAssignmentRequest represents the request payload for assigning permissions to menus.
type PermissionAssignmentRequest struct {
	RoleID      int                 `json:"role_id" binding:"required"`
	Permissions []MenuPermissionSet `json:"permissions" binding:"required"`
}

// MenuPermissionSet represents a menu and its combined permission string.
type MenuPermissionSet struct {
	MenuID     int    `json:"menu_id" binding:"required"`
	Permission string `json:"permission" binding:"required"` // e.g., C, CR, CRU, CRUD
}

type PermissionHandler struct {
	permissionService *services.PermissionService
}

func NewPermissionHandler(permissionService *services.PermissionService) *PermissionHandler {
	return &PermissionHandler{permissionService: permissionService}
}

// AssignPermission godoc
// @Summary Assign combined permissions to a role for multiple menus
// @Description Assign specific combined permissions (e.g., C, CR, CRUD) to a role on given menus
// @Tags Permissions
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param permission body PermissionAssignmentRequest true "Permission Assignment"
// @Success 200 {string} string "Permissions assigned successfully"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string
// @Router /api/permissions/assign [post]
func (h *PermissionHandler) AssignPermission(c *gin.Context) {
	var req PermissionAssignmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	validPermissions := map[string]bool{
		"C": true, "R": true, "U": true, "D": true,
		"CR": true, "CU": true, "RU": true, "CRUD": true,
		"RD": true, "CRU": true, "CD": true, "RUD": true, "CUD": true,
	}

	for _, menu := range req.Permissions {
		if !validPermissions[menu.Permission] {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid permission: %s", menu.Permission)})
			return
		}

		if err := h.permissionService.AssignPermission(req.RoleID, menu.MenuID, menu.Permission); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("Failed to assign permission to menu %d: %s", menu.MenuID, err.Error()),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Permissions assigned successfully"})
}
