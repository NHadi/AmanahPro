package models

import (
	"time"
)

// User model
type User struct {
	UserID         int       `gorm:"column:user_id;primaryKey;autoIncrement"`
	Username       string    `gorm:"type:varchar(100);unique;not null"`
	Email          string    `gorm:"type:varchar(100);unique;not null"`
	Password       string    `gorm:"type:varchar(100);not null"`
	Status         bool      `gorm:"default:true"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
	OrganizationID *int      `gorm:"column:organization_id"`

	// Relationships
	Organization *SysOrganization `gorm:"foreignKey:OrganizationID"`
	UserRoles    []UserRole       `gorm:"foreignKey:UserID"`
}

// TableName overrides the table name in GORM
func (User) TableName() string {
	return "Users"
}
