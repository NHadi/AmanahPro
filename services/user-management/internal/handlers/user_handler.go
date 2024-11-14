package handlers

import (
	"AmanahPro/services/user-management/internal/application/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserHandler represents the HTTP handler for user operations
type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with provided details
// @Security BearerAuth
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.User true "User Data"
// @Success 201 {object} models.User
// @Failure 400 {object} map[string]string
// @Router /api/users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var userData struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Bind JSON input to struct
	if err := c.ShouldBindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Create user through service
	user, err := h.userService.CreateUser(userData.Username, userData.Email, userData.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with created user data
	c.JSON(http.StatusCreated, user)
}
