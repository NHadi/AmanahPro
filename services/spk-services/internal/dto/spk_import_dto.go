package dto

// SphDTO represents the data transfer object for SPH
type SpkImportDTO struct {
	ProjectId *int    `form:"ProjectId" binding:"required"` // Required field
	Subject   *string `form:"Subject"`                      // Optional
	Mandor    *string `json:"Mandor"`
}
