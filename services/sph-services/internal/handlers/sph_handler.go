package handlers

import (
	"AmanahPro/services/sph-services/common/helpers"
	"AmanahPro/services/sph-services/internal/application/services"
	"AmanahPro/services/sph-services/internal/domain/models"
	"AmanahPro/services/sph-services/internal/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SphHandler struct {
	sphService *services.SphService
}

// NewSphHandler creates a new SphHandler instance
func NewSphHandler(sphService *services.SphService) *SphHandler {
	return &SphHandler{
		sphService: sphService,
	}
}

// SPH CRUD

// FilterSphs
// @Summary Filter SPHs
// @Description Filter SPHs by organization ID, SPH ID, and project ID
// @Tags SPHs
// @Security BearerAuth
// @Param organization_id query int true "Organization ID"
// @Param sph_id query int false "SPH ID"
// @Param project_id query int false "Project ID"
// @Produce json
// @Success 200 {array} models.Sph
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string
// @Router /api/sph/filter [get]
func (h *SphHandler) FilterSphs(c *gin.Context) {
	claims, err := helpers.GetClaims(c)
	if err != nil {
		// Error already handled in the helper
		return
	}

	organizationID := int(*claims.OrganizationId)

	// Parse optional query parameters
	sphIDStr := c.Query("sph_id")
	var sphID *int
	if sphIDStr != "" {
		id, err := strconv.Atoi(sphIDStr)
		if err != nil || id <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid SPH ID"})
			return
		}
		sphID = &id
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

	// Call the service to filter SPHs
	sphs, err := h.sphService.Filter(organizationID, sphID, projectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, sphs)
}

// CreateSph
// @Summary Create SPH
// @Description Create a new SPH
// @Tags SPHs
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param sph body dto.SphDTO true "SPH Data"
// @Success 201 {object} map[string]interface{} "Created SPH"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string
// @Router /api/sph [post]
func (h *SphHandler) CreateSph(c *gin.Context) {
	var sphDTO dto.SphDTO
	if err := c.ShouldBindJSON(&sphDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	claims, err := helpers.GetClaims(c)
	if err != nil {
		return
	}

	// Map DTO to Model
	sph := models.Sph{
		ProjectId:      sphDTO.ProjectId,
		ProjectName:    sphDTO.ProjectName,
		Subject:        sphDTO.Subject,
		Location:       sphDTO.Location,
		Date:           sphDTO.Date,
		OrganizationId: claims.OrganizationId,
		CreatedBy:      &claims.UserID,
		RecepientName:  sphDTO.RecepientName,
	}

	if err := h.sphService.CreateSph(&sph); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Map back to DTO for response
	sphDTO.SphId = sph.SphId

	c.JSON(http.StatusCreated, gin.H{
		"message": "SPH created successfully",
		"data":    sphDTO,
	})
}

// UpdateSph
// @Summary Update SPH
// @Description Update an existing SPH
// @Tags SPHs
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param sph_id path int true "SPH ID"
// @Param sph body dto.SphDTO true "SPH Data"
// @Success 200 {object} map[string]interface{} "Updated SPH"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "SPH Not Found"
// @Failure 500 {object} map[string]string
// @Router /api/sph/{sph_id} [put]
func (h *SphHandler) UpdateSph(c *gin.Context) {
	idStr := c.Param("sph_id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid SPH ID"})
		return
	}

	var sphDTO dto.SphDTO
	if err := c.ShouldBindJSON(&sphDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	claims, err := helpers.GetClaims(c)
	if err != nil {
		return
	}

	existingSph, err := h.sphService.GetSphByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "SPH not found"})
		return
	}

	updatedSph := sphDTO.ToModelForUpdate(existingSph, claims.UserID)
	updatedSph.OrganizationId = claims.OrganizationId

	if err := h.sphService.UpdateSph(updatedSph); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "SPH updated successfully",
		"data":    sphDTO,
	})
}

// DeleteSph
// @Summary Delete SPH
// @Description Delete an SPH by ID
// @Tags SPHs
// @Security BearerAuth
// @Param sph_id path int true "SPH ID"
// @Produce json
// @Success 204 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "SPH Not Found"
// @Failure 500 {object} map[string]string
// @Router /api/sph/{sph_id} [delete]
func (h *SphHandler) DeleteSph(c *gin.Context) {
	idStr := c.Param("sph_id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid SPH ID"})
		return
	}

	if err := h.sphService.DeleteSph(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// CreateSphSection
// @Summary Create SPH Section
// @Description Create a new SPH Section
// @Tags SPHSections
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param sph_id path int true "SPH ID"
// @Param section body dto.SphSectionDTO true "SPH Section Data"
// @Success 201 {object} map[string]interface{} "Created SPH Section"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string
// @Router /api/sph/{sph_id}/sections [post]
func (h *SphHandler) CreateSphSection(c *gin.Context) {
	var sectionDTO dto.SphSectionDTO
	if err := c.ShouldBindJSON(&sectionDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	sphIDStr := c.Param("sph_id")
	sphID, err := strconv.Atoi(sphIDStr)
	if err != nil || sphID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid SPH ID"})
		return
	}

	claims, err := helpers.GetClaims(c)
	if err != nil {
		return
	}

	section := sectionDTO.ToModel(sphID, claims.UserID)
	section.OrganizationId = claims.OrganizationId

	if err := h.sphService.CreateSphSection(section); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "SPH section created successfully",
		"data":    section,
	})
}

// UpdateSphSection
// @Summary Update SPH Section
// @Description Update an existing SPH Section
// @Tags SPHSections
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param sph_id path int true "SPH ID"
// @Param section_id path int true "Section ID"
// @Param section body dto.SphSectionDTO true "SPH Section Data"
// @Success 200 {object} map[string]interface{} "Updated SPH Section"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "SPH Section Not Found"
// @Failure 500 {object} map[string]string
// @Router /api/sph/{sph_id}/sections/{section_id} [put]
func (h *SphHandler) UpdateSphSection(c *gin.Context) {
	sphIDStr := c.Param("sph_id")
	sphID, err := strconv.Atoi(sphIDStr)
	if err != nil || sphID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid SPH ID"})
		return
	}

	sectionIDStr := c.Param("section_id")
	sectionID, err := strconv.Atoi(sectionIDStr)
	if err != nil || sectionID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Section ID"})
		return
	}

	var sectionDTO dto.SphSectionDTO
	if err := c.ShouldBindJSON(&sectionDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	claims, err := helpers.GetClaims(c)
	if err != nil {
		return
	}

	existingSection, err := h.sphService.GetSphSectionByID(sectionID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "SPH section not found"})
		return
	}

	updatedSection := sectionDTO.ToModelForUpdate(existingSection, claims.UserID)
	updatedSection.OrganizationId = claims.OrganizationId

	if err := h.sphService.UpdateSphSection(updatedSection); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "SPH section updated successfully",
		"data":    updatedSection,
	})
}

// DeleteSphSection
// @Summary Delete SPH Section
// @Description Delete an SPH Section by ID
// @Tags SPHSections
// @Security BearerAuth
// @Param sph_id path int true "SPH ID"
// @Param section_id path int true "Section ID"
// @Produce json
// @Success 204 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "SPH Section Not Found"
// @Failure 500 {object} map[string]string
// @Router /api/sph/{sph_id}/sections/{section_id} [delete]
func (h *SphHandler) DeleteSphSection(c *gin.Context) {
	sphIDStr := c.Param("sph_id")
	sphID, err := strconv.Atoi(sphIDStr)
	if err != nil || sphID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid SPH ID"})
		return
	}

	sectionIDStr := c.Param("section_id")
	sectionID, err := strconv.Atoi(sectionIDStr)
	if err != nil || sectionID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Section ID"})
		return
	}

	if err := h.sphService.DeleteSphSection(sectionID, sphID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// CreateSphDetail
// @Summary Create SPH Detail
// @Description Create a new SPH Detail
// @Tags SPHDetails
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param sph_id path int true "SPH ID"
// @Param section_id path int true "Section ID"
// @Param detail body dto.SphDetailDTO true "SPH Detail Data"
// @Success 201 {object} map[string]interface{} "Created SPH Detail"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string
// @Router /api/sph/{sph_id}/sections/{section_id}/details [post]
func (h *SphHandler) CreateSphDetail(c *gin.Context) {
	sphIDStr := c.Param("sph_id")
	sphID, err := strconv.Atoi(sphIDStr)
	if err != nil || sphID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid SPH ID"})
		return
	}

	sectionIDStr := c.Param("section_id")
	sectionID, err := strconv.Atoi(sectionIDStr)
	if err != nil || sectionID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Section ID"})
		return
	}

	var detailDTO dto.SphDetailDTO
	if err := c.ShouldBindJSON(&detailDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	claims, err := helpers.GetClaims(c)
	if err != nil {
		return
	}

	detail := detailDTO.ToModel(sectionID, claims.UserID)
	detail.OrganizationId = claims.OrganizationId

	if err := h.sphService.CreateSphDetail(detail, sphID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "SPH detail created successfully",
		"data":    detail,
	})
}

// UpdateSphDetail
// @Summary Update SPH Detail
// @Description Update an existing SPH Detail
// @Tags SPHDetails
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param sph_id path int true "SPH ID"
// @Param section_id path int true "Section ID"
// @Param detail_id path int true "Detail ID"
// @Param detail body dto.SphDetailDTO true "SPH Detail Data"
// @Success 200 {object} map[string]interface{} "Updated SPH Detail"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "SPH Detail Not Found"
// @Failure 500 {object} map[string]string
// @Router /api/sph/{sph_id}/sections/{section_id}/details/{detail_id} [put]
func (h *SphHandler) UpdateSphDetail(c *gin.Context) {
	sphIDStr := c.Param("sph_id")
	sphID, err := strconv.Atoi(sphIDStr)
	if err != nil || sphID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid SPH ID"})
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

	var detailDTO dto.SphDetailDTO
	if err := c.ShouldBindJSON(&detailDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	claims, err := helpers.GetClaims(c)
	if err != nil {
		return
	}

	existingDetail, err := h.sphService.GetSphDetailByID(detailID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "SPH detail not found"})
		return
	}

	updatedDetail := detailDTO.ToModelForUpdate(existingDetail, claims.UserID)
	updatedDetail.OrganizationId = claims.OrganizationId

	if err := h.sphService.UpdateSphDetail(updatedDetail, sphID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "SPH detail updated successfully",
		"data":    updatedDetail,
	})
}

// DeleteSphDetail
// @Summary Delete SPH Detail
// @Description Delete an SPH Detail by ID
// @Tags SPHDetails
// @Security BearerAuth
// @Param sph_id path int true "SPH ID"
// @Param section_id path int true "Section ID"
// @Param detail_id path int true "Detail ID"
// @Produce json
// @Success 204 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "SPH Detail Not Found"
// @Failure 500 {object} map[string]string
// @Router /api/sph/{sph_id}/sections/{section_id}/details/{detail_id} [delete]
func (h *SphHandler) DeleteSphDetail(c *gin.Context) {
	sphIDStr := c.Param("sph_id")
	sphID, err := strconv.Atoi(sphIDStr)
	if err != nil || sphID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid SPH ID"})
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

	if err := h.sphService.DeleteSphDetail(detailID, sphID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
