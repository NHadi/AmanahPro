package models

import (
	"time"
)

type UserRole struct {
	UserID     int       `gorm:"column:user_id;primaryKey;autoIncrement"`
	RoleID     int       `gorm:"column:role_id;primaryKey"`
	AssignedAt time.Time `gorm:"default:GETDATE()"`
}
