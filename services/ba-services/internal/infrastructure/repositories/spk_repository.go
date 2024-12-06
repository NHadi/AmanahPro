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

// NewSPKRepository creates a new instance of SPKRepository
func NewSPKRepository(db *gorm.DB, esClient *elasticsearch.Client, esIndex string) repositories.SPKRepository {
	return &baRepositoryImpl{
		db:       db,
		esClient: esClient,
		esIndex:  esIndex,
	}
}

// Filter retrieves SPKs from Elasticsearch by organization ID with optional filters
func (r *baRepositoryImpl) Filter(organizationID int, baID *int, projectID *int) ([]models.SPK, error) {
	log.Printf("Filtering SPKs from Elasticsearch by OrganizationID: %d, SpkID: %v, ProjectID: %v", organizationID, baID, projectID)

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

	// Add optional filters for SpkID
	if baID != nil {
		query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = append(
			query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"].([]map[string]interface{}),
			map[string]interface{}{
				"term": map[string]interface{}{
					"SpkId": *baID,
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
	bas, err := helpers.ParseResponse[models.SPK](res)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Elasticsearch response: %w", err)
	}

	log.Printf("Found %d SPKs for OrganizationID: %d", len(bas), organizationID)
	return bas, nil
}

// Create inserts a new SPK record into the database
func (r *baRepositoryImpl) Create(ba *models.SPK) error {
	log.Printf("Creating SPK: %+v", ba)

	if err := r.db.Create(ba).Error; err != nil {
		log.Printf("Failed to create SPK: %v", err)
		return fmt.Errorf("failed to create SPK: %w", err)
	}

	log.Printf("Successfully created SPK: %+v", ba)
	return nil
}

// Update modifies an existing SPK record in the database
func (r *baRepositoryImpl) Update(ba *models.SPK) error {
	log.Printf("Updating SPK ID: %d", ba.SpkId)

	if err := r.db.Save(ba).Error; err != nil {
		log.Printf("Failed to update SPK ID %d: %v", ba.SpkId, err)
		return fmt.Errorf("failed to update SPK: %w", err)
	}

	log.Printf("Successfully updated SPK ID: %d", ba.SpkId)
	return nil
}

// Delete removes a SPK record from the database
func (r *baRepositoryImpl) Delete(baId int) error {
	log.Printf("Deleting SPK ID: %d", baId)

	if err := r.db.Delete(&models.SPK{}, baId).Error; err != nil {
		log.Printf("Failed to delete SPK ID %d: %v", baId, err)
		return fmt.Errorf("failed to delete SPK: %w", err)
	}

	log.Printf("Successfully deleted SPK ID: %d", baId)
	return nil
}

// GetByID retrieves a SPK record by its ID
func (r *baRepositoryImpl) GetByID(baId int) (*models.SPK, error) {
	log.Printf("Retrieving SPK by ID: %d", baId)

	var ba models.SPK
	if err := r.db.Preload("Sections.Details").First(&ba, baId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("SPK ID %d not found", baId)
			return nil, nil
		}
		log.Printf("Failed to retrieve SPK ID %d: %v", baId, err)
		return nil, fmt.Errorf("failed to retrieve SPK: %w", err)
	}

	log.Printf("Successfully retrieved SPK: %+v", ba)
	return &ba, nil
}
