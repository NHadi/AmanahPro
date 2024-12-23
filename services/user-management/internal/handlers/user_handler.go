package handlers

import (
	"AmanahPro/services/user-management/internal/application/services"
	"net/http"
	"strconv"

	"AmanahPro/services/user-management/internal/helpers"

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
// @Description Create a new user with provided details and assigns a default role
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
		RoleID   int    `json:"roleID"` // Role ID to assign
	}

	// Bind JSON input to struct
	if err := c.ShouldBindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	claims, _ := helpers.GetClaims(c)

	// Create user through service
	user, err := h.userService.CreateUser(userData.Username, userData.Email, userData.Password, claims.OrganizationId, userData.RoleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with created user data
	c.JSON(http.StatusCreated, user)
}

// LoadUsersByOrganization godoc
// @Summary Load users by organization ID
// @Description Get a list of users associated with a specific organization
// @Security BearerAuth
// @Tags Users
// @Accept json
// @Produce json
// @Param organization_id path int true "Organization ID"
// @Success 200 {array} models.User
// @Failure 400 {object} map[string]string
// @Router /api/users/organizations [get]
func (h *UserHandler) LoadUsersByOrganization(c *gin.Context) {
	claims, _ := helpers.GetClaims(c)

	// Fetch users by organization ID
	users, err := h.userService.LoadByOrganizationID(*claims.OrganizationId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with the list of users
	c.JSON(http.StatusOK, users)
}

// UpdateUser godoc
// @Summary Update an existing user
// @Description Update a user's details and re-assign their role if needed
// @Security BearerAuth
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.User true "User Data"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]string
// @Router /api/users/{user_id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	var userData struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
		RoleID   int    `json:"roleID"`
	}

	// Convert the user_id parameter from string to int
	userIDStr := c.Param("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Bind JSON input to struct
	if err := c.ShouldBindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	claims, _ := helpers.GetClaims(c)

	// Call the service to update the user
	updatedUser, err := h.userService.UpdateUser(userID, userData.Username, userData.Email, userData.Password, claims.OrganizationId, userData.RoleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with the updated user data
	c.JSON(http.StatusOK, updatedUser)
}
