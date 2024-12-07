package handlers

import (
	"AmanahPro/services/breakdown-services/common/helpers"
	"AmanahPro/services/breakdown-services/internal/application/services"
	"AmanahPro/services/breakdown-services/internal/domain/models"
	"AmanahPro/services/breakdown-services/internal/dto"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BreakdownHandler struct {
	breakdownService *services.BreakdownService
}

// NewBreakdownHandler creates a new BreakdownHandler instance
func NewBreakdownHandler(breakdownService *services.BreakdownService) *BreakdownHandler {
	return &BreakdownHandler{
		breakdownService: breakdownService,
	}
}

// Breakdown CRUD

// FilterBreakdowns
// @Summary Filter Breakdowns
// @Description Filter breakdowns by organization ID, breakdown ID, and project ID
// @Tags Breakdowns
// @Security BearerAuth
// @Param organization_id query int true "Organization ID"
// @Param breakdown_id query int false "Breakdown ID"
// @Param project_id query int false "Project ID"
// @Produce json
// @Success 200 {array} models.Breakdown
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string
// @Router /api/breakdowns/filter [get]
func (h *BreakdownHandler) FilterBreakdowns(c *gin.Context) {
	claims, err := helpers.GetClaims(c)
	if err != nil {
		// Error already handled in the helper
		return
	}

	organizationID := int(*claims.OrganizationId)

	// Parse optional query parameters
	breakdownIDStr := c.Query("breakdown_id")
	var breakdownID *int
	if breakdownIDStr != "" {
		id, err := strconv.Atoi(breakdownIDStr)
		if err != nil || id <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid breakdown ID"})
			return
		}
		breakdownID = &id
	}

	projectIDStr := c.Query("project_id")
	var projectID *int
	if projectIDStr != "" {
		id, err := strconv.Atoi(projectIDStr)
		if err != nil || id <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
			return
		}
		projectID = &id
	}

	// Call the service to filter breakdowns
	breakdowns, err := h.breakdownService.FilterBreakdowns(organizationID, breakdownID, projectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, breakdowns)
}

// CreateBreakdown
// @Summary Create Breakdown
// @Description Create a new Breakdown
// @Tags Breakdowns
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param breakdown body dto.BreakdownDTO true "Breakdown Data"
// @Success 201 {object} map[string]interface{} "Created Breakdown"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string
// @Router /api/breakdowns [post]
func (h *BreakdownHandler) CreateBreakdown(c *gin.Context) {
	var breakdownDTO dto.BreakdownDTO
	if err := c.ShouldBindJSON(&breakdownDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	claims, err := helpers.GetClaims(c)
	if err != nil {
		return
	}

	// Map DTO to Model
	breakdown := models.Breakdown{
		ProjectId:      breakdownDTO.ProjectId,
		ProjectName:    breakdownDTO.ProjectName,
		Subject:        breakdownDTO.Subject,
		Location:       breakdownDTO.Location,
		Date:           breakdownDTO.Date,
		OrganizationId: claims.OrganizationId,
		CreatedBy:      &claims.UserID,
	}

	if err := h.breakdownService.CreateBreakdown(&breakdown); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Map back to DTO for response
	breakdownDTO.BreakdownId = breakdown.BreakdownId

	c.JSON(http.StatusCreated, gin.H{
		"message": "Breakdown created successfully",
		"data":    breakdownDTO,
	})
}

// UpdateBreakdown
// @Summary Update Breakdown
// @Description Update an existing Breakdown
// @Tags Breakdowns
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param breakdown_id path int true "Breakdown ID"
// @Param breakdown body dto.BreakdownDTO true "Breakdown Data"
// @Success 200 {object} map[string]interface{} "Updated Breakdown"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "Breakdown Not Found"
// @Failure 500 {object} map[string]string
// @Router /api/breakdowns/{breakdown_id} [put]
func (h *BreakdownHandler) UpdateBreakdown(c *gin.Context) {
	idStr := c.Param("breakdown_id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid breakdown ID"})
		return
	}

	var breakdownDTO dto.BreakdownDTO
	if err := c.ShouldBindJSON(&breakdownDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	claims, err := helpers.GetClaims(c)
	if err != nil {
		return
	}

	// Fetch existing breakdown from the database
	existingBreakdown, err := h.breakdownService.GetBreakdownByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Breakdown not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch breakdown"})
		}
		return
	}

	// Map incoming data to existing breakdown (only overwrite non-zero fields)
	breakdown := models.Breakdown{
		BreakdownId:    id,
		ProjectId:      firstNonZeroInt(breakdownDTO.ProjectId, existingBreakdown.ProjectId),
		ProjectName:    firstNonEmptyString(breakdownDTO.ProjectName, existingBreakdown.ProjectName),
		Subject:        firstNonEmptyString(breakdownDTO.Subject, existingBreakdown.Subject),
		Location:       firstNonEmptyStringPointer(breakdownDTO.Location, existingBreakdown.Location),
		Date:           firstNonZeroCustomDate(breakdownDTO.Date, existingBreakdown.Date),
		OrganizationId: claims.OrganizationId,
		UpdatedBy:      &claims.UserID,
	}

	if err := h.breakdownService.UpdateBreakdown(&breakdown); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Map back to DTO for response
	breakdownDTO.BreakdownId = breakdown.BreakdownId

	c.JSON(http.StatusOK, gin.H{
		"message": "Breakdown updated successfully",
		"data":    breakdownDTO,
	})
}

// DeleteBreakdown
// @Summary Delete Breakdown
// @Description Delete a Breakdown by ID
// @Tags Breakdowns
// @Security BearerAuth
// @Param breakdown_id path int true "Breakdown ID"
// @Produce json
// @Success 204 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "Breakdown Not Found"
// @Failure 500 {object} map[string]string
// @Router /api/breakdowns/{breakdown_id} [delete]
func (h *BreakdownHandler) DeleteBreakdown(c *gin.Context) {
	idStr := c.Param("breakdown_id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid breakdown ID"})
		return
	}

	if err := h.breakdownService.DeleteBreakdown(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// BreakdownSection CRUD

// CreateBreakdownSection
// @Summary Create Breakdown Section
// @Description Create a new Breakdown Section
// @Tags BreakdownSections
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param section body dto.BreakdownSectionDTO true "Breakdown Section Data"
// @Success 201 {object} map[string]interface{} "Created Breakdown Section"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string
// @Router /api/breakdowns/{breakdown_id}/sections [post]
func (h *BreakdownHandler) CreateBreakdownSection(c *gin.Context) {
	// Parse the request body into the DTO
	var sectionDTO dto.BreakdownSectionDTO
	if err := c.ShouldBindJSON(&sectionDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Parse and validate the breakdown ID from the URL
	breakdownIDStr := c.Param("breakdown_id")
	breakdownID, err := strconv.Atoi(breakdownIDStr)
	if err != nil || breakdownID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid breakdown ID"})
		return
	}

	// Extract claims for the user context
	claims, err := helpers.GetClaims(c)
	if err != nil {
		// Error already handled in the helper
		return
	}

	// Map the DTO to the model
	section := sectionDTO.ToModel(breakdownID, claims.UserID)

	organizationID := int(*claims.OrganizationId)
	section.OrganizationId = &organizationID

	// Pass the model to the service layer for creation
	if err := h.breakdownService.CreateBreakdownSection(section); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a successful response with the created section data
	c.JSON(http.StatusCreated, gin.H{
		"message": "Breakdown section created successfully",
		"data":    section,
	})
}

// UpdateBreakdownSection
// @Summary Update Breakdown Section
// @Description Update an existing Breakdown Section
// @Tags BreakdownSections
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Breakdown Section ID"
// @Param section body dto.BreakdownSectionDTO true "Breakdown Section Data"
// @Success 200 {object} map[string]interface{} "Updated Breakdown Section"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "Breakdown Section Not Found"
// @Failure 500 {object} map[string]string
// @Router /api/breakdowns/{breakdown_id}/sections/{id} [put]
func (h *BreakdownHandler) UpdateBreakdownSection(c *gin.Context) {
	// Parse the section ID from the URL
	idStr := c.Param("section_id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid section ID"})
		return
	}

	// Parse the breakdown ID from the URL
	breakdownIDStr := c.Param("breakdown_id")
	breakdownID, err := strconv.Atoi(breakdownIDStr)
	if err != nil || breakdownID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid breakdown ID"})
		return
	}

	// Parse the request body into the DTO
	var sectionDTO dto.BreakdownSectionDTO
	if err := c.ShouldBindJSON(&sectionDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Retrieve the user claims
	claims, err := helpers.GetClaims(c)
	if err != nil {
		// Error already handled in the helper
		return
	}

	// Retrieve the existing section from the service
	existingSection, err := h.breakdownService.GetBreakdownSectionByID(id, breakdownID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Breakdown section not found"})
		return
	}

	// Map the DTO to the model
	updatedSection := sectionDTO.ToModelForUpdate(existingSection, claims.UserID)
	organizationID := int(*claims.OrganizationId)
	updatedSection.OrganizationId = &organizationID

	// Call the service to update the breakdown section
	if err := h.breakdownService.UpdateBreakdownSection(updatedSection); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success response with the updated section data
	c.JSON(http.StatusOK, gin.H{
		"message": "Breakdown section updated successfully",
		"data":    updatedSection,
	})
}

// DeleteBreakdownSection
// @Summary Delete Breakdown Section
// @Description Delete a Breakdown Section by ID
// @Tags BreakdownSections
// @Security BearerAuth
// @Param id path int true "Breakdown Section ID"
// @Param breakdown_id path int true "Breakdown ID"
// @Produce json
// @Success 204 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "Breakdown Section Not Found"
// @Failure 500 {object} map[string]string
// @Router /api/breakdowns/{breakdown_id}/sections/{id} [delete]
func (h *BreakdownHandler) DeleteBreakdownSection(c *gin.Context) {
	idStr := c.Param("section_id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid section ID"})
		return
	}

	breakdownIDStr := c.Param("breakdown_id")
	breakdownID, err := strconv.Atoi(breakdownIDStr)
	if err != nil || breakdownID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid breakdown ID"})
		return
	}

	if err := h.breakdownService.DeleteBreakdownSection(id, breakdownID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// BreakdownItem CRUD

// CreateBreakdownItem
// @Summary Create Breakdown Item
// @Description Create a new Breakdown Item
// @Tags BreakdownItems
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param section_id path int true "Breakdown Section ID"
// @Param item body dto.BreakdownItemDTO true "Breakdown Item Data"
// @Success 201 {object} map[string]interface{} "Created Breakdown Item"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string
// @Router /api/breakdowns/{breakdown_id}/sections/{section_id}/items [post]
func (h *BreakdownHandler) CreateBreakdownItem(c *gin.Context) {
	var item dto.BreakdownItemDTO
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	sectionIDStr := c.Param("section_id")
	sectionID, err := strconv.Atoi(sectionIDStr)
	if err != nil || sectionID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid section ID"})
		return
	}

	breakdownIDStr := c.Param("breakdown_id")
	breakdownID, err := strconv.Atoi(breakdownIDStr)
	if err != nil || breakdownID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid breakdown ID"})
		return
	}

	claims, err := helpers.GetClaims(c)
	if err != nil {
		// Error already handled inside the helper
		return
	}

	// Map the DTO to the model
	data := item.ToModel(sectionID, claims.UserID)
	organizationID := int(*claims.OrganizationId)
	data.OrganizationId = &organizationID

	if err := h.breakdownService.CreateBreakdownItem(data, breakdownID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Breakdown item created successfully",
		"data":    item,
	})
}

// UpdateBreakdownItem
// @Summary Update Breakdown Item
// @Description Update an existing Breakdown Item
// @Tags BreakdownItems
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Breakdown Item ID"
// @Param item body dto.BreakdownItemDTO true "Breakdown Item Data"
// @Success 200 {object} map[string]interface{} "Updated Breakdown Item"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "Breakdown Item Not Found"
// @Failure 500 {object} map[string]string
// @Router /api/breakdowns/{breakdown_id}/sections/{section_id}/items/{id} [put]
func (h *BreakdownHandler) UpdateBreakdownItem(c *gin.Context) {
	idStr := c.Param("item_id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	var item dto.BreakdownItemDTO
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	sectionIDStr := c.Param("section_id")
	sectionID, err := strconv.Atoi(sectionIDStr)
	if err != nil || sectionID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid section ID"})
		return
	}

	breakdownIDStr := c.Param("breakdown_id")
	breakdownID, err := strconv.Atoi(breakdownIDStr)
	if err != nil || breakdownID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid breakdown ID"})
		return
	}

	claims, err := helpers.GetClaims(c)
	if err != nil {
		// Error already handled inside the helper
		return
	}

	// Retrieve the existing section from the service
	existingItem, err := h.breakdownService.GetBreakdownItemByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Breakdown section not found"})
		return
	}

	// Map the DTO to the model
	updatedItem := item.ToModelForUpdate(existingItem, claims.UserID)

	if err := h.breakdownService.UpdateBreakdownItem(updatedItem, breakdownID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Breakdown item updated successfully",
		"data":    item,
	})
}

// DeleteBreakdownItem
// @Summary Delete Breakdown Item
// @Description Delete a Breakdown Item by ID
// @Tags BreakdownItems
// @Security BearerAuth
// @Param id path int true "Breakdown Item ID"
// @Produce json
// @Success 204 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "Breakdown Item Not Found"
// @Failure 500 {object} map[string]string
// @Router /api/breakdowns/{breakdown_id}/sections/{section_id}/items/{id} [delete]
func (h *BreakdownHandler) DeleteBreakdownItem(c *gin.Context) {
	idStr := c.Param("item_id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	sectionIDStr := c.Param("section_id")
	sectionID, err := strconv.Atoi(sectionIDStr)
	if err != nil || sectionID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid section ID"})
		return
	}

	breakdownIDStr := c.Param("breakdown_id")
	breakdownID, err := strconv.Atoi(breakdownIDStr)
	if err != nil || breakdownID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid breakdown ID"})
		return
	}

	if err := h.breakdownService.DeleteBreakdownItem(id, breakdownID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
