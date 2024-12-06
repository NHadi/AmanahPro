package models

import "time"

type ProjectUser struct {
	ID             int        `gorm:"primaryKey;autoIncrement"` // Primary key
	ProjectID      int        `gorm:"not null"`                 // Foreign key to Projects
	UserID         int        `gorm:"not null"`                 // Foreign key to Users
	Role           *string    `gorm:"type:nvarchar(50);null"`   // Role of the user in the project
	CreatedAt      *time.Time `gorm:"autoCreateTime"`           // Creation timestamp
	CreatedBy      *int       `gorm:"null"`                     // Created by user ID
	UpdatedBy      *int       `gorm:"null"`                     // Updated by user ID
	UpdatedAt      *time.Time `gorm:"autoUpdateTime"`           // Update timestamp
	DeletedBy      *int       `gorm:"null"`                     // Deleted by user ID
	DeletedAt      *time.Time `gorm:"index;null"`               // Deletion timestamp
	OrganizationID *int       `gorm:"null"`                     // Organization ID

	// Relationships
	Project Project `gorm:"foreignKey:ProjectID;references:ProjectID"` // Many-to-One relationship with Project
}

// TableName specifies the table name for ProjectUser
func (ProjectUser) TableName() string {
	return "ProjectUser"
}
