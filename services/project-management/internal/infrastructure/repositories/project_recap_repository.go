package repositories

import (
	"AmanahPro/services/project-management/internal/domain/models"
	"AmanahPro/services/project-management/internal/domain/repositories"

	"gorm.io/gorm"
)

type projectRecapRepositoryImpl struct {
	db *gorm.DB
}

func NewProjectRecapRepository(db *gorm.DB) repositories.ProjectRecapRepository {
	return &projectRecapRepositoryImpl{db: db}
}

func (r *projectRecapRepositoryImpl) Create(projectRecap *models.ProjectRecap) error {
	return r.db.Create(projectRecap).Error
}

func (r *projectRecapRepositoryImpl) Update(projectRecap *models.ProjectRecap) error {
	return r.db.Save(projectRecap).Error
}

func (r *projectRecapRepositoryImpl) Delete(id int) error {
	return r.db.Delete(&models.ProjectRecap{}, id).Error
}

func (r *projectRecapRepositoryImpl) FindByID(id int) (*models.ProjectRecap, error) {
	var projectRecap models.ProjectRecap
	err := r.db.First(&projectRecap, id).Error
	return &projectRecap, err
}

func (r *projectRecapRepositoryImpl) FindAllByOrganizationID(organizationID int) ([]models.ProjectRecap, error) {
	var projectRecaps []models.ProjectRecap
	err := r.db.Where("OrganizationID = ?", organizationID).Find(&projectRecaps).Error
	return projectRecaps, err
}
