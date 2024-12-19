package dto

// SphDTO represents the data transfer object for SPH
type SphImportDTO struct {
	ProjectId     *int    `form:"ProjectId" binding:"required"` // Required field
	ProjectName   *string `form:"ProjectName"`                  // Optional
	Subject       *string `form:"Subject"`                      // Optional
	Location      *string `form:"Location"`                     // Optional
	RecepientName *string `form:"RecepientName"`                // Optional
}
