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
// @Description Create a new role with provided details
// @Tags Roles
// @Accept json
// @Produce json
// @Param role body models.Role true "Role Data"
// @Success 201 {object} models.Role
// @Failure 400 {object} map[string]string
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
// @Description Retrieve all roles in the system
// @Tags Roles
// @Accept json
// @Produce json
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
