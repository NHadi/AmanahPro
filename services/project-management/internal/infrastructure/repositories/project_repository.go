package repositories

import (
	"AmanahPro/services/project-management/internal/domain/models"
	"AmanahPro/services/project-management/internal/domain/repositories"

	"gorm.io/gorm"
)

type projectRepositoryImpl struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) repositories.ProjectRepository {
	return &projectRepositoryImpl{db: db}
}

func (r *projectRepositoryImpl) Create(project *models.Project) error {
	return r.db.Create(project).Error
}

func (r *projectRepositoryImpl) Update(project *models.Project) error {
	return r.db.Save(project).Error
}

func (r *projectRepositoryImpl) Delete(id int) error {
	return r.db.Delete(&models.Project{}, id).Error
}

func (r *projectRepositoryImpl) FindByID(id int) (*models.Project, error) {
	var project models.Project
	err := r.db.First(&project, id).Error
	return &project, err
}

func (r *projectRepositoryImpl) FindAllByOrganizationID(organizationID int) ([]models.Project, error) {
	var projects []models.Project
	err := r.db.Where("OrganizationID = ?", organizationID).Find(&projects).Error
	return projects, err
}
