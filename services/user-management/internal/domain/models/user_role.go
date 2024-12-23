package models

import (
	"time"
)

// UserRole model
type UserRole struct {
	UserID     int       `gorm:"column:user_id;primaryKey"`
	RoleID     int       `gorm:"column:role_id;primaryKey"`
	AssignedAt time.Time `gorm:"column:assigned_at;autoCreateTime"`

	// Relationships
	User *User `gorm:"foreignKey:UserID"`
	Role *Role `gorm:"foreignKey:RoleID"`
}

// TableName overrides the table name in GORM
func (UserRole) TableName() string {
	return "UserRoles"
}
