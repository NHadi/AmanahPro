package handlers

import (
	"AmanahPro/services/breakdown-services/common/helpers"
	"AmanahPro/services/breakdown-services/internal/application/services"
	"AmanahPro/services/breakdown-services/internal/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MstBreakdownHandler struct {
	breakdownService *services.BreakdownService
}

// NewMstBreakdownHandler initializes a new handler
func NewMstBreakdownHandler(breakdownService *services.BreakdownService) *MstBreakdownHandler {
	return &MstBreakdownHandler{
		breakdownService: breakdownService,
	}
}

// MstBreakdownSection Handlers

// FilterMstBreakdownSections
// @Summary Filter Master Breakdown Sections
// @Description Retrieve all master breakdown sections for an organization
// @Tags MstBreakdownSections
// @Security BearerAuth
// @Produce json
// @Success 200 {array} models.MstBreakdownSection
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string
// @Router /api/mst-breakdown-sections [get]
func (h *MstBreakdownHandler) FilterMstBreakdownSections(c *gin.Context) {
	claims, err := helpers.GetClaims(c)
	if err != nil {
		return
	}

	sections, err := h.breakdownService.FilterMstBreakdownSections(claims.OrganizationId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, sections)
}

// CreateMstBreakdownSection
// @Summary Create Master Breakdown Section
// @Description Create a new master breakdown section
// @Tags MstBreakdownSections
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param section body dto.MstBreakdownSectionDTO true "Master Breakdown Section Data"
// @Success 201 {object} map[string]interface{} "Created Master Breakdown Section"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string
// @Router /api/mst-breakdown-sections [post]
func (h *MstBreakdownHandler) CreateMstBreakdownSection(c *gin.Context) {
	var sectionDTO dto.MstBreakdownSectionDTO
	if err := c.ShouldBindJSON(&sectionDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	claims, err := helpers.GetClaims(c)
	if err != nil {
		return
	}

	section := sectionDTO.ToModel(claims.UserID)
	section.OrganizationId = claims.OrganizationId

	if err := h.breakdownService.CreateMstBreakdownSection(section); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Master breakdown section created successfully",
		"data":    section,
	})
}

// UpdateMstBreakdownSection
// @Summary Update Master Breakdown Section
// @Description Update an existing master breakdown section
// @Tags MstBreakdownSections
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Master Breakdown Section ID"
// @Param section body dto.MstBreakdownSectionDTO true "Master Breakdown Section Data"
// @Success 200 {object} map[string]interface{} "Updated Master Breakdown Section"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "Not Found"
// @Failure 500 {object} map[string]string
// @Router /api/mst-breakdown-sections/{id} [put]
func (h *MstBreakdownHandler) UpdateMstBreakdownSection(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid section ID"})
		return
	}

	var sectionDTO dto.MstBreakdownSectionDTO
	if err := c.ShouldBindJSON(&sectionDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	claims, err := helpers.GetClaims(c)
	if err != nil {
		return
	}

	existingSection, err := h.breakdownService.GetMstBreakdownSectionyID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Breakdown section not found"})
		return
	}

	updatedSection := sectionDTO.ToModelForUpdate(existingSection, claims.UserID)
	updatedSection.OrganizationId = claims.OrganizationId

	if err := h.breakdownService.UpdateMstBreakdownSection(updatedSection); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Master breakdown section updated successfully",
		"data":    updatedSection,
	})
}

// DeleteMstBreakdownSection
// @Summary Delete Master Breakdown Section
// @Description Delete a master breakdown section by ID
// @Tags MstBreakdownSections
// @Security BearerAuth
// @Param id path int true "Master Breakdown Section ID"
// @Produce json
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "Not Found"
// @Failure 500 {object} map[string]string
// @Router /api/mst-breakdown-sections/{id} [delete]
func (h *MstBreakdownHandler) DeleteMstBreakdownSection(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid section ID"})
		return
	}

	if err := h.breakdownService.DeleteMstBreakdownSection(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// MstBreakdownItem Handlers

// CreateMstBreakdownItem
// @Summary Create Master Breakdown Item
// @Description Create a new master breakdown item
// @Tags MstBreakdownItems
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Master Breakdown Section ID"
// @Param item body dto.MstBreakdownItemDTO true "Master Breakdown Item Data"
// @Success 201 {object} map[string]interface{} "Created Master Breakdown Item"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string
// @Router /api/mst-breakdown-sections/{id}/items [post]
func (h *MstBreakdownHandler) CreateMstBreakdownItem(c *gin.Context) {
	var itemDTO dto.MstBreakdownItemDTO
	if err := c.ShouldBindJSON(&itemDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	sectionID, err := strconv.Atoi(c.Param("id"))
	if err != nil || sectionID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid section ID"})
		return
	}

	claims, err := helpers.GetClaims(c)
	if err != nil {
		return
	}

	item := itemDTO.ToModel(sectionID, claims.UserID)
	item.OrganizationId = claims.OrganizationId

	if err := h.breakdownService.CreateMstBreakdownItem(item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Master breakdown item created successfully",
		"data":    item,
	})
}

// UpdateMstBreakdownItem
// @Summary Update Master Breakdown Item
// @Description Update an existing master breakdown item
// @Tags MstBreakdownItems
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Master Breakdown Section ID"
// @Param item_id path int true "Master Breakdown Item ID"
// @Param item body dto.MstBreakdownItemDTO true "Master Breakdown Item Data"
// @Success 200 {object} map[string]interface{} "Updated Master Breakdown Item"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "Not Found"
// @Failure 500 {object} map[string]string
// @Router /api/mst-breakdown-sections/{id}/items/{item_id} [put]
func (h *MstBreakdownHandler) UpdateMstBreakdownItem(c *gin.Context) {
	itemID, err := strconv.Atoi(c.Param("item_id"))
	if err != nil || itemID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	var itemDTO dto.MstBreakdownItemDTO
	if err := c.ShouldBindJSON(&itemDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	claims, err := helpers.GetClaims(c)
	if err != nil {
		return
	}

	existingItem, err := h.breakdownService.GetMstBreakdownItemByID(itemID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Breakdown item not found"})
		return
	}

	updatedItem := itemDTO.ToModelForUpdate(existingItem, claims.UserID)

	if err := h.breakdownService.UpdateMstBreakdownItem(updatedItem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Master breakdown item updated successfully",
		"data":    updatedItem,
	})
}

// DeleteMstBreakdownItem
// @Summary Delete Master Breakdown Item
// @Description Delete a master breakdown item by ID
// @Tags MstBreakdownItems
// @Security BearerAuth
// @Param id path int true "Master Breakdown Section ID"
// @Param item_id path int true "Master Breakdown Item ID"
// @Produce json
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "Not Found"
// @Failure 500 {object} map[string]string
// @Router /api/mst-breakdown-sections/{id}/items/{item_id} [delete]
func (h *MstBreakdownHandler) DeleteMstBreakdownItem(c *gin.Context) {
	itemID, err := strconv.Atoi(c.Param("item_id"))
	if err != nil || itemID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	if err := h.breakdownService.DeleteMstBreakdownItem(itemID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
