package repositories

import "AmanahPro/services/user-management/internal/domain/models"

type GoogleDriveTokenRepository interface {
	SaveToken(token *models.GoogleDriveToken) error
	GetTokenByUserID(userID string) (*models.GoogleDriveToken, error)
	DeleteTokenByUserID(userID string) error
}
