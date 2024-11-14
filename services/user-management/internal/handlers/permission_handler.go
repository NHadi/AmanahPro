package handlers

import (
	"net/http"

	"AmanahPro/services/user-management/internal/application/services"

	"github.com/gin-gonic/gin"
)

// PermissionAssignmentRequest represents the request body for assigning a permission
type PermissionAssignmentRequest struct {
	RoleID     int    `json:"role_id"`
	MenuID     int    `json:"menu_id"`
	Permission string `json:"permission"`
}

type PermissionHandler struct {
	permissionService *services.PermissionService
}

func NewPermissionHandler(permissionService *services.PermissionService) *PermissionHandler {
	return &PermissionHandler{permissionService: permissionService}
}

// AssignPermission godoc
// @Summary Assign permission to a role for a menu
// @Description Assign a specific permission to a role on a given menu
// @Tags Permissions
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param permission body PermissionAssignmentRequest true "Permission Assignment"
// @Success 200 {string} string "Permission assigned successfully"
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

	if err := h.permissionService.AssignPermission(req.RoleID, req.MenuID, req.Permission); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to assign permission"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Permission assigned successfully"})
}
