package repositories

import (
	"AmanahPro/services/project-management/common/helpers"
	"AmanahPro/services/project-management/internal/domain/models"
	"AmanahPro/services/project-management/internal/domain/repositories"
	"AmanahPro/services/project-management/internal/dto"
	"context"
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
	"gorm.io/gorm"
)

type projectRepositoryImpl struct {
	db       *gorm.DB
	esClient *elasticsearch.Client
	esIndex  string
}

func NewProjectRepository(db *gorm.DB, esClient *elasticsearch.Client, esIndex string) repositories.ProjectRepository {
	return &projectRepositoryImpl{
		db:       db,
		esClient: esClient,
		esIndex:  esIndex,
	}
}

// SQL Operations (Create, Update, Delete)

// Create adds a new project to the database
func (r *projectRepositoryImpl) Create(project *models.Project) error {
	log.Printf("Attempting to create a new project: %+v", project)

	err := r.db.Create(project).Error
	if err != nil {
		log.Printf("Error creating project: %+v, Error: %v", project, err)
		return fmt.Errorf("failed to create project: %w", err)
	}

	log.Printf("Successfully created project with ID: %d", project.ProjectID)
	return nil
}

// Update modifies an existing project in the database
func (r *projectRepositoryImpl) Update(project *models.Project) error {
	log.Printf("Attempting to update project with ID: %d", project.ProjectID)

	err := r.db.Save(project).Error
	if err != nil {
		log.Printf("Error updating project with ID: %d, Error: %v", project.ProjectID, err)
		return fmt.Errorf("failed to update project with ID %d: %w", project.ProjectID, err)
	}

	log.Printf("Successfully updated project with ID: %d", project.ProjectID)
	return nil
}

// Delete removes a project from the database
func (r *projectRepositoryImpl) Delete(id int) error {
	log.Printf("Attempting to delete project with ID: %d", id)

	err := r.db.Delete(&models.Project{}, id).Error
	if err != nil {
		log.Printf("Error deleting project with ID: %d, Error: %v", id, err)
		return fmt.Errorf("failed to delete project with ID %d: %w", id, err)
	}

	log.Printf("Successfully deleted project with ID: %d", id)
	return nil
}

// Elasticsearch Read Operations

func (r *projectRepositoryImpl) SearchProjectsByOrganization(organizationID int, query string) ([]dto.ProjectDTO, error) {
	// Build Elasticsearch query
	esQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []interface{}{
					map[string]interface{}{
						"term": map[string]interface{}{
							"OrganizationID": organizationID,
						},
					},
					map[string]interface{}{
						"multi_match": map[string]interface{}{
							"query":  query,
							"fields": []string{"ProjectName", "Location"},
						},
					},
				},
			},
		},
	}

	// Execute the search query
	queryReader, err := helpers.MapToReader(esQuery)
	if err != nil {
		return nil, fmt.Errorf("error preparing query: %v", err)
	}

	res, err := r.esClient.Search(
		r.esClient.Search.WithIndex(r.esIndex),
		r.esClient.Search.WithBody(queryReader),
		r.esClient.Search.WithContext(context.Background()),
	)
	if err != nil {
		return nil, fmt.Errorf("error executing search query: %v", err)
	}
	defer res.Body.Close()

	// Parse Elasticsearch response into DTOs
	projects, err := helpers.ParseResponse[dto.ProjectDTO](res)
	if err != nil {
		return nil, fmt.Errorf("error parsing Elasticsearch response: %v", err)
	}

	// Return an empty slice if no projects are found
	if len(projects) == 0 {
		return []dto.ProjectDTO{}, nil
	}

	return projects, nil
}

// helper

func ToProjectDTO(project *models.Project, recap *models.ProjectRecap, keyUsers []models.ProjectUser) *dto.ProjectDTO {
	var recapDTO *dto.ProjectRecapDTO
	if recap != nil {
		recapDTO = &dto.ProjectRecapDTO{
			TotalOpname:      recap.TotalOpname,
			TotalPengeluaran: recap.TotalPengeluaran,
			Margin:           recap.Margin,
			MarginPercentage: recap.MarginPercentage,
		}
	}

	keyUsersDTO := make([]dto.ProjectUserDTO, len(keyUsers))
	for i, user := range keyUsers {
		keyUsersDTO[i] = dto.ProjectUserDTO{
			UserID: user.UserID,
			Role:   *user.Role,
		}
	}

	return &dto.ProjectDTO{
		ProjectID:      project.ProjectID,
		ProjectName:    project.ProjectName,
		Location:       *project.Location,
		Status:         *project.Status,
		OrganizationID: project.OrganizationID,
		Recap:          recapDTO,
		KeyUsers:       keyUsersDTO,
		CreatedAt:      *project.CreatedAt,
		UpdatedAt:      *project.UpdatedAt,
	}
}
