package repositories

import (
	"AmanahPro/services/project-management/internal/domain/models"
	"AmanahPro/services/project-management/internal/domain/repositories"

	"gorm.io/gorm"
)

type projectUserRepositoryImpl struct {
	db *gorm.DB
}

func NewProjectUserRepository(db *gorm.DB) repositories.ProjectUserRepository {
	return &projectUserRepositoryImpl{db: db}
}

func (r *projectUserRepositoryImpl) Create(projectUser *models.ProjectUser) error {
	return r.db.Create(projectUser).Error
}

func (r *projectUserRepositoryImpl) Update(projectUser *models.ProjectUser) error {
	return r.db.Save(projectUser).Error
}

func (r *projectUserRepositoryImpl) Delete(id int) error {
	return r.db.Delete(&models.ProjectUser{}, id).Error
}

func (r *projectUserRepositoryImpl) FindByID(id int) (*models.ProjectUser, error) {
	var projectUser models.ProjectUser
	err := r.db.First(&projectUser, id).Error
	return &projectUser, err
}

func (r *projectUserRepositoryImpl) FindAllByProjectID(projectID int) ([]models.ProjectUser, error) {
	var projectUsers []models.ProjectUser
	err := r.db.Where("ProjectID = ?", projectID).Find(&projectUsers).Error
	return projectUsers, err
}
