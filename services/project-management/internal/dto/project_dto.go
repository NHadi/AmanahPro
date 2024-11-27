package dto

import "time"

type ProjectDTO struct {
	ProjectID      int              `json:"ProjectID"`
	ProjectName    string           `json:"ProjectName"`
	Location       string           `json:"Location"`
	Status         string           `json:"Status"`
	OrganizationID *int             `json:"OrganizationID"`
	Recap          *ProjectRecapDTO `json:"Recap"`    // Nested Recap object
	KeyUsers       []ProjectUserDTO `json:"KeyUsers"` // Nested KeyUsers array
	CreatedAt      time.Time        `json:"CreatedAt"`
	UpdatedAt      time.Time        `json:"UpdatedAt"`
}
