package repositories

import (
	"AmanahPro/services/sph-services/internal/domain/models"
	"AmanahPro/services/sph-services/internal/domain/repositories"
	"context"
	"fmt"
	"log"

	"github.com/NHadi/AmanahPro-common/helpers"

	"github.com/elastic/go-elasticsearch/v8"
	"gorm.io/gorm"
)

type sphRepositoryImpl struct {
	db       *gorm.DB
	esClient *elasticsearch.Client
	esIndex  string
}

// NewSphRepository creates a new instance of SphRepository
func NewSphRepository(db *gorm.DB, esClient *elasticsearch.Client, esIndex string) repositories.SphRepository {
	return &sphRepositoryImpl{
		db:       db,
		esClient: esClient,
		esIndex:  esIndex,
	}
}

// Begin starts a transaction
func (r *sphRepositoryImpl) Begin() (repositories.SphRepository, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &sphRepositoryImpl{db: tx}, nil
}

// Commit commits the transaction
func (r *sphRepositoryImpl) Commit() error {
	return r.db.Commit().Error
}

// Rollback rolls back the transaction
func (r *sphRepositoryImpl) Rollback() error {
	return r.db.Rollback().Error
}

// FilterSPHs retrieves SPHs from Elasticsearch by organization ID with optional filters
func (r *sphRepositoryImpl) Filter(organizationID int, sphID *int, projectID *int) ([]models.Sph, error) {
	log.Printf("Filtering SPHs from Elasticsearch by OrganizationID: %d, SphID: %v, ProjectID: %v", organizationID, sphID, projectID)

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

	// Add optional filters for SphId
	if sphID != nil {
		query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = append(
			query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"].([]map[string]interface{}),
			map[string]interface{}{
				"term": map[string]interface{}{
					"SphId": *sphID,
				},
			},
		)
	}

	// Add optional filters for ProjectId
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

	// Convert the query map to a JSON reader using the helper
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

	// Parse the response using the helper
	sphs, err := helpers.ParseResponse[models.Sph](res)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Elasticsearch response: %w", err)
	}

	log.Printf("Found %d SPHs for OrganizationID: %d", len(sphs), organizationID)
	return sphs, nil
}

// Create inserts a new SPH record into the database
func (r *sphRepositoryImpl) Create(sph *models.Sph) error {
	log.Printf("Creating SPH: %+v", sph)

	if err := r.db.Create(sph).Error; err != nil {
		log.Printf("Failed to create SPH: %v", err)
		return fmt.Errorf("failed to create SPH: %w", err)
	}

	log.Printf("Successfully created SPH: %+v", sph)
	return nil
}

// Update modifies an existing SPH record in the database
func (r *sphRepositoryImpl) Update(sph *models.Sph) error {
	log.Printf("Updating SPH ID: %d", sph.SphId)

	if err := r.db.Save(sph).Error; err != nil {
		log.Printf("Failed to update SPH ID %d: %v", sph.SphId, err)
		return fmt.Errorf("failed to update SPH: %w", err)
	}

	log.Printf("Successfully updated SPH ID: %d", sph.SphId)
	return nil
}

// Delete removes a SPH record from the database
func (r *sphRepositoryImpl) Delete(sphID int) error {
	log.Printf("Deleting SPH ID: %d", sphID)

	if err := r.db.Delete(&models.Sph{}, sphID).Error; err != nil {
		log.Printf("Failed to delete SPH ID %d: %v", sphID, err)
		return fmt.Errorf("failed to delete SPH: %w", err)
	}

	log.Printf("Successfully deleted SPH ID: %d", sphID)
	return nil
}

// GetByID retrieves a SPH record by its ID
func (r *sphRepositoryImpl) GetByID(sphID int) (*models.Sph, error) {
	log.Printf("Retrieving SPH by ID: %d", sphID)

	var sph models.Sph
	if err := r.db.Preload("Sections.Details").First(&sph, sphID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("SPH ID %d not found", sphID)
			return nil, nil
		}
		log.Printf("Failed to retrieve SPH ID %d: %v", sphID, err)
		return nil, fmt.Errorf("failed to retrieve SPH: %w", err)
	}

	log.Printf("Successfully retrieved SPH: %+v", sph)
	return &sph, nil
}
