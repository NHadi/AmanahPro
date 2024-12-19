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

type ProjectFinancialHandler struct {
	service *services.ProjectFinancialService
}

// NewProjectFinancialHandler initializes ProjectFinancialHandler
func NewProjectFinancialHandler(service *services.ProjectFinancialService) *ProjectFinancialHandler {
	return &ProjectFinancialHandler{service: service}
}

// CreateProjectFinancial
// @Summary Create Financial Record
// @Description Add a new financial record for a project
// @Tags ProjectFinancial
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param financial body dto.ProjectFinancialDTO true "Financial Record Data"
// @Success 201 {object} map[string]interface{} "Created"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string
// @Router /api/project-financials [post]
func (h *ProjectFinancialHandler) CreateProjectFinancial(c *gin.Context) {
	var dto dto.ProjectFinancialDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	claims, _ := helpers.GetClaims(c)
	traceID, _ := c.Get(middleware.TraceIDHeader)

	model := dto.ToModel(claims.UserID, *claims.OrganizationId)

	if err := h.service.CreateProjectFinancial(model, traceID.(string)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Financial record created successfully"})
}

// UpdateProjectFinancial
// @Summary Update Financial Record
// @Description Update an existing financial record
// @Tags ProjectFinancial
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param financial_id path int true "Financial Record ID"
// @Param financial body dto.ProjectFinancialDTO true "Financial Record Data"
// @Success 200 {object} map[string]interface{} "Updated"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "Not Found"
// @Failure 500 {object} map[string]string
// @Router /api/project-financials/{financial_id} [put]
func (h *ProjectFinancialHandler) UpdateProjectFinancial(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("financial_id"))

	var dto dto.ProjectFinancialDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	claims, _ := helpers.GetClaims(c)
	traceID, _ := c.Get(middleware.TraceIDHeader)

	record, err := h.service.GetProjectFinancialByID(id, traceID.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	model := dto.ToModelForUpdate(record, claims.UserID)

	if err := h.service.UpdateProjectFinancial(model, traceID.(string)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Financial record updated successfully"})
}

// DeleteProjectFinancial
// @Summary Delete Financial Record
// @Description Remove a financial record by ID
// @Tags ProjectFinancial
// @Security BearerAuth
// @Param financial_id path int true "Financial Record ID"
// @Produce json
// @Success 204 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "Not Found"
// @Failure 500 {object} map[string]string
// @Router /api/project-financials/{financial_id} [delete]
func (h *ProjectFinancialHandler) DeleteProjectFinancial(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("financial_id"))
	claims, _ := helpers.GetClaims(c)
	traceID, _ := c.Get(middleware.TraceIDHeader)

	if err := h.service.DeleteProjectFinancial(id, traceID.(string), claims.UserID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetProjectFinancialByID
// @Summary Get Financial Record
// @Description Retrieve financial record by ID
// @Tags ProjectFinancial
// @Security BearerAuth
// @Param id path int true "Financial Record ID"
// @Produce json
// @Success 200 {object} models.ProjectFinancial
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "Not Found"
// @Router /api/project-financials/{id} [get]
func (h *ProjectFinancialHandler) GetProjectFinancialByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	traceID, _ := c.Get(middleware.TraceIDHeader)

	record, err := h.service.GetProjectFinancialByID(id, traceID.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, record)
}

// GetAllFinancialByProjectID
// @Summary Get All Financial Records
// @Description Retrieve all financial records for a given ProjectID
// @Tags ProjectFinancial
// @Security BearerAuth
// @Param project_id query int true "Project ID"
// @Produce json
// @Success 200 {array} models.ProjectFinancial
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string
// @Router /api/project-financials/project/{project_id} [get]
func (h *ProjectFinancialHandler) GetAllFinancialByProjectID(c *gin.Context) {
	projectID, _ := strconv.Atoi(c.Param("project_id"))
	traceID, _ := c.Get(middleware.TraceIDHeader)

	records, err := h.service.GetAllFinancialByProjectID(projectID, traceID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, records)
}

// GetProjectFinancialSummary
// @Summary Get Project Financial Summary
// @Description Get summarized financial data for all projects by Organization ID
// @Tags Projects
// @Security BearerAuth
// @Produce json
// @Success 200 {array} dto.ProjectFinancialSummaryDTO
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string
// @Router /api/project-financials/financial-summary [get]
func (h *ProjectFinancialHandler) GetProjectFinancialSummary(c *gin.Context) {
	// Extract claims for OrganizationID
	claims, err := helpers.GetClaims(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	organizationID := *claims.OrganizationId

	// Call the service to fetch the summary
	summary, err := h.service.GetProjectFinancialSummary(organizationID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, summary)
}
