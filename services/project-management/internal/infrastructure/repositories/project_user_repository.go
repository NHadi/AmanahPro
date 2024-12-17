package repositories

import (
	"AmanahPro/services/project-management/internal/domain/models"
	"AmanahPro/services/project-management/internal/domain/repositories"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type projectUserRepositoryImpl struct {
	db *gorm.DB
}

// NewProjectUserRepository creates a new instance of ProjectUserRepository
func NewProjectUserRepository(db *gorm.DB) repositories.ProjectUserRepository {
	return &projectUserRepositoryImpl{
		db: db,
	}
}

// Create inserts a new ProjectUser record into the database
func (r *projectUserRepositoryImpl) Create(user *models.ProjectUser) error {
	log.Printf("Creating ProjectUser: %+v", user)

	if err := r.db.Create(user).Error; err != nil {
		log.Printf("Failed to create ProjectUser: %v", err)
		return fmt.Errorf("failed to create ProjectUser: %w", err)
	}

	log.Printf("Successfully created ProjectUser: %+v", user)
	return nil
}

// Update modifies an existing ProjectUser record in the database
func (r *projectUserRepositoryImpl) Update(user *models.ProjectUser) error {
	log.Printf("Updating ProjectUser ID: %d", user.ID)

	if err := r.db.Save(user).Error; err != nil {
		log.Printf("Failed to update ProjectUser ID %d: %v", user.ID, err)
		return fmt.Errorf("failed to update ProjectUser: %w", err)
	}

	log.Printf("Successfully updated ProjectUser ID: %d", user.ID)
	return nil
}

// Delete removes a ProjectUser record from the database
func (r *projectUserRepositoryImpl) Delete(userID int) error {
	log.Printf("Deleting ProjectUser ID: %d", userID)

	if err := r.db.Delete(&models.ProjectUser{}, userID).Error; err != nil {
		log.Printf("Failed to delete ProjectUser ID %d: %v", userID, err)
		return fmt.Errorf("failed to delete ProjectUser: %w", err)
	}

	log.Printf("Successfully deleted ProjectUser ID: %d", userID)
	return nil
}

// GetByID retrieves a ProjectUser record by its ID
func (r *projectUserRepositoryImpl) GetByID(userID int) (*models.ProjectUser, error) {
	log.Printf("Retrieving ProjectUser by ID: %d", userID)

	var user models.ProjectUser
	if err := r.db.First(&user, userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("ProjectUser ID %d not found", userID)
			return nil, nil
		}
		log.Printf("Failed to retrieve ProjectUser ID %d: %v", userID, err)
		return nil, fmt.Errorf("failed to retrieve ProjectUser: %w", err)
	}

	log.Printf("Successfully retrieved ProjectUser: %+v", user)
	return &user, nil
}

// GetByProjectID retrieves all ProjectUser records for a specific ProjectID
func (r *projectUserRepositoryImpl) GetByProjectID(projectID int) ([]models.ProjectUser, error) {
	log.Printf("Retrieving ProjectUsers for ProjectID: %d", projectID)

	var users []models.ProjectUser
	if err := r.db.Where("ProjectID = ?", projectID).Find(&users).Error; err != nil {
		log.Printf("Failed to retrieve ProjectUsers for ProjectID %d: %v", projectID, err)
		return nil, fmt.Errorf("failed to retrieve ProjectUsers: %w", err)
	}

	log.Printf("Successfully retrieved %d ProjectUsers for ProjectID: %d", len(users), projectID)
	return users, nil
}

// UpdateCategoryByProjectUser updates the Category in ProjectFinancial where ProjectUserID and ProjectID match
func (r *projectUserRepositoryImpl) UpdateCategoryByProjectUser(projectID int, projectUserID int, userName string) error {
	log.Printf("Updating Category in ProjectFinancial for ProjectID: %d, ProjectUserID: %d to UserName: %s", projectID, projectUserID, userName)

	// Perform the update in the database
	result := r.db.Model(&models.ProjectFinancial{}).
		Where("ProjectID = ? AND ProjectUserID = ?", projectID, projectUserID).
		Update("Category", userName)

	// Check for errors during the update
	if result.Error != nil {
		log.Printf("Error updating Category: %v", result.Error)
		return fmt.Errorf("failed to update Category: %w", result.Error)
	}

	// Check if any rows were affected
	if result.RowsAffected == 0 {
		log.Printf("No records updated for ProjectID: %d, ProjectUserID: %d", projectID, projectUserID)
		return fmt.Errorf("no matching records found to update")
	}

	log.Printf("Successfully updated Category to '%s' for ProjectID: %d, ProjectUserID: %d", userName, projectID, projectUserID)
	return nil
}
