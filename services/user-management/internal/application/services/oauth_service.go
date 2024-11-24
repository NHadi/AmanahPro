package services

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"time"

	"AmanahPro/services/user-management/internal/domain/models"
	"AmanahPro/services/user-management/internal/domain/repositories"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
)

type OAuthService struct {
	config          *oauth2.Config
	tokenRepository repositories.GoogleDriveTokenRepository
}

func NewOAuthService(ClientID, ClientSecret string, tokenRepo repositories.GoogleDriveTokenRepository) *OAuthService {
	config := &oauth2.Config{
		ClientID:     ClientID,
		ClientSecret: ClientSecret,
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost:8080/user-management/api/oauth2/callback",
		Scopes:       []string{"https://www.googleapis.com/auth/drive"},
	}
	return &OAuthService{
		config:          config,
		tokenRepository: tokenRepo,
	}
}

func (s *OAuthService) GetAuthURL() string {
	return s.config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
}

func (s *OAuthService) HandleCallback(code string, userID string) error {

	token, err := s.config.Exchange(context.Background(), code)
	if err != nil {
		return err
	}

	t := &models.GoogleDriveToken{
		UserID:       userID,
		AccessToken:  token.AccessToken,
		ExpiresIn:    token.Expiry.Unix(),
		RefreshToken: token.RefreshToken,
	}
	return s.tokenRepository.SaveToken(t)
}

func (s *OAuthService) GetToken(userID string) (*oauth2.Token, error) {

	tokenModel, err := s.tokenRepository.GetTokenByUserID(userID)
	if err != nil {
		return nil, err
	}

	return &oauth2.Token{
		AccessToken:  tokenModel.AccessToken,
		TokenType:    "Bearer",
		RefreshToken: tokenModel.RefreshToken,
		Expiry:       time.Unix(tokenModel.ExpiresIn, 0),
	}, nil
}

func (s *OAuthService) RefreshTokenIfNeeded(userID string) (*oauth2.Token, error) {
	tokenModel, err := s.tokenRepository.GetTokenByUserID(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve token: %w", err)
	}

	token := &oauth2.Token{
		AccessToken:  tokenModel.AccessToken,
		TokenType:    "Bearer",
		RefreshToken: tokenModel.RefreshToken,
		Expiry:       time.Unix(tokenModel.ExpiresIn, 0),
	}

	if token.Valid() {
		return token, nil
	}

	// Refresh token
	tokenSource := s.config.TokenSource(context.Background(), token)
	newToken, err := tokenSource.Token()
	if err != nil {
		return nil, fmt.Errorf("failed to refresh token: %w", err)
	}

	tokenModel.AccessToken = newToken.AccessToken
	tokenModel.ExpiresIn = newToken.Expiry.Unix()
	tokenModel.RefreshToken = newToken.RefreshToken
	err = s.tokenRepository.SaveToken(tokenModel)
	if err != nil {
		return nil, fmt.Errorf("failed to save refreshed token: %w", err)
	}

	return newToken, nil
}

func (s *OAuthService) DownloadFileContent(userID, fileID string) ([]byte, error) {
	// Get a valid token for the user
	token, err := s.RefreshTokenIfNeeded(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve or refresh token: %w", err)
	}

	// Create a client using the token
	client := s.config.Client(context.Background(), token)

	// Initialize Google Drive service
	driveService, err := drive.New(client)
	if err != nil {
		return nil, fmt.Errorf("failed to create Drive client: %w", err)
	}

	// Download the file content
	response, err := driveService.Files.Get(fileID).Download()
	if err != nil {
		return nil, fmt.Errorf("failed to download file: %w", err)
	}
	defer response.Body.Close()

	// Read file content into a byte slice
	content, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read file content: %w", err)
	}

	return content, nil
}

func (s *OAuthService) UploadFileToGoogleDrive(userID, organizationID, fileName string, fileData []byte) (string, error) {
	// Retrieve or refresh token if needed
	token, err := s.RefreshTokenIfNeeded(userID)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve or refresh token: %w", err)
	}

	// Create an HTTP client with the token
	client := s.config.Client(context.Background(), token)

	// Initialize Google Drive service
	driveService, err := drive.New(client)
	if err != nil {
		return "", fmt.Errorf("failed to create Drive client: %w", err)
	}

	// Get the current year and month
	now := time.Now()
	year := fmt.Sprintf("%d", now.Year())
	month := fmt.Sprintf("%02d", now.Month())

	// Ensure folder hierarchy exists
	rootFolderID, err := s.ensureFolder(driveService, "AmanahPro", "root")
	if err != nil {
		return "", fmt.Errorf("failed to ensure root folder: %w", err)
	}

	orgFolderID, err := s.ensureFolder(driveService, organizationID, rootFolderID)
	if err != nil {
		return "", fmt.Errorf("failed to ensure organization folder: %w", err)
	}

	yearFolderID, err := s.ensureFolder(driveService, year, orgFolderID)
	if err != nil {
		return "", fmt.Errorf("failed to ensure year folder: %w", err)
	}

	monthFolderID, err := s.ensureFolder(driveService, month, yearFolderID)
	if err != nil {
		return "", fmt.Errorf("failed to ensure month folder: %w", err)
	}

	// Create a Google Drive file object
	driveFile := &drive.File{
		Name:    fileName,
		Parents: []string{monthFolderID}, // Upload file to the month folder
	}

	// Create a reader for the file data
	fileReader := bytes.NewReader(fileData)

	// Upload the file
	uploadedFile, err := driveService.Files.Create(driveFile).Media(fileReader).Do()
	if err != nil {
		return "", fmt.Errorf("failed to upload file to Google Drive: %w", err)
	}

	return uploadedFile.Id, nil
}

func (s *OAuthService) ensureFolder(driveService *drive.Service, folderName, parentID string) (string, error) {
	// Query to check if the folder already exists
	query := fmt.Sprintf("name='%s' and mimeType='application/vnd.google-apps.folder' and trashed=false", folderName)
	if parentID != "root" {
		query += fmt.Sprintf(" and '%s' in parents", parentID)
	}

	// List folders with the specified name under the parent
	fileList, err := driveService.Files.List().Q(query).Fields("files(id, name)").Do()
	if err != nil {
		return "", fmt.Errorf("failed to list folders: %w", err)
	}

	// If the folder exists, return its ID
	if len(fileList.Files) > 0 {
		return fileList.Files[0].Id, nil
	}

	// Folder does not exist, create it
	folder := &drive.File{
		Name:     folderName,
		MimeType: "application/vnd.google-apps.folder",
	}
	if parentID != "root" {
		folder.Parents = []string{parentID}
	}

	// Create the folder in Google Drive
	createdFolder, err := driveService.Files.Create(folder).Fields("id").Do()
	if err != nil {
		return "", fmt.Errorf("failed to create folder: %w", err)
	}

	// Return the ID of the newly created folder
	return createdFolder.Id, nil
}
