package models

import (
	"time"
)

type Menu struct {
	MenuID    int       `gorm:"column:menu_id;primaryKey;autoIncrement"`
	ParentID  *int      `gorm:"column:parent_id"` // Nullable to allow root items
	MenuName  string    `gorm:"column:name;type:nvarchar(50);not null"`
	Path      string    `gorm:"column:path;type:nvarchar(100);not null"`
	Icon      string    `gorm:"column:icon;type:nvarchar(50)"`
	Order     int       `gorm:"column:order;type:int"`
	CreatedAt time.Time `gorm:"column:created_at;default:GETDATE()"`
}
