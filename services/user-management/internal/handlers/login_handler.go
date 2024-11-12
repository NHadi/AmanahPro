package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"AmanahPro/services/user-management/internal/application/services"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

type LoginHandler struct {
	userService *services.UserService
}

func NewLoginHandler(userService *services.UserService) *LoginHandler {
	return &LoginHandler{userService: userService}
}

// LoginRequest represents the login request payload.
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse represents the response payload for a successful login.
type LoginResponse struct {
	Token string `json:"token"`
}

// JWTClaims defines the custom claims for JWT.
type JWTClaims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

// Login godoc
// @Summary User login
// @Description Logs in a user and returns a JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param login body LoginRequest true "Login credentials"
// @Success 200 {object} LoginResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /login [post]
func (h *LoginHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logrus.WithError(err).Warn("Invalid login request payload")
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Authenticate the user using userService
	user, err := h.userService.Authenticate(req.Username, req.Password)
	if err != nil {
		logrus.WithError(err).Warn("Unauthorized login attempt")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Create JWT token with user ID in claims
	claims := &JWTClaims{
		UserID: user.UserID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("your_secret_key"))
	if err != nil {
		logrus.WithError(err).Error("Failed to sign JWT token")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	logrus.WithField("user_id", user.UserID).Info("User logged in successfully")

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(LoginResponse{Token: tokenString}); err != nil {
		logrus.WithError(err).Error("Failed to encode login response")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
