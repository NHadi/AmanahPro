package models

import (
	"time"
)

// RoleMenu model
type RoleMenu struct {
	RoleID     int       `gorm:"column:role_id;primaryKey"`
	MenuID     int       `gorm:"column:menu_id;primaryKey"`
	Permission string    `gorm:"type:varchar(10);not null"`
	AssignedAt time.Time `gorm:"autoCreateTime"`

	// Relationships
	Role *Role `gorm:"foreignKey:RoleID"`
	Menu *Menu `gorm:"foreignKey:MenuID"`
}

func (RoleMenu) TableName() string {
	return "RoleMenus" // Match the exact name from the database
}
