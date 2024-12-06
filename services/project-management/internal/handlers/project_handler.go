package handlers

import (
	"AmanahPro/services/project-management/internal/application/services"
	"AmanahPro/services/project-management/internal/dto"
	"net/http"
	"strconv"

	"github.com/NHadi/AmanahPro-common/helpers"
	"github.com/NHadi/AmanahPro-common/middleware"

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

// CreateProject
// @Summary Create Project
// @Description Create a new project
// @Tags Projects
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param project body dto.ProjectDTO true "Project Data"
// @Success 201 {object} map[string]interface{} "Created Project"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string
// @Router /api/projects [post]
func (h *ProjectHandler) CreateProject(c *gin.Context) {
	var projectDTO dto.ProjectDTO
	if err := c.ShouldBindJSON(&projectDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	claims, err := helpers.GetClaims(c)
	if err != nil {
		return
	}

	traceID, exists := c.Get(middleware.TraceIDHeader)
	if !exists {
		traceID = "unknown"
	}

	// Map DTO to Model
	project := projectDTO.ToModel(claims.UserID, *claims.OrganizationId)

	// Call the service to create the project
	if err := h.projectService.CreateProject(project, traceID.(string)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Map back to DTO for response
	projectDTO.ProjectID = project.ProjectID

	c.JSON(http.StatusCreated, gin.H{
		"message": "Project created successfully",
		"data":    projectDTO,
	})
}

// UpdateProject
// @Summary Update Project
// @Description Update an existing project
// @Tags Projects
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param project_id path int true "Project ID"
// @Param project body dto.ProjectDTO true "Project Data"
// @Success 200 {object} map[string]interface{} "Updated Project"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "Project Not Found"
// @Failure 500 {object} map[string]string
// @Router /api/projects/{project_id} [put]
func (h *ProjectHandler) UpdateProject(c *gin.Context) {
	idStr := c.Param("project_id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	var projectDTO dto.ProjectDTO
	if err := c.ShouldBindJSON(&projectDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	claims, err := helpers.GetClaims(c)
	if err != nil {
		return
	}

	traceID, exists := c.Get(middleware.TraceIDHeader)
	if !exists {
		traceID = "unknown"
	}

	// Retrieve the existing project from the database
	existingProject, err := h.projectService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	// Map updated data from DTO to Model
	updatedProject := projectDTO.ToModelForUpdate(existingProject, claims.UserID)
	updatedProject.OrganizationID = claims.OrganizationId

	// Call the service to update the project
	if err := h.projectService.UpdateProject(updatedProject, traceID.(string)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Project updated successfully",
		"data":    projectDTO,
	})
}

// DeleteProject
// @Summary Delete Project
// @Description Delete a project by ID
// @Tags Projects
// @Security BearerAuth
// @Param project_id path int true "Project ID"
// @Produce json
// @Success 204 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "Project Not Found"
// @Failure 500 {object} map[string]string
// @Router /api/projects/{project_id} [delete]
func (h *ProjectHandler) DeleteProject(c *gin.Context) {
	idStr := c.Param("project_id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	claims, err := helpers.GetClaims(c)
	if err != nil {
		return
	}

	traceID, exists := c.Get(middleware.TraceIDHeader)
	if !exists {
		traceID = "unknown"
	}

	// Call the service to delete the project
	if err := h.projectService.DeleteProject(id, traceID.(string), claims.UserID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
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
