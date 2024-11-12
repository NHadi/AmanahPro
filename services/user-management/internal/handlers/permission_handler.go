package handlers

import (
	"encoding/json"
	"net/http"

	"AmanahPro/services/user-management/internal/application/services"
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
// @Accept json
// @Produce json
// @Param permission body PermissionAssignmentRequest true "Permission Assignment"
// @Success 200 {string} string "Permission assigned successfully"
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /permissions/assign [post]
func (h *PermissionHandler) AssignPermission(w http.ResponseWriter, r *http.Request) {
	var req PermissionAssignmentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Use `err` here without redeclaring it
	err := h.permissionService.AssignPermission(req.RoleID, req.MenuID, req.Permission)
	if err != nil {
		http.Error(w, "Failed to assign permission", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Permission assigned successfully"))
}
