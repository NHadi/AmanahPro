package handlers

import (
	"AmanahPro/services/spk-services/internal/application/services"
	"AmanahPro/services/spk-services/internal/domain/models"
	"AmanahPro/services/spk-services/internal/dto"
	"io"
	"net/http"
	"strconv"

	"github.com/NHadi/AmanahPro-common/helpers"
	"github.com/NHadi/AmanahPro-common/middleware"

	"github.com/gin-gonic/gin"
)

type SpkHandler struct {
	spkService *services.SpkService
}

// NewSpkHandler creates a new SpkHandler instance
func NewSpkHandler(spkService *services.SpkService) *SpkHandler {
	return &SpkHandler{
		spkService: spkService,
	}
}

// SPK CRUD

// FilterSpks
// @Summary Filter SPKs
// @Description Filter SPKs by organization ID, SPK ID, and project ID
// @Tags SPKs
// @Security BearerAuth
// @Param organization_id query int true "Organization ID"
// @Param spk_id query int false "SPK ID"
// @Param project_id query int false "Project ID"
// @Produce json
// @Success 200 {array} models.SPK
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string
// @Router /api/spk/filter [get]
func (h *SpkHandler) FilterSpks(c *gin.Context) {
	claims, err := helpers.GetClaims(c)
	if err != nil {
		return
	}

	organizationID := int(*claims.OrganizationId)

	// Parse optional query parameters
	spkIDStr := c.Query("spk_id")
	var spkID *int
	if spkIDStr != "" {
		id, err := strconv.Atoi(spkIDStr)
		if err != nil || id <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid SPK ID"})
			return
		}
		spkID = &id
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
	spks, err := h.spkService.Filter(organizationID, spkID, projectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, spks)
}

// CreateSpk
// @Summary Create SPK
// @Description Create a new SPK
// @Tags SPKs
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param spk body dto.SpkDTO true "SPK Data"
// @Success 201 {object} map[string]interface{} "Created SPK"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string
// @Router /api/spk [post]
func (h *SpkHandler) CreateSpk(c *gin.Context) {
	var spkDTO dto.SpkDTO
	if err := c.ShouldBindJSON(&spkDTO); err != nil {
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
	spk := models.SPK{
		ProjectId:      spkDTO.ProjectId,
		ProjectName:    spkDTO.ProjectName,
		Subject:        spkDTO.Subject,
		Date:           spkDTO.Date,
		OrganizationId: claims.OrganizationId,
		CreatedBy:      &claims.UserID,
		SphId:          &spkDTO.SphId,
	}

	// Call the service to create SPK
	if err := h.spkService.CreateSpk(&spk, int32(spkDTO.SphId), traceID.(string)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Map back to DTO for response
	spkDTO.SpkId = spk.SpkId

	c.JSON(http.StatusCreated, gin.H{
		"message": "SPK created successfully",
		"data":    spkDTO,
	})
}

// UpdateSpk
// @Summary Update SPK
// @Description Update an existing SPK
// @Tags SPKs
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param spk_id path int true "SPK ID"
// @Param spk body dto.SpkDTO true "SPK Data"
// @Success 200 {object} map[string]interface{} "Updated SPK"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "SPK Not Found"
// @Failure 500 {object} map[string]string
// @Router /api/spk/{spk_id} [put]
func (h *SpkHandler) UpdateSpk(c *gin.Context) {
	idStr := c.Param("spk_id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid SPK ID"})
		return
	}

	var spkDTO dto.SpkDTO
	if err := c.ShouldBindJSON(&spkDTO); err != nil {
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

	existingSpk, err := h.spkService.GetSpkByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "SPK not found"})
		return
	}

	updatedSpk := spkDTO.ToModelForUpdate(existingSpk, claims.UserID)
	updatedSpk.OrganizationId = claims.OrganizationId

	if err := h.spkService.UpdateSpk(updatedSpk, traceID.(string)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "SPK updated successfully",
		"data":    spkDTO,
	})
}

// DeleteSpk
// @Summary Delete SPK
// @Description Delete an SPK by ID
// @Tags SPKs
// @Security BearerAuth
// @Param spk_id path int true "SPK ID"
// @Produce json
// @Success 204 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "SPK Not Found"
// @Failure 500 {object} map[string]string
// @Router /api/spk/{spk_id} [delete]
func (h *SpkHandler) DeleteSpk(c *gin.Context) {
	idStr := c.Param("spk_id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid SPK ID"})
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

	if err := h.spkService.DeleteSpk(id, traceID.(string), claims.UserID); err != nil {
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
// @Param spk_id path int true "SPK ID"
// @Param section body dto.SpkSectionDTO true "SPK Section Data"
// @Success 201 {object} map[string]interface{} "Created SPK Section"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string
// @Router /api/spk/{spk_id}/sections [post]
func (h *SpkHandler) CreateSpkSection(c *gin.Context) {
	var sectionDTO dto.SpkSectionDTO
	if err := c.ShouldBindJSON(&sectionDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	spkIDStr := c.Param("spk_id")
	spkID, err := strconv.Atoi(spkIDStr)
	if err != nil || spkID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid SPK ID"})
		return
	}

	claims, err := helpers.GetClaims(c)
	if err != nil {
		return
	}

	section := sectionDTO.ToModel(claims.UserID)
	section.OrganizationId = claims.OrganizationId

	if err := h.spkService.CreateSpkSection(section, spkID); err != nil {
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
// @Param spk_id path int true "SPK ID"
// @Param section_id path int true "Section ID"
// @Param section body dto.SpkSectionDTO true "SPK Section Data"
// @Success 200 {object} map[string]interface{} "Updated SPK Section"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "SPK Section Not Found"
// @Failure 500 {object} map[string]string
// @Router /api/spk/{spk_id}/sections/{section_id} [put]
func (h *SpkHandler) UpdateSpkSection(c *gin.Context) {
	spkIDStr := c.Param("spk_id")
	spkID, err := strconv.Atoi(spkIDStr)
	if err != nil || spkID <= 0 {
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

	existingSection, err := h.spkService.GetSpkSectionByID(sectionID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "SPK section not found"})
		return
	}

	updatedSection := sectionDTO.ToModelForUpdate(existingSection, claims.UserID)
	updatedSection.OrganizationId = claims.OrganizationId

	if err := h.spkService.UpdateSpkSection(updatedSection); err != nil {
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
// @Param spk_id path int true "SPK ID"
// @Param section_id path int true "Section ID"
// @Produce json
// @Success 204 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "SPK Section Not Found"
// @Failure 500 {object} map[string]string
// @Router /api/spk/{spk_id}/sections/{section_id} [delete]
func (h *SpkHandler) DeleteSpkSection(c *gin.Context) {
	spkIDStr := c.Param("spk_id")
	spkID, err := strconv.Atoi(spkIDStr)
	if err != nil || spkID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid SPK ID"})
		return
	}

	sectionIDStr := c.Param("section_id")
	sectionID, err := strconv.Atoi(sectionIDStr)
	if err != nil || sectionID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Section ID"})
		return
	}

	if err := h.spkService.DeleteSpkSection(sectionID, spkID); err != nil {
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
// @Param spk_id path int true "SPK ID"
// @Param section_id path int true "Section ID"
// @Param detail body dto.SpkDetailDTO true "SPK Detail Data"
// @Success 201 {object} map[string]interface{} "Created SPK Detail"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string
// @Router /api/spk/{spk_id}/sections/{section_id}/details [post]
func (h *SpkHandler) CreateSpkDetail(c *gin.Context) {
	spkIDStr := c.Param("spk_id")
	spkID, err := strconv.Atoi(spkIDStr)
	if err != nil || spkID <= 0 {
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

	if err := h.spkService.CreateSpkDetail(detail, spkID, sectionID); err != nil {
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
// @Param spk_id path int true "SPK ID"
// @Param section_id path int true "Section ID"
// @Param detail_id path int true "Detail ID"
// @Param detail body dto.SpkDetailDTO true "SPK Detail Data"
// @Success 200 {object} map[string]interface{} "Updated SPK Detail"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "SPK Detail Not Found"
// @Failure 500 {object} map[string]string
// @Router /api/spk/{spk_id}/sections/{section_id}/details/{detail_id} [put]
func (h *SpkHandler) UpdateSpkDetail(c *gin.Context) {
	spkIDStr := c.Param("spk_id")
	spkID, err := strconv.Atoi(spkIDStr)
	if err != nil || spkID <= 0 {
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

	existingDetail, err := h.spkService.GetSpkDetailByID(detailID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "SPK detail not found"})
		return
	}

	updatedDetail := detailDTO.ToModelForUpdate(existingDetail, claims.UserID)
	updatedDetail.OrganizationId = claims.OrganizationId

	if err := h.spkService.UpdateSpkDetail(updatedDetail, spkID); err != nil {
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
// @Param spk_id path int true "SPK ID"
// @Param section_id path int true "Section ID"
// @Param detail_id path int true "Detail ID"
// @Produce json
// @Success 204 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "SPK Detail Not Found"
// @Failure 500 {object} map[string]string
// @Router /api/spk/{spk_id}/sections/{section_id}/details/{detail_id} [delete]
func (h *SpkHandler) DeleteSpkDetail(c *gin.Context) {
	spkIDStr := c.Param("spk_id")
	spkID, err := strconv.Atoi(spkIDStr)
	if err != nil || spkID <= 0 {
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

	if err := h.spkService.DeleteSpkDetail(detailID, spkID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// ImportSphFromExcel
// @Summary Import SPH data from an Excel file
// @Description Import SPH data along with metadata
// @Tags SPHs
// @Security BearerAuth
// @Accept multipart/form-data
// @Produce json
// @Param ProjectId formData string true "ProjectId"
// @Param Subject formData string true "Subject"
// @Param Date formData string true "Date (yyyy-MM-dd)"
// @Param Mandor formData string true "Mandor"
// @Param file formData file true "Excel File"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string
// @Router /api/spk/import [post]
func (h *SpkHandler) ImportFromExcel(c *gin.Context) {
	var metadata dto.SpkImportDTO

	// Bind form fields to metadata
	if err := c.ShouldBind(&metadata); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid metadata"})
		return
	}

	// Parse the file from the request
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File upload failed"})
		return
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
		return
	}

	claims, err := helpers.GetClaims(c)
	if err != nil {
		return
	}

	// Call the service to process the file and get the grand total
	spk, err := h.spkService.ImportSpkFromExcel(metadata, fileBytes, *claims.OrganizationId, claims.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the grand total as part of the response
	c.JSON(http.StatusOK, gin.H{
		"message":       "SPKP import successful",
		"totalJasa":     spk.TotalJasa,
		"totalMaterial": spk.TotalMaterial,
	})
}
