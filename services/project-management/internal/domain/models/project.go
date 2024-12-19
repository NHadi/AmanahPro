package models

import (
	"time"
)

type Project struct {
	ProjectID      int         `gorm:"column:ProjectID;primaryKey;autoIncrement"`                 // Primary key
	ProjectName    string      `gorm:"column:ProjectName;type:varchar(255);not null"`             // Project name
	Location       *string     `gorm:"column:Location;type:varchar(255);null"`                    // Location
	StartDate      *CustomDate `gorm:"column:StartDate;type:date;null"`                           // Start date
	EndDate        *CustomDate `gorm:"column:EndDate;type:date;null"`                             // End date
	Description    *string     `gorm:"column:Description;type:text;null"`                         // Description
	Status         *string     `gorm:"column:Status;type:varchar(20);default:'in-progress';null"` // Status
	PO_SPHDate     *CustomDate `gorm:"column:PO_SPHDate;type:date;null"`                          // PO/SPH date
	SPH            *float64    `gorm:"column:SPH;type:decimal(18,2;null"`                         // SPH
	Termin         *float64    `gorm:"column:Termin;decimal(18,2;null"`                           // Termin
	TotalSPK       *float64    `gorm:"column:TotalSPK;type:decimal(18,2);null"`
	CreatedBy      *int        `gorm:"column:CreatedBy;null"`           // Created by
	CreatedAt      *time.Time  `gorm:"column:CreatedAt;autoCreateTime"` // Created at
	UpdatedBy      *int        `gorm:"column:UpdatedBy;null"`           // Updated by
	UpdatedAt      *time.Time  `gorm:"column:UpdatedAt;autoUpdateTime"` // Updated at
	DeletedBy      *int        `gorm:"column:DeletedBy;null"`           // Deleted by
	DeletedAt      *time.Time  `gorm:"column:DeletedAt;null"`           // Deleted at
	OrganizationID *int        `gorm:"column:OrganizationID;null"`      // Organization ID

	// Relationships
	ProjectRecap []ProjectRecap `gorm:"foreignKey:ProjectID;references:ProjectID"` // One-to-Many relationship with ProjectRecap
	ProjectUsers []ProjectUser  `gorm:"foreignKey:ProjectID;references:ProjectID"` // One-to-Many relationship with ProjectUser
}

// TableName specifies the table name for Project
func (Project) TableName() string {
	return "Projects"
}
