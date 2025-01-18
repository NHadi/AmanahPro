package dto

// SphDTO represents the data transfer object for SPH
type BAImportDTO struct {
	ProjectId     *int    `form:"ProjectId" binding:"required"` // Required field
	Subject       *string `form:"Subject"`                      // Optional
	RecepientName *string `json:"RecepientName"`
}
