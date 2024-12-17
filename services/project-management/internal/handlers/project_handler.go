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

// SearchProjectUsersByProject handles the request to search projects
// @Summary Search Project Users
// @Description Search project Users by Project ID
// @Tags Projects
// @Security BearerAuth
// @Param project_id query int true "project_id ID"
// @Produce json
// @Success 200 {array} models.ProjectUser
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string
// @Router /api/projects/{project_id}/users [get]
func (h *ProjectHandler) SearchProjectUsersByProject(c *gin.Context) {
	idStr := c.Param("project_id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}
	claims, err := helpers.GetClaims(c)
	if err != nil {
		// Error already handled inside the helper
		return
	}

	organizationID := int(*claims.OrganizationId)

	projects, err := h.projectService.SearchProjectUsersByProject(id, organizationID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, projects)
}

// AssignUser
// @Summary Assign User to Project
// @Description Assign a user (with or without UserID) to a project
// @Tags Projects
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param project_id path int true "Project ID"
// @Param assignUser body dto.ProjectUserDTO true "User Data"
// @Success 200 {object} map[string]interface{} "User assigned successfully"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string
// @Router /api/projects/{project_id}/assign-user [post]
func (h *ProjectHandler) AssignUser(c *gin.Context) {
	idStr := c.Param("project_id")
	projectID, err := strconv.Atoi(idStr)
	if err != nil || projectID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	var userDTO dto.ProjectUserDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	claims, err := helpers.GetClaims(c)
	if err != nil {
		return
	}

	organizationID := int(*claims.OrganizationId)

	if err := h.projectService.AssignUser(projectID, userDTO.UserID, userDTO.UserName, userDTO.Role, claims.UserID, organizationID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User assigned successfully"})
}

// UnAssignUser
// @Summary Unassign User from Project
// @Description Remove a user from a project by UserName
// @Tags Projects
// @Security BearerAuth
// @Param project_id path int true "Project ID"
// @Param user_name query string true "User Name to Unassign"
// @Produce json
// @Success 200 {object} map[string]interface{} "User unassigned successfully"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string
// @Router /api/projects/{project_id}/unassign-user [delete]
func (h *ProjectHandler) UnAssignUser(c *gin.Context) {
	idStr := c.Param("project_id")
	projectID, err := strconv.Atoi(idStr)
	if err != nil || projectID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	// Parse user ID from the query parameter
	userIDStr := c.Query("user_name")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil || userID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := h.projectService.UnAssignUser(projectID, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User unassigned successfully"})
}

// ChangeUser
// @Summary Change Assigned User
// @Description Update a user's role or name in a project
// @Tags Projects
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param project_id path int true "Project ID"
// @Param user_id path int true "User ID"
// @Param changeUser body dto.ProjectUserDTO true "Updated User Data (UserName, Role)"
// @Success 200 {object} map[string]interface{} "User updated successfully"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "User not found"
// @Failure 500 {object} map[string]string
// @Router /api/projects/{project_id}/change-user/{user_id} [put]
func (h *ProjectHandler) ChangeUser(c *gin.Context) {
	// Parse project ID from the path
	projectIDStr := c.Param("project_id")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil || projectID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	// Parse user ID from the path
	userIDStr := c.Param("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil || userID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Parse the user input DTO
	var userDTO dto.ProjectUserDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Extract claims for auditing
	claims, err := helpers.GetClaims(c)
	if err != nil {
		return
	}

	// Retrieve existing ProjectUser by ID
	projectUser, err := h.projectService.SearchProjectUsersById(userID, *claims.OrganizationId)
	if err != nil || projectUser == nil || projectUser.ProjectID != projectID {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found in the specified project"})
		return
	}

	// Update allowed fields
	if userDTO.UserName != "" {
		projectUser.UserName = userDTO.UserName
	}
	if userDTO.Role != "" {
		projectUser.Role = userDTO.Role
	}
	projectUser.UpdatedBy = &claims.UserID

	// Call the service to save updates
	if err := h.projectService.UpdateProjectUser(projectUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}
