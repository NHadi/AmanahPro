package handlers

import (
	"AmanahPro/services/ba-services/internal/application/services"
	"AmanahPro/services/ba-services/internal/domain/models"
	"AmanahPro/services/ba-services/internal/dto"
	"net/http"
	"strconv"

	"github.com/NHadi/AmanahPro-common/helpers"

	"github.com/gin-gonic/gin"
)

type SpkHandler struct {
	baService *services.SpkService
}

// NewSpkHandler creates a new SpkHandler instance
func NewSpkHandler(baService *services.SpkService) *SpkHandler {
	return &SpkHandler{
		baService: baService,
	}
}

// SPK CRUD

// FilterSpks
// @Summary Filter SPKs
// @Description Filter SPKs by organization ID, SPK ID, and project ID
// @Tags SPKs
// @Security BearerAuth
// @Param organization_id query int true "Organization ID"
// @Param ba_id query int false "SPK ID"
// @Param project_id query int false "Project ID"
// @Produce json
// @Success 200 {array} models.SPK
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string
// @Router /api/ba/filter [get]
func (h *SpkHandler) FilterSpks(c *gin.Context) {
	claims, err := helpers.GetClaims(c)
	if err != nil {
		return
	}

	organizationID := int(*claims.OrganizationId)

	// Parse optional query parameters
	baIDStr := c.Query("ba_id")
	var baID *int
	if baIDStr != "" {
		id, err := strconv.Atoi(baIDStr)
		if err != nil || id <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid SPK ID"})
			return
		}
		baID = &id
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

	// Call the service to filter SPKs
	bas, err := h.baService.Filter(organizationID, baID, projectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, bas)
}

// CreateSpk
// @Summary Create SPK
// @Description Create a new SPK
// @Tags SPKs
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param ba body dto.SpkDTO true "SPK Data"
// @Success 201 {object} map[string]interface{} "Created SPK"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string
// @Router /api/ba [post]
func (h *SpkHandler) CreateSpk(c *gin.Context) {
	var baDTO dto.SpkDTO
	if err := c.ShouldBindJSON(&baDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	claims, err := helpers.GetClaims(c)
	if err != nil {
		return
	}

	// Map DTO to Model
	ba := models.SPK{
		ProjectId:      baDTO.ProjectId,
		ProjectName:    baDTO.ProjectName,
		Subject:        baDTO.Subject,
		Date:           baDTO.Date,
		OrganizationId: claims.OrganizationId,
		CreatedBy:      &claims.UserID,
		SphId:          baDTO.SphId,
	}

	// Call the service to create SPK
	if err := h.baService.CreateSpk(&ba, int32(baDTO.SphId)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Map back to DTO for response
	baDTO.SpkId = ba.SpkId

	c.JSON(http.StatusCreated, gin.H{
		"message": "SPK created successfully",
		"data":    baDTO,
	})
}

// UpdateSpk
// @Summary Update SPK
// @Description Update an existing SPK
// @Tags SPKs
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param ba_id path int true "SPK ID"
// @Param ba body dto.SpkDTO true "SPK Data"
// @Success 200 {object} map[string]interface{} "Updated SPK"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "SPK Not Found"
// @Failure 500 {object} map[string]string
// @Router /api/ba/{ba_id} [put]
func (h *SpkHandler) UpdateSpk(c *gin.Context) {
	idStr := c.Param("ba_id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid SPK ID"})
		return
	}

	var baDTO dto.SpkDTO
	if err := c.ShouldBindJSON(&baDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	claims, err := helpers.GetClaims(c)
	if err != nil {
		return
	}

	existingSpk, err := h.baService.GetSpkByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "SPK not found"})
		return
	}

	updatedSpk := baDTO.ToModelForUpdate(existingSpk, claims.UserID)
	updatedSpk.OrganizationId = claims.OrganizationId

	if err := h.baService.UpdateSpk(updatedSpk); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "SPK updated successfully",
		"data":    baDTO,
	})
}

// DeleteSpk
// @Summary Delete SPK
// @Description Delete an SPK by ID
// @Tags SPKs
// @Security BearerAuth
// @Param ba_id path int true "SPK ID"
// @Produce json
// @Success 204 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "SPK Not Found"
// @Failure 500 {object} map[string]string
// @Router /api/ba/{ba_id} [delete]
func (h *SpkHandler) DeleteSpk(c *gin.Context) {
	idStr := c.Param("ba_id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid SPK ID"})
		return
	}

	if err := h.baService.DeleteSpk(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// CreateSpkSection
// @Summary Create SPK Section
// @Description Create a new SPK Section
// @Tags SPKSections
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param ba_id path int true "SPK ID"
// @Param section body dto.SpkSectionDTO true "SPK Section Data"
// @Success 201 {object} map[string]interface{} "Created SPK Section"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string
// @Router /api/ba/{ba_id}/sections [post]
func (h *SpkHandler) CreateSpkSection(c *gin.Context) {
	var sectionDTO dto.SpkSectionDTO
	if err := c.ShouldBindJSON(&sectionDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	baIDStr := c.Param("ba_id")
	baID, err := strconv.Atoi(baIDStr)
	if err != nil || baID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid SPK ID"})
		return
	}

	claims, err := helpers.GetClaims(c)
	if err != nil {
		return
	}

	section := sectionDTO.ToModel(claims.UserID)
	section.OrganizationId = claims.OrganizationId

	if err := h.baService.CreateSpkSection(section, baID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "SPK section created successfully",
		"data":    section,
	})
}

// UpdateSpkSection
// @Summary Update SPK Section
// @Description Update an existing SPK Section
// @Tags SPKSections
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param ba_id path int true "SPK ID"
// @Param section_id path int true "Section ID"
// @Param section body dto.SpkSectionDTO true "SPK Section Data"
// @Success 200 {object} map[string]interface{} "Updated SPK Section"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "SPK Section Not Found"
// @Failure 500 {object} map[string]string
// @Router /api/ba/{ba_id}/sections/{section_id} [put]
func (h *SpkHandler) UpdateSpkSection(c *gin.Context) {
	baIDStr := c.Param("ba_id")
	baID, err := strconv.Atoi(baIDStr)
	if err != nil || baID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid SPK ID"})
		return
	}

	sectionIDStr := c.Param("section_id")
	sectionID, err := strconv.Atoi(sectionIDStr)
	if err != nil || sectionID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Section ID"})
		return
	}

	var sectionDTO dto.SpkSectionDTO
	if err := c.ShouldBindJSON(&sectionDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	claims, err := helpers.GetClaims(c)
	if err != nil {
		return
	}

	existingSection, err := h.baService.GetSpkSectionByID(sectionID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "SPK section not found"})
		return
	}

	updatedSection := sectionDTO.ToModelForUpdate(existingSection, claims.UserID)
	updatedSection.OrganizationId = claims.OrganizationId

	if err := h.baService.UpdateSpkSection(updatedSection); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "SPK section updated successfully",
		"data":    updatedSection,
	})
}

// DeleteSpkSection
// @Summary Delete SPK Section
// @Description Delete an SPK Section by ID
// @Tags SPKSections
// @Security BearerAuth
// @Param ba_id path int true "SPK ID"
// @Param section_id path int true "Section ID"
// @Produce json
// @Success 204 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "SPK Section Not Found"
// @Failure 500 {object} map[string]string
// @Router /api/ba/{ba_id}/sections/{section_id} [delete]
func (h *SpkHandler) DeleteSpkSection(c *gin.Context) {
	baIDStr := c.Param("ba_id")
	baID, err := strconv.Atoi(baIDStr)
	if err != nil || baID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid SPK ID"})
		return
	}

	sectionIDStr := c.Param("section_id")
	sectionID, err := strconv.Atoi(sectionIDStr)
	if err != nil || sectionID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Section ID"})
		return
	}

	if err := h.baService.DeleteSpkSection(sectionID, baID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// CreateSpkDetail
// @Summary Create SPK Detail
// @Description Create a new SPK Detail
// @Tags SPKDetails
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param ba_id path int true "SPK ID"
// @Param section_id path int true "Section ID"
// @Param detail body dto.SpkDetailDTO true "SPK Detail Data"
// @Success 201 {object} map[string]interface{} "Created SPK Detail"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string
// @Router /api/ba/{ba_id}/sections/{section_id}/details [post]
func (h *SpkHandler) CreateSpkDetail(c *gin.Context) {
	baIDStr := c.Param("ba_id")
	baID, err := strconv.Atoi(baIDStr)
	if err != nil || baID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid SPK ID"})
		return
	}

	sectionIDStr := c.Param("section_id")
	sectionID, err := strconv.Atoi(sectionIDStr)
	if err != nil || sectionID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Section ID"})
		return
	}

	var detailDTO dto.SpkDetailDTO
	if err := c.ShouldBindJSON(&detailDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	claims, err := helpers.GetClaims(c)
	if err != nil {
		return
	}

	detail := detailDTO.ToModel(claims.UserID)
	detail.OrganizationId = claims.OrganizationId

	if err := h.baService.CreateSpkDetail(detail, baID, sectionID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "SPK detail created successfully",
		"data":    detail,
	})
}

// UpdateSpkDetail
// @Summary Update SPK Detail
// @Description Update an existing SPK Detail
// @Tags SPKDetails
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param ba_id path int true "SPK ID"
// @Param section_id path int true "Section ID"
// @Param detail_id path int true "Detail ID"
// @Param detail body dto.SpkDetailDTO true "SPK Detail Data"
// @Success 200 {object} map[string]interface{} "Updated SPK Detail"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "SPK Detail Not Found"
// @Failure 500 {object} map[string]string
// @Router /api/ba/{ba_id}/sections/{section_id}/details/{detail_id} [put]
func (h *SpkHandler) UpdateSpkDetail(c *gin.Context) {
	baIDStr := c.Param("ba_id")
	baID, err := strconv.Atoi(baIDStr)
	if err != nil || baID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid SPK ID"})
		return
	}

	sectionIDStr := c.Param("section_id")
	sectionID, err := strconv.Atoi(sectionIDStr)
	if err != nil || sectionID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Section ID"})
		return
	}

	detailIDStr := c.Param("detail_id")
	detailID, err := strconv.Atoi(detailIDStr)
	if err != nil || detailID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Detail ID"})
		return
	}

	var detailDTO dto.SpkDetailDTO
	if err := c.ShouldBindJSON(&detailDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	claims, err := helpers.GetClaims(c)
	if err != nil {
		return
	}

	existingDetail, err := h.baService.GetSpkDetailByID(detailID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "SPK detail not found"})
		return
	}

	updatedDetail := detailDTO.ToModelForUpdate(existingDetail, claims.UserID)
	updatedDetail.OrganizationId = claims.OrganizationId

	if err := h.baService.UpdateSpkDetail(updatedDetail, baID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "SPK detail updated successfully",
		"data":    updatedDetail,
	})
}

// DeleteSpkDetail
// @Summary Delete SPK Detail
// @Description Delete an SPK Detail by ID
// @Tags SPKDetails
// @Security BearerAuth
// @Param ba_id path int true "SPK ID"
// @Param section_id path int true "Section ID"
// @Param detail_id path int true "Detail ID"
// @Produce json
// @Success 204 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "SPK Detail Not Found"
// @Failure 500 {object} map[string]string
// @Router /api/ba/{ba_id}/sections/{section_id}/details/{detail_id} [delete]
func (h *SpkHandler) DeleteSpkDetail(c *gin.Context) {
	baIDStr := c.Param("ba_id")
	baID, err := strconv.Atoi(baIDStr)
	if err != nil || baID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid SPK ID"})
		return
	}

	sectionIDStr := c.Param("section_id")
	sectionID, err := strconv.Atoi(sectionIDStr)
	if err != nil || sectionID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Section ID"})
		return
	}

	detailIDStr := c.Param("detail_id")
	detailID, err := strconv.Atoi(detailIDStr)
	if err != nil || detailID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Detail ID"})
		return
	}

	if err := h.baService.DeleteSpkDetail(detailID, baID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
