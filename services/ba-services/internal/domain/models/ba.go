package models

import (
	"encoding/json"
	"time"
)

// BA represents the BA table in the database
type BA struct {
	BAId           int        `gorm:"primaryKey;column:BAId;autoIncrement"`        // Primary key
	SphId          *int       `gorm:"column:SphId;null"`                           // Reference to SPH ID
	ProjectId      *int       `gorm:"column:ProjectId;null"`                       // Reference to Project ID
	ProjectName    *string    `gorm:"column:ProjectName;type:varchar(250);null"`   // Project Name
	BADate         CustomDate `gorm:"column:BADate;type:date;not null"`            // BA Date
	BASubject      string     `gorm:"column:BASubject;type:varchar(255);not null"` // BA Subject
	RecepientName  *string    `gorm:"column:RecepientName;type:varchar(255);null"` // Recepient Name
	CreatedBy      *int       `gorm:"column:CreatedBy;null"`                       // Created by user ID
	CreatedAt      *time.Time `gorm:"column:CreatedAt;autoCreateTime"`             // Creation timestamp
	UpdatedBy      *int       `gorm:"column:UpdatedBy;null"`                       // Updated by user ID
	UpdatedAt      *time.Time `gorm:"column:UpdatedAt;autoUpdateTime;null"`        // Update timestamp
	DeletedBy      *int       `gorm:"column:DeletedBy;null"`                       // Deleted by user ID
	DeletedAt      *time.Time `gorm:"column:DeletedAt;index;null"`                 // Deletion timestamp
	OrganizationId *int       `gorm:"column:OrganizationId;null"`                  // Organization ID

	Sections []BASection `gorm:"foreignKey:BAID;references:BAId"` // Relationship to BASection
}

// TableName specifies the table name for BA
func (BA) TableName() string {
	return "BA"
}

func (b BA) MarshalJSON() ([]byte, error) {
	type Alias BA
	return json.Marshal(&struct {
		Sections []BASection `json:"Sections"`
		Alias
	}{
		Sections: b.Sections, // `BASection` handles its own custom marshaling
		Alias:    (Alias)(b),
	})
}
