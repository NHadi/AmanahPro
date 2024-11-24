package repositories

import (
	"AmanahPro/services/user-management/internal/domain/models"

	"gorm.io/gorm"
)

type googleDriveTokenRepository struct {
	db *gorm.DB
}

func NewGoogleDriveTokenRepository(db *gorm.DB) *googleDriveTokenRepository {
	return &googleDriveTokenRepository{db: db}
}

func (r *googleDriveTokenRepository) SaveToken(token *models.GoogleDriveToken) error {
	return r.db.Save(token).Error
}

func (r *googleDriveTokenRepository) GetTokenByUserID(userID string) (*models.GoogleDriveToken, error) {
	var token models.GoogleDriveToken
	err := r.db.Where("UserID = ?", userID).First(&token).Error
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (r *googleDriveTokenRepository) DeleteTokenByUserID(userID string) error {
	return r.db.Where("UserID = ?", userID).Delete(&models.GoogleDriveToken{}).Error
}
