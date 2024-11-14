package handlers

import (
	"AmanahPro/services/user-management/internal/application/services"
	"net/http"

	"github.com/gin-gonic/gin"
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
// @Security BearerAuth
// @Tags Roles
// @Accept json
// @Produce json
// @Param role body models.Role true "Role Data"
// @Success 201 {object} models.Role
// @Failure 400 {object} map[string]string
// @Router /api/roles [post]
func (h *RoleHandler) CreateRole(c *gin.Context) {
	var roleData struct {
		RoleName    string `json:"role_name"`
		Description string `json:"description"`
	}

	// Bind JSON input to struct
	if err := c.ShouldBindJSON(&roleData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Create role through service
	role, err := h.roleService.CreateRole(roleData.RoleName, roleData.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with created role data
	c.JSON(http.StatusCreated, role)
}

// GetRoles godoc
// @Summary Get all roles
// @Description Retrieve all roles in the system
// @Security BearerAuth
// @Tags Roles
// @Accept json
// @Produce json
// @Success 200 {array} models.Role
// @Router /api/roles [get]
func (h *RoleHandler) GetRoles(c *gin.Context) {
	roles, err := h.roleService.GetRoles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with roles data
	c.JSON(http.StatusOK, roles)
}
