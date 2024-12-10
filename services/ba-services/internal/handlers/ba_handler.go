package handlers

import (
	"AmanahPro/services/ba-services/internal/application/services"
	"AmanahPro/services/ba-services/internal/dto"
	"net/http"
	"strconv"

	"github.com/NHadi/AmanahPro-common/helpers"
	"github.com/gin-gonic/gin"
)

type BAHandler struct {
	baService *services.BAService
}

// NewBAHandler creates a new BAHandler instance
func NewBAHandler(baService *services.BAService) *BAHandler {
	return &BAHandler{
		baService: baService,
	}
}

// FilterBAs
// @Summary Filter BAs
// @Description Filter BAs by organization ID, BA ID, and project ID
// @Tags BAs
// @Security BearerAuth
// @Param organization_id query int true "Organization ID"
// @Param ba_id query int false "BA ID"
// @Param project_id query int false "Project ID"
// @Produce json
// @Success 200 {array} models.BA
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string
// @Router /api/ba/filter [get]
func (h *BAHandler) FilterBAs(c *gin.Context) {
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
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid BA ID"})
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

	// Call the service to filter BAs
	bas, err := h.baService.Filter(organizationID, baID, projectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, bas)
}

// CreateBA
// @Summary Create BA
// @Description Create a new BA
// @Tags BAs
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param ba body dto.BADTO true "BA Data"
// @Success 201 {object} map[string]interface{} "Created BA"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string
// @Router /api/ba [post]
func (h *BAHandler) CreateBA(c *gin.Context) {
	var baDTO dto.BADTO
	if err := c.ShouldBindJSON(&baDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	claims, err := helpers.GetClaims(c)
	if err != nil {
		return
	}

	// Map DTO to Model
	ba := baDTO.ToModel(claims.UserID)
	ba.OrganizationId = claims.OrganizationId
	// Call the service to create BA
	if err := h.baService.CreateBA(ba, int32(baDTO.SphId)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "BA created successfully",
		"data":    baDTO,
	})
}

// UpdateBA
// @Summary Update BA
// @Description Update an existing BA
// @Tags BAs
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param ba_id path int true "BA ID"
// @Param ba body dto.BADTO true "BA Data"
// @Success 200 {object} map[string]interface{} "Updated BA"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "BA Not Found"
// @Failure 500 {object} map[string]string
// @Router /api/ba/{ba_id} [put]
func (h *BAHandler) UpdateBA(c *gin.Context) {
	baIDStr := c.Param("ba_id")
	baID, err := strconv.Atoi(baIDStr)
	if err != nil || baID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid BA ID"})
		return
	}

	var baDTO dto.BADTO
	if err := c.ShouldBindJSON(&baDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	claims, err := helpers.GetClaims(c)
	if err != nil {
		return
	}

	existingBA, err := h.baService.GetBAByID(baID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "BA not found"})
		return
	}

	updatedBA := baDTO.ToModelForUpdate(existingBA, claims.UserID)
	if err := h.baService.UpdateBA(updatedBA); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "BA updated successfully",
		"data":    baDTO,
	})
}

// DeleteBA
// @Summary Delete BA
// @Description Delete an existing BA
// @Tags BAs
// @Security BearerAuth
// @Param ba_id path int true "BA ID"
// @Produce json
// @Success 204 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "BA Not Found"
// @Failure 500 {object} map[string]string
// @Router /api/ba/{ba_id} [delete]
func (h *BAHandler) DeleteBA(c *gin.Context) {
	baIDStr := c.Param("ba_id")
	baID, err := strconv.Atoi(baIDStr)
	if err != nil || baID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid BA ID"})
		return
	}

	claims, err := helpers.GetClaims(c)
	if err != nil {
		return
	}

	if err := h.baService.DeleteBA(baID, claims.UserID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// CreateBASection
// @Summary Create BA Section
// @Description Create a new BA Section
// @Tags BASections
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param ba_id path int true "BA ID"
// @Param section body dto.BASectionDTO true "BA Section Data"
// @Success 201 {object} map[string]interface{} "Created BA Section"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string
// @Router /api/ba/{ba_id}/sections [post]
func (h *BAHandler) CreateBASection(c *gin.Context) {
	baIDStr := c.Param("ba_id")
	baID, err := strconv.Atoi(baIDStr)
	if err != nil || baID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid BA ID"})
		return
	}

	var sectionDTO dto.BASectionDTO
	if err := c.ShouldBindJSON(&sectionDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	claims, err := helpers.GetClaims(c)
	if err != nil {
		return
	}

	section := sectionDTO.ToModel(claims.UserID)
	section.BAID = &baID
	section.OrganizationId = claims.OrganizationId

	if err := h.baService.CreateSection(section); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "BA section created successfully",
		"data":    sectionDTO,
	})
}

// UpdateBASection
// @Summary Update BA Section
// @Description Update an existing BA Section
// @Tags BASections
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param ba_id path int true "BA ID"
// @Param section_id path int true "Section ID"
// @Param section body dto.BASectionDTO true "BA Section Data"
// @Success 200 {object} map[string]interface{} "Updated BA Section"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "BA Section Not Found"
// @Failure 500 {object} map[string]string
// @Router /api/ba/{ba_id}/sections/{section_id} [put]
func (h *BAHandler) UpdateBASection(c *gin.Context) {
	baIDStr := c.Param("ba_id")
	baID, err := strconv.Atoi(baIDStr)
	if err != nil || baID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid BA ID"})
		return
	}

	sectionIDStr := c.Param("section_id")
	sectionID, err := strconv.Atoi(sectionIDStr)
	if err != nil || sectionID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Section ID"})
		return
	}

	var sectionDTO dto.BASectionDTO
	if err := c.ShouldBindJSON(&sectionDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	claims, err := helpers.GetClaims(c)
	if err != nil {
		return
	}

	existingSection, err := h.baService.GetBASectionByID(sectionID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "BA section not found"})
		return
	}

	updatedSection := sectionDTO.ToModelForUpdate(existingSection, claims.UserID)
	updatedSection.BAID = &baID
	updatedSection.OrganizationId = claims.OrganizationId

	if err := h.baService.UpdateSection(updatedSection); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "BA section updated successfully",
		"data":    sectionDTO,
	})
}

// DeleteBASection
// @Summary Delete BA Section
// @Description Delete an existing BA Section
// @Tags BASections
// @Security BearerAuth
// @Param ba_id path int true "BA ID"
// @Param section_id path int true "Section ID"
// @Produce json
// @Success 204 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "BA Section Not Found"
// @Failure 500 {object} map[string]string
// @Router /api/ba/{ba_id}/sections/{section_id} [delete]
func (h *BAHandler) DeleteBASection(c *gin.Context) {
	baIDStr := c.Param("ba_id")
	baID, err := strconv.Atoi(baIDStr)
	if err != nil || baID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid BA ID"})
		return
	}

	sectionIDStr := c.Param("section_id")
	sectionID, err := strconv.Atoi(sectionIDStr)
	if err != nil || sectionID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Section ID"})
		return
	}

	claims, err := helpers.GetClaims(c)
	if err != nil {
		return
	}

	if err := h.baService.DeleteSection(sectionID, baID, claims.UserID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// CreateBADetail
// @Summary Create BA Detail
// @Description Create a new BA Detail
// @Tags BADetails
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param ba_id path int true "BA ID"
// @Param section_id path int true "Section ID"
// @Param detail body dto.BADetailDTO true "BA Detail Data"
// @Success 201 {object} map[string]interface{} "Created BA Detail"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string
// @Router /api/ba/{ba_id}/sections/{section_id}/details [post]
func (h *BAHandler) CreateBADetail(c *gin.Context) {
	baIDStr := c.Param("ba_id")
	baID, err := strconv.Atoi(baIDStr)
	if err != nil || baID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid BA ID"})
		return
	}

	sectionIDStr := c.Param("section_id")
	sectionID, err := strconv.Atoi(sectionIDStr)
	if err != nil || sectionID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Section ID"})
		return
	}

	var detailDTO dto.BADetailDTO
	if err := c.ShouldBindJSON(&detailDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	claims, err := helpers.GetClaims(c)
	if err != nil {
		return
	}

	detail := detailDTO.ToModel(claims.UserID)
	detail.SectionID = &sectionID
	detail.OrganizationId = claims.OrganizationId

	if err := h.baService.CreateDetail(detail, baID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "BA detail created successfully",
		"data":    detailDTO,
	})
}

// UpdateBADetail
// @Summary Update BA Detail
// @Description Update an existing BA Detail
// @Tags BADetails
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param ba_id path int true "BA ID"
// @Param section_id path int true "Section ID"
// @Param detail_id path int true "Detail ID"
// @Param detail body dto.BADetailDTO true "BA Detail Data"
// @Success 200 {object} map[string]interface{} "Updated BA Detail"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "BA Detail Not Found"
// @Failure 500 {object} map[string]string
// @Router /api/ba/{ba_id}/sections/{section_id}/details/{detail_id} [put]
func (h *BAHandler) UpdateBADetail(c *gin.Context) {
	baIDStr := c.Param("ba_id")
	baID, err := strconv.Atoi(baIDStr)
	if err != nil || baID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid BA ID"})
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

	var detailDTO dto.BADetailDTO
	if err := c.ShouldBindJSON(&detailDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	claims, err := helpers.GetClaims(c)
	if err != nil {
		return
	}

	existingDetail, err := h.baService.GetBADetailByID(detailID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "BA detail not found"})
		return
	}

	updatedDetail := detailDTO.ToModelForUpdate(existingDetail, claims.UserID)
	updatedDetail.SectionID = &sectionID
	updatedDetail.OrganizationId = claims.OrganizationId

	if err := h.baService.UpdateDetail(updatedDetail, baID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "BA detail updated successfully",
		"data":    updatedDetail,
	})
}

// DeleteBADetail
// @Summary Delete BA Detail
// @Description Delete an existing BA Detail
// @Tags BADetails
// @Security BearerAuth
// @Param ba_id path int true "BA ID"
// @Param section_id path int true "Section ID"
// @Param detail_id path int true "Detail ID"
// @Produce json
// @Success 204 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "BA Detail Not Found"
// @Failure 500 {object} map[string]string
// @Router /api/ba/{ba_id}/sections/{section_id}/details/{detail_id} [delete]
func (h *BAHandler) DeleteBADetail(c *gin.Context) {
	baIDStr := c.Param("ba_id")
	baID, err := strconv.Atoi(baIDStr)
	if err != nil || baID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid BA ID"})
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

	claims, err := helpers.GetClaims(c)
	if err != nil {
		return
	}

	if err := h.baService.DeleteDetail(detailID, baID, claims.UserID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// CreateBAProgress
// @Summary Create BA Progress
// @Description Create a new BA Progress entry
// @Tags BAProgress
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param ba_id path int true "BA ID"
// @Param detail_id path int true "Detail ID"
// @Param progress body dto.BAProgressDTO true "BA Progress Data"
// @Success 201 {object} map[string]interface{} "Created BA Progress"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string
// @Router /api/ba/{ba_id}/details/{detail_id}/progress [post]
func (h *BAHandler) CreateBAProgress(c *gin.Context) {
	baIDStr := c.Param("ba_id")
	baID, err := strconv.Atoi(baIDStr)
	if err != nil || baID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid BA ID"})
		return
	}

	detailIDStr := c.Param("detail_id")
	detailID, err := strconv.Atoi(detailIDStr)
	if err != nil || detailID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Detail ID"})
		return
	}

	var progressDTO dto.BAProgressDTO
	if err := c.ShouldBindJSON(&progressDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	claims, err := helpers.GetClaims(c)
	if err != nil {
		return
	}

	progress := progressDTO.ToModel(claims.UserID)
	progress.DetailId = detailID
	progress.OrganizationId = claims.OrganizationId

	if err := h.baService.CreateProgress(progress, baID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "BA progress created successfully",
		"data":    progressDTO,
	})
}

// UpdateBAProgress
// @Summary Update BA Progress
// @Description Update an existing BA Progress entry
// @Tags BAProgress
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param ba_id path int true "BA ID"
// @Param detail_id path int true "Detail ID"
// @Param progress_id path int true "Progress ID"
// @Param progress body dto.BAProgressDTO true "BA Progress Data"
// @Success 200 {object} map[string]interface{} "Updated BA Progress"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "BA Progress Not Found"
// @Failure 500 {object} map[string]string
// @Router /api/ba/{ba_id}/details/{detail_id}/progress/{progress_id} [put]
func (h *BAHandler) UpdateBAProgress(c *gin.Context) {
	baIDStr := c.Param("ba_id")
	baID, err := strconv.Atoi(baIDStr)
	if err != nil || baID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid BA ID"})
		return
	}

	detailIDStr := c.Param("detail_id")
	detailID, err := strconv.Atoi(detailIDStr)
	if err != nil || detailID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Detail ID"})
		return
	}

	progressIDStr := c.Param("progress_id")
	progressID, err := strconv.Atoi(progressIDStr)
	if err != nil || progressID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Progress ID"})
		return
	}

	var progressDTO dto.BAProgressDTO
	if err := c.ShouldBindJSON(&progressDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	claims, err := helpers.GetClaims(c)
	if err != nil {
		return
	}

	existingProgress, err := h.baService.GetBAProgressByID(progressID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "BA progress not found"})
		return
	}

	updatedProgress := progressDTO.ToModelForUpdate(existingProgress, claims.UserID)
	updatedProgress.DetailId = detailID
	updatedProgress.OrganizationId = claims.OrganizationId

	if err := h.baService.UpdateProgress(updatedProgress, baID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "BA progress updated successfully",
		"data":    progressDTO,
	})
}

// DeleteBAProgress
// @Summary Delete BA Progress
// @Description Delete an existing BA Progress entry
// @Tags BAProgress
// @Security BearerAuth
// @Param ba_id path int true "BA ID"
// @Param detail_id path int true "Detail ID"
// @Param progress_id path int true "Progress ID"
// @Produce json
// @Success 204 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "BA Progress Not Found"
// @Failure 500 {object} map[string]string
// @Router /api/ba/{ba_id}/details/{detail_id}/progress/{progress_id} [delete]
func (h *BAHandler) DeleteBAProgress(c *gin.Context) {
	baIDStr := c.Param("ba_id")
	baID, err := strconv.Atoi(baIDStr)
	if err != nil || baID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid BA ID"})
		return
	}

	detailIDStr := c.Param("detail_id")
	detailID, err := strconv.Atoi(detailIDStr)
	if err != nil || detailID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Detail ID"})
		return
	}

	progressIDStr := c.Param("progress_id")
	progressID, err := strconv.Atoi(progressIDStr)
	if err != nil || progressID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Progress ID"})
		return
	}

	claims, err := helpers.GetClaims(c)
	if err != nil {
		return
	}

	if err := h.baService.DeleteProgress(progressID, baID, claims.UserID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
