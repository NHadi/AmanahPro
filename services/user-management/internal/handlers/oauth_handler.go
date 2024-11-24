package handlers

import (
	"fmt"
	"net/http"

	"AmanahPro/services/user-management/internal/application/services"

	jwtModels "github.com/NHadi/AmanahPro-common/models"
	"github.com/gin-gonic/gin"
)

type OAuthHandler struct {
	oauthService *services.OAuthService
}

// NewOAuthHandler creates a new instance of OAuthHandler
func NewOAuthHandler(oauthService *services.OAuthService) *OAuthHandler {
	return &OAuthHandler{oauthService: oauthService}
}

// Authorize handles the request to get the Google OAuth2 authorization URL
// @Summary Get Google OAuth2 Authorization URL
// @Security BearerAuth
// @Description Returns the Google OAuth2 authorization URL for the client to redirect the user
// @Tags OAuth2
// @Produce json
// @Success 200 {object} map[string]string "auth_url"
// @Router /api/oauth2/authorize [get]
func (h *OAuthHandler) Authorize(c *gin.Context) {
	url := h.oauthService.GetAuthURL()
	c.JSON(http.StatusOK, gin.H{"auth_url": url})
}

// Callback handles the Google OAuth2 callback and saves the token
// @Summary Handle OAuth2 Callback
// @Security BearerAuth
// @Description Handles the Google OAuth2 callback and saves the token in the database
// @Tags OAuth2
// @Produce json
// @Param code query string true "Authorization Code"
// @Success 200 {object} map[string]string "message"
// @Failure 400 {object} map[string]string "error"
// @Failure 500 {object} map[string]string "error"
// @Router /api/oauth2/callback [get]
func (h *OAuthHandler) Callback(c *gin.Context) {

	// Extract claims from the JWT token
	userClaims, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	claims, ok := userClaims.(*jwtModels.JWTClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		return
	}

	userID := claims.Username
	code := c.Query("code")

	if code == "" || userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing code or user_id"})
		return
	}

	err := h.oauthService.HandleCallback(code, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OAuth successful"})
}

// UploadFileHandler uploads a file to Google Drive
// @Summary Upload File
// @Security BearerAuth
// @Description Uploads a file to Google Drive under a specific folder hierarchy
// @Tags File
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File to Upload"
// @Success 200 {object} map[string]string "file_id"
// @Failure 400 {object} map[string]string "error"
// @Failure 500 {object} map[string]string "error"
// @Router /api/upload [post]
func (h *OAuthHandler) UploadFileHandler(c *gin.Context) {

	// Extract claims from the JWT token
	userClaims, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	claims, ok := userClaims.(*jwtModels.JWTClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		return
	}

	userID := claims.Username
	organizationID := fmt.Sprintf("%v", claims.OrganizationId)

	// Get file from form-data
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read file"})
		return
	}
	defer file.Close()

	// Read the file content
	fileData := make([]byte, header.Size)
	_, err = file.Read(fileData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file content"})
		return
	}

	// Call the upload function
	fileID, err := h.oauthService.UploadFileToGoogleDrive(userID, organizationID, header.Filename, fileData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"file_id": fileID})
}

// DownloadFileHandler downloads a file from Google Drive
// @Summary Download File
// @Security BearerAuth
// @Description Downloads a file from Google Drive using its file ID
// @Tags File
// @Produce json
// @Param file_id query string true "File ID"
// @Success 200 {object} map[string]string "file_content"
// @Failure 400 {object} map[string]string "error"
// @Failure 500 {object} map[string]string "error"
// @Router /api/download [get]
func (h *OAuthHandler) DownloadFileHandler(c *gin.Context) {

	userClaims, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	claims, ok := userClaims.(*jwtModels.JWTClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		return
	}

	userID := claims.Username
	fileID := c.Query("file_id")
	if userID == "" || fileID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing user_id or file_id"})
		return
	}

	// Call the download function
	fileContent, err := h.oauthService.DownloadFileContent(userID, fileID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"file_content": string(fileContent)})
}
