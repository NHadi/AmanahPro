package models

import (
	"time"
)

// Project represents the Projects table in the database
type Project struct {
	ProjectID      int         `gorm:"column:ProjectID;primaryKey;autoIncrement"`                 // Maps to [ProjectID]
	ProjectName    string      `gorm:"column:ProjectName;type:varchar(255);not null"`             // Maps to [ProjectName]
	Location       *string     `gorm:"column:Location;type:varchar(255);null"`                    // Maps to [Location]
	StartDate      *CustomDate `gorm:"column:StartDate;type:date;null"`                           // Maps to [StartDate]
	EndDate        *CustomDate `gorm:"column:EndDate;type:date;null"`                             // Maps to [EndDate]
	Description    *string     `gorm:"column:Description;type:text;null"`                         // Maps to [Description]
	Status         *string     `gorm:"column:Status;type:varchar(20);default:'in-progress';null"` // Maps to [Status]
	CreatedBy      *int        `gorm:"column:CreatedBy;null"`                                     // Maps to [CreatedBy]
	CreatedAt      *time.Time  `gorm:"column:CreatedAt;autoCreateTime"`                           // Maps to [CreatedAt]
	UpdatedBy      *int        `gorm:"column:UpdatedBy;null"`                                     // Maps to [UpdatedBy]
	UpdatedAt      *time.Time  `gorm:"column:UpdatedAt;autoUpdateTime"`                           // Maps to [UpdatedAt]
	DeletedBy      *int        `gorm:"column:DeletedBy;null"`                                     // Maps to [DeletedBy]
	DeletedAt      *time.Time  `gorm:"column:DeletedAt;null"`                                     // Maps to [DeletedAt]
	OrganizationID *int        `gorm:"column:OrganizationID;null"`                                // Maps to [OrganizationID]
}

// TableName specifies the table name for Project
func (Project) TableName() string {
	return "Projects"
}
