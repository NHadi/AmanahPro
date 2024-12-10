package repositories

import (
	"AmanahPro/services/ba-services/internal/domain/models"
	"AmanahPro/services/ba-services/internal/domain/repositories"
	"context"
	"fmt"
	"log"

	"github.com/NHadi/AmanahPro-common/helpers"
	"github.com/elastic/go-elasticsearch/v8"
	"gorm.io/gorm"
)

type baRepositoryImpl struct {
	db       *gorm.DB
	esClient *elasticsearch.Client
	esIndex  string
}

// NewBARepository creates a new instance of BARepository
func NewBARepository(db *gorm.DB, esClient *elasticsearch.Client, esIndex string) repositories.BARepository {
	return &baRepositoryImpl{
		db:       db,
		esClient: esClient,
		esIndex:  esIndex}
}

func (r *baRepositoryImpl) Create(ba *models.BA) error {
	log.Printf("Creating BA: %+v", ba)
	if err := r.db.Create(ba).Error; err != nil {
		log.Printf("Failed to create BA: %v", err)
		return fmt.Errorf("failed to create BA: %w", err)
	}
	return nil
}

func (r *baRepositoryImpl) Update(ba *models.BA) error {
	log.Printf("Updating BA ID: %d", ba.BAId)
	if err := r.db.Save(ba).Error; err != nil {
		log.Printf("Failed to update BA ID %d: %v", ba.BAId, err)
		return fmt.Errorf("failed to update BA: %w", err)
	}
	return nil
}

func (r *baRepositoryImpl) Delete(baId int) error {
	log.Printf("Deleting BA ID: %d", baId)
	if err := r.db.Delete(&models.BA{}, baId).Error; err != nil {
		log.Printf("Failed to delete BA ID %d: %v", baId, err)
		return fmt.Errorf("failed to delete BA: %w", err)
	}
	return nil
}

func (r *baRepositoryImpl) GetByID(baId int, loadSections bool) (*models.BA, error) {
	log.Printf("Retrieving BA by ID: %d, loadSections: %v", baId, loadSections)
	var ba models.BA
	query := r.db
	if loadSections {
		query = query.Preload("Sections.Details.Progress")
	}
	if err := query.First(&ba, baId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("BA ID %d not found", baId)
			return nil, nil
		}
		log.Printf("Failed to retrieve BA ID %d: %v", baId, err)
		return nil, fmt.Errorf("failed to retrieve BA: %w", err)
	}
	return &ba, nil
}

// Filter retrieves BAs from Elasticsearch by organization ID with optional filters
func (r *baRepositoryImpl) Filter(organizationID int, baID *int, projectID *int) ([]models.BA, error) {
	log.Printf("Filtering BAs from Elasticsearch by OrganizationID: %d, BAID: %v, ProjectID: %v", organizationID, baID, projectID)

	// Build the Elasticsearch query
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{
					{
						"term": map[string]interface{}{
							"OrganizationId": organizationID,
						},
					},
				},
			},
		},
	}

	// Add optional filters for BAID
	if baID != nil {
		query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = append(
			query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"].([]map[string]interface{}),
			map[string]interface{}{
				"term": map[string]interface{}{
					"BAId": *baID,
				},
			},
		)
	}

	// Add optional filters for ProjectID
	if projectID != nil {
		query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = append(
			query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"].([]map[string]interface{}),
			map[string]interface{}{
				"term": map[string]interface{}{
					"ProjectId": *projectID,
				},
			},
		)
	}

	// Convert the query map to a JSON reader
	queryReader, err := helpers.MapToReader(query)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare query: %w", err)
	}

	// Execute the Elasticsearch query
	res, err := r.esClient.Search(
		r.esClient.Search.WithIndex(r.esIndex),
		r.esClient.Search.WithBody(queryReader),
		r.esClient.Search.WithContext(context.Background()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to query Elasticsearch: %w", err)
	}
	defer res.Body.Close()

	// Parse the response
	bas, err := ParseResponse[models.BA](res)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Elasticsearch response: %w", err)
	}

	log.Printf("Found %d BAs for OrganizationID: %d", len(bas), organizationID)
	return bas, nil
}
