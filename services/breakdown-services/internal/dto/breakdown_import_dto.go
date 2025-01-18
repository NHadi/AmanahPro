package dto

// SphDTO represents the data transfer object for SPH
type BreakdownImportDTO struct {
	ProjectId *int    `form:"ProjectId" binding:"required"` // Required field
	Subject   *string `form:"Subject"`                      // Optional
	Location  *string `json:"Location"`
}
