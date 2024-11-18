package models

import (
	"time"
)

type RoleMenu struct {
	RoleID     int       `gorm:"column:role_id;primaryKey"`
	MenuID     int       `gorm:"column:menu_id;primaryKey"`
	Permission string    `gorm:"column:permission;type:varchar(10);not null"`
	AssignedAt time.Time `gorm:"column:assigned_at;default:GETDATE()"`
}

func (RoleMenu) TableName() string {
	return "RoleMenus" // Match the exact name from the database
}
