package models

import (
	"time"
)

// Role model
type Role struct {
	RoleID      int       `gorm:"column:role_id;primaryKey;autoIncrement"`
	Name        string    `gorm:"type:varchar(100);unique;not null"`
	Description string    `gorm:"type:text"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`

	// Relationships
	UserRoles []UserRole `gorm:"foreignKey:RoleID"`
	RoleMenus []RoleMenu `gorm:"foreignKey:RoleID"`
}
