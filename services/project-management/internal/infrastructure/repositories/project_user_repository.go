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

type projectUserRepositoryImpl struct {
	db       *gorm.DB
	esClient *elasticsearch.Client
	esIndex  string
}

func NewProjectUserRepository(db *gorm.DB, esClient *elasticsearch.Client, esIndex string) repositories.ProjectUserRepository {
	return &projectUserRepositoryImpl{
		db:       db,
		esClient: esClient,
		esIndex:  esIndex,
	}
}

// Create inserts a new ProjectUser record in the database
func (r *projectUserRepositoryImpl) Create(user *models.ProjectUser) error {
	log.Printf("Repository: Creating project user for ProjectID: %d, UserID: %d", user.ProjectID, user.UserID)

	if err := r.db.Create(user).Error; err != nil {
		log.Printf("Repository: Error creating project user: %v", err)
		return fmt.Errorf("failed to create project user: %w", err)
	}

	log.Printf("Repository: Successfully created project user: %+v", user)
	return nil
}

// Update updates an existing ProjectUser record in the database
func (r *projectUserRepositoryImpl) Update(user *models.ProjectUser) error {
	log.Printf("Repository: Updating project user ID: %d", user.ID)

	if err := r.db.Save(user).Error; err != nil {
		log.Printf("Repository: Error updating project user: %v", err)
		return fmt.Errorf("failed to update project user: %w", err)
	}

	log.Printf("Repository: Successfully updated project user ID: %d", user.ID)
	return nil
}

// Delete removes a ProjectUser record from the database
func (r *projectUserRepositoryImpl) Delete(id int) error {
	log.Printf("Repository: Deleting project user ID: %d", id)

	if err := r.db.Delete(&models.ProjectUser{}, id).Error; err != nil {
		log.Printf("Repository: Error deleting project user ID: %d, %v", id, err)
		return fmt.Errorf("failed to delete project user: %w", err)
	}

	log.Printf("Repository: Successfully deleted project user ID: %d", id)
	return nil
}

// FindByProjectID retrieves all ProjectUser records for a given ProjectID from Elasticsearch
func (r *projectUserRepositoryImpl) FindByProjectID(projectID int, organizationID *int) ([]dto.ProjectUserDTO, error) {
	log.Printf("Repository: Searching project users for ProjectID: %d", projectID)

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"term": map[string]interface{}{
				"ProjectID": projectID,
			},
		},
	}

	// Add the OrganizationID filter if provided
	if organizationID != nil {
		orgFilter := map[string]interface{}{
			"term": map[string]interface{}{
				"OrganizationID": *organizationID,
			},
		}
		query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = append(
			query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"].([]interface{}),
			orgFilter,
		)
	}

	queryReader, err := helpers.MapToReader(query)
	if err != nil {
		log.Printf("Repository: Error preparing query for project users: %v", err)
		return nil, fmt.Errorf("error preparing query: %w", err)
	}

	res, err := r.esClient.Search(
		r.esClient.Search.WithIndex(r.esIndex),
		r.esClient.Search.WithBody(queryReader),
		r.esClient.Search.WithContext(context.Background()),
	)
	if err != nil {
		log.Printf("Repository: Error searching project users in Elasticsearch: %v", err)
		return nil, fmt.Errorf("failed to search project users in Elasticsearch: %w", err)
	}
	defer res.Body.Close()

	users, err := helpers.ParseResponse[dto.ProjectUserDTO](res)
	if err != nil {
		log.Printf("Repository: Error parsing Elasticsearch response for project users: %v", err)
		return nil, fmt.Errorf("failed to parse project users from Elasticsearch response: %w", err)
	}

	if len(users) == 0 {
		return []dto.ProjectUserDTO{}, nil
	}

	log.Printf("Repository: Found %d project users for ProjectID: %d", len(users), projectID)
	return users, nil
}

func (r *projectUserRepositoryImpl) FindByUserAndProject(userID, projectID int, organizationID *int) (*dto.ProjectUserDTO, error) {
	// Build the base query
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []interface{}{
					map[string]interface{}{
						"term": map[string]interface{}{
							"UserID": userID,
						},
					},
					map[string]interface{}{
						"term": map[string]interface{}{
							"ProjectID": projectID,
						},
					},
				},
			},
		},
	}

	// Add the OrganizationID filter if provided
	if organizationID != nil {
		orgFilter := map[string]interface{}{
			"term": map[string]interface{}{
				"OrganizationID": *organizationID,
			},
		}
		query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = append(
			query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"].([]interface{}),
			orgFilter,
		)
	}

	// Convert the query to JSON reader
	queryReader, err := helpers.MapToReader(query)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare query: %w", err)
	}

	// Execute the Elasticsearch search
	res, err := r.esClient.Search(
		r.esClient.Search.WithIndex(r.esIndex),
		r.esClient.Search.WithBody(queryReader),
		r.esClient.Search.WithContext(context.Background()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to execute search query: %w", err)
	}
	defer res.Body.Close()

	// Parse the Elasticsearch response
	projectUsers, err := helpers.ParseResponse[dto.ProjectUserDTO](res)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Elasticsearch response: %w", err)
	}

	// Return the first match if found, otherwise return nil
	if len(projectUsers) == 0 {
		log.Printf("user not found in project (UserID: %d, ProjectID: %d)", userID, projectID)
		return nil, nil
	}

	return &projectUsers[0], nil
}
