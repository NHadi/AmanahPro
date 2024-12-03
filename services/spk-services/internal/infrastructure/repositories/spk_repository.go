package repositories

import (
	"AmanahPro/services/spk-services/common/helpers"
	"AmanahPro/services/spk-services/internal/domain/models"
	"AmanahPro/services/spk-services/internal/domain/repositories"
	"context"
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
	"gorm.io/gorm"
)

type spkRepositoryImpl struct {
	db       *gorm.DB
	esClient *elasticsearch.Client
	esIndex  string
}

// NewSPKRepository creates a new instance of SPKRepository
func NewSPKRepository(db *gorm.DB, esClient *elasticsearch.Client, esIndex string) repositories.SPKRepository {
	return &spkRepositoryImpl{
		db:       db,
		esClient: esClient,
		esIndex:  esIndex,
	}
}

// Filter retrieves SPKs from Elasticsearch by organization ID with optional filters
func (r *spkRepositoryImpl) Filter(organizationID int, spkID *int, projectID *int) ([]models.SPK, error) {
	log.Printf("Filtering SPKs from Elasticsearch by OrganizationID: %d, SpkID: %v, ProjectID: %v", organizationID, spkID, projectID)

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
	if spkID != nil {
		query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = append(
			query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"].([]map[string]interface{}),
			map[string]interface{}{
				"term": map[string]interface{}{
					"SpkId": *spkID,
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
	spks, err := helpers.ParseResponse[models.SPK](res)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Elasticsearch response: %w", err)
	}

	log.Printf("Found %d SPKs for OrganizationID: %d", len(spks), organizationID)
	return spks, nil
}

// Create inserts a new SPK record into the database
func (r *spkRepositoryImpl) Create(spk *models.SPK) error {
	log.Printf("Creating SPK: %+v", spk)

	if err := r.db.Create(spk).Error; err != nil {
		log.Printf("Failed to create SPK: %v", err)
		return fmt.Errorf("failed to create SPK: %w", err)
	}

	log.Printf("Successfully created SPK: %+v", spk)
	return nil
}

// Update modifies an existing SPK record in the database
func (r *spkRepositoryImpl) Update(spk *models.SPK) error {
	log.Printf("Updating SPK ID: %d", spk.SpkId)

	if err := r.db.Save(spk).Error; err != nil {
		log.Printf("Failed to update SPK ID %d: %v", spk.SpkId, err)
		return fmt.Errorf("failed to update SPK: %w", err)
	}

	log.Printf("Successfully updated SPK ID: %d", spk.SpkId)
	return nil
}

// Delete removes a SPK record from the database
func (r *spkRepositoryImpl) Delete(spkId int) error {
	log.Printf("Deleting SPK ID: %d", spkId)

	if err := r.db.Delete(&models.SPK{}, spkId).Error; err != nil {
		log.Printf("Failed to delete SPK ID %d: %v", spkId, err)
		return fmt.Errorf("failed to delete SPK: %w", err)
	}

	log.Printf("Successfully deleted SPK ID: %d", spkId)
	return nil
}

// GetByID retrieves a SPK record by its ID
func (r *spkRepositoryImpl) GetByID(spkId int) (*models.SPK, error) {
	log.Printf("Retrieving SPK by ID: %d", spkId)

	var spk models.SPK
	if err := r.db.Preload("Sections.Details").First(&spk, spkId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("SPK ID %d not found", spkId)
			return nil, nil
		}
		log.Printf("Failed to retrieve SPK ID %d: %v", spkId, err)
		return nil, fmt.Errorf("failed to retrieve SPK: %w", err)
	}

	log.Printf("Successfully retrieved SPK: %+v", spk)
	return &spk, nil
}
