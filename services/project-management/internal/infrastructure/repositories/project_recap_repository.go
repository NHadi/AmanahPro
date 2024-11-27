package repositories

import (
	"AmanahPro/services/project-management/common/helpers"
	"AmanahPro/services/project-management/internal/domain/models"
	"AmanahPro/services/project-management/internal/domain/repositories"
	"AmanahPro/services/project-management/internal/dto"
	"context"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
	"gorm.io/gorm"
)

type projectRecapRepositoryImpl struct {
	db       *gorm.DB
	esClient *elasticsearch.Client
	esIndex  string
}

func NewProjectRecapRepository(db *gorm.DB, esClient *elasticsearch.Client, esIndex string) repositories.ProjectRecapRepository {
	return &projectRecapRepositoryImpl{
		db:       db,
		esClient: esClient,
		esIndex:  esIndex,
	}
}

func (r *projectRecapRepositoryImpl) Create(recap *models.ProjectRecap) error {
	if err := r.db.Create(recap).Error; err != nil {
		return fmt.Errorf("failed to create project recap: %w", err)
	}
	return nil
}

func (r *projectRecapRepositoryImpl) Update(recap *models.ProjectRecap) error {
	if err := r.db.Save(recap).Error; err != nil {
		return fmt.Errorf("failed to update project recap: %w", err)
	}
	return nil
}

func (r *projectRecapRepositoryImpl) Delete(id int) error {
	if err := r.db.Delete(&models.ProjectRecap{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete project recap: %w", err)
	}
	return nil
}

func (r *projectRecapRepositoryImpl) FindByProjectID(projectID int) (*dto.ProjectRecapDTO, error) {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"term": map[string]interface{}{
				"ProjectID": projectID,
			},
		},
	}

	// Execute the search query
	queryReader, err := helpers.MapToReader(query)
	if err != nil {
		return nil, fmt.Errorf("error preparing query: %v", err)
	}

	res, err := r.esClient.Search(
		r.esClient.Search.WithIndex(r.esIndex),
		r.esClient.Search.WithBody(queryReader),
		r.esClient.Search.WithContext(context.Background()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to search project recap in Elasticsearch: %w", err)
	}
	defer res.Body.Close()

	recaps, err := helpers.ParseResponse[dto.ProjectRecapDTO](res)
	if err != nil || len(recaps) == 0 {
		return nil, fmt.Errorf("no recap found for project ID: %d", projectID)
	}

	return &recaps[0], nil
}
