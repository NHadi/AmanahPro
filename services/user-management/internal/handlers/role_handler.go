package handlers

import (
	"encoding/json"
	"net/http"

	"AmanahPro/services/user-management/internal/application/services"
)

// RoleHandler represents the HTTP handler for role operations
type RoleHandler struct {
	roleService *services.RoleService
}

func NewRoleHandler(roleService *services.RoleService) *RoleHandler {
	return &RoleHandler{roleService: roleService}
}

// CreateRole godoc
// @Summary Create a new role
// @Param role body models.Role true "Role"
// @Success 201 {object} models.Role
// @Router /roles [post]
func (h *RoleHandler) CreateRole(w http.ResponseWriter, r *http.Request) {
	var roleData struct {
		RoleName    string `json:"role_name"`
		Description string `json:"description"`
	}
	if err := json.NewDecoder(r.Body).Decode(&roleData); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	role, err := h.roleService.CreateRole(roleData.RoleName, roleData.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(role)
}

// GetRoles godoc
// @Summary Get all roles
// @Success 200 {array} models.Role
// @Router /roles [get]
func (h *RoleHandler) GetRoles(w http.ResponseWriter, r *http.Request) {
	roles, err := h.roleService.GetRoles()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(roles)
}
