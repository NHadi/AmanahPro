package models

import (
	"time"
)

type GoogleDriveToken struct {
	TokenID      uint       `gorm:"primaryKey;autoIncrement"`                     // Primary key
	UserID       string     `gorm:"size:255;not null;unique"`                     // Unique user ID
	AccessToken  string     `gorm:"type:text;not null"`                           // Access token
	ExpiresIn    int64      `gorm:"not null"`                                     // Expiry time (in seconds)
	RefreshToken string     `gorm:"type:text;not null"`                           // Refresh token
	CreatedAt    time.Time  `gorm:"autoCreateTime"`                               // Automatically sets when the record is created
	UpdatedAt    time.Time  `gorm:"autoUpdateTime"`                               // Automatically sets when the record is updated
	DeletedAt    *time.Time `gorm:"column:DeletedAt" json:"deleted_at,omitempty"` // Soft delete column
}

func (GoogleDriveToken) TableName() string {
	return "GoogleDriveTokens" // Match the exact name from the database
}
