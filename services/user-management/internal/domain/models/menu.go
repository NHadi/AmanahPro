package models

import (
	"time"
)

// Menu model
type Menu struct {
	MenuID    int       `gorm:"column:menu_id;primaryKey;autoIncrement"`
	ParentID  *int      `gorm:"column:parent_id"`
	Name      string    `gorm:"type:nvarchar(50);not null"`
	Path      string    `gorm:"type:nvarchar(100);not null"`
	Icon      string    `gorm:"type:nvarchar(50)"`
	Order     *int      `gorm:"column:order"`
	CreatedAt time.Time `gorm:"autoCreateTime"`

	// Relationships
	Children  []Menu     `gorm:"foreignKey:ParentID"`
	RoleMenus []RoleMenu `gorm:"foreignKey:MenuID"`
}
