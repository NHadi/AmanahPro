package handlers

import (
	"AmanahPro/services/project-management/common/helpers"
	"AmanahPro/services/project-management/internal/application/services"
	"AmanahPro/services/project-management/internal/domain/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProjectHandler struct {
	projectService *services.ProjectService
}

// NewProjectHandler creates a new ProjectHandler instance
func NewProjectHandler(projectService *services.ProjectService) *ProjectHandler {
	return &ProjectHandler{
		projectService: projectService,
	}
}

// CreateProject handles the request to create a new project
// @Summary Create Project
// @Description Create a new project
// @Tags Projects
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param project body models.Project true "Project Data"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string
// @Router /api/projects [post]
func (h *ProjectHandler) CreateProject(c *gin.Context) {
	var project models.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	claims, err := helpers.GetClaims(c)
	if err != nil {
		// Error already handled inside the helper
		return
	}

	organizationID := int(*claims.OrganizationId)
	project.CreatedBy = &claims.UserID
	project.OrganizationID = &organizationID

	if err := h.projectService.Create(&project); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Project created successfully"})
}

// UpdateProject handles the request to update an existing project
// @Summary Update Project
// @Description Update an existing project
// @Tags Projects
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param project body models.Project true "Project Data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string
// @Router /api/projects [put]
func (h *ProjectHandler) UpdateProject(c *gin.Context) {
	var project models.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	claims, err := helpers.GetClaims(c)
	if err != nil {
		// Error already handled inside the helper
		return
	}

	organizationID := int(*claims.OrganizationId)
	project.OrganizationID = &organizationID
	project.UpdatedBy = &claims.UserID

	if err := h.projectService.Update(&project); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project updated successfully"})
}

// DeleteProject handles the request to delete a project
// @Summary Delete Project
// @Description Delete a project by ID
// @Tags Projects
// @Security BearerAuth
// @Param id path int true "Project ID"
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string
// @Router /api/projects/{id} [delete]
func (h *ProjectHandler) DeleteProject(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	if err := h.projectService.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project deleted successfully"})
}

// SearchProjectsByOrganization handles the request to search projects
// @Summary Search Projects
// @Description Search projects by organization ID and query string
// @Tags Projects
// @Security BearerAuth
// @Param organization_id query int true "Organization ID"
// @Param query query string true "Search Query"
// @Produce json
// @Success 200 {array} dto.ProjectDTO
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string
// @Router /api/projects/search [get]
func (h *ProjectHandler) SearchProjectsByOrganization(c *gin.Context) {

	claims, err := helpers.GetClaims(c)
	if err != nil {
		// Error already handled inside the helper
		return
	}

	organizationID := int(*claims.OrganizationId)

	query := c.Query("query")

	projects, err := h.projectService.SearchProjectsByOrganization(organizationID, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, projects)
}
