package handlers

import (
	"encoding/json"
	"net/http"

	"AmanahPro/services/user-management/internal/application/services"
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
// @Param user body models.User true "User"
// @Success 201 {object} models.User
// @Router /users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userData struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	user, err := h.userService.CreateUser(userData.Username, userData.Email, userData.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
