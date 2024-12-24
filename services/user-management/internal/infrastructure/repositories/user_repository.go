package repositories

import (
	"AmanahPro/services/user-management/internal/domain/models"
	"AmanahPro/services/user-management/internal/domain/repositories"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db *gorm.DB) repositories.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

// Update updates an existing user in the database
func (r *userRepository) Update(user *models.User) error {
	// Update user details in the database
	return r.db.Save(user).Error
}

func (r *userRepository) FindByID(id int) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, "user_id = ?", id).Error
	return &user, err
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Preload("UserRoles.Role").First(&user, "Email = ? OR username = ?", email, email).Error
	return &user, err
}

func (r *userRepository) FindAll() ([]models.User, error) {
	var users []models.User
	err := r.db.Preload("Organization").Preload("UserRoles.Role").Find(&users).Error
	return users, err
}

func (r *userRepository) FindByOrganizationID(organizationId int) ([]models.User, error) {
	var users []models.User
	err := r.db.
		Preload("Organization").   // Preload related Organization
		Preload("UserRoles.Role"). // Preload UserRoles and nested Role
		Where("organization_id = ?", organizationId).
		Find(&users).Error
	return users, err
}

func (r *userRepository) DeleteByID(id int) error {
	return r.db.Delete(&models.User{}, "user_id = ?", id).Error
}
