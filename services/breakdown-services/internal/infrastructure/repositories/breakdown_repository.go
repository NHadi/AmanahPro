package repositories

import (
	"AmanahPro/services/breakdown-services/common/helpers"
	"AmanahPro/services/breakdown-services/internal/domain/models"
	"AmanahPro/services/breakdown-services/internal/domain/repositories"
	"context"
	"fmt"
	"log"
	"sort"

	"github.com/elastic/go-elasticsearch/v8"
	"gorm.io/gorm"
)

type breakdownRepositoryImpl struct {
	db       *gorm.DB
	esClient *elasticsearch.Client
	esIndex  string
}

func NewBreakdownRepository(db *gorm.DB, esClient *elasticsearch.Client, esIndex string) repositories.BreakdownRepository {
	return &breakdownRepositoryImpl{
		db:       db,
		esClient: esClient,
		esIndex:  esIndex,
	}
}

// FilterBreakdowns retrieves breakdowns from Elasticsearch by organization ID with optional filters
func (r *breakdownRepositoryImpl) FilterBreakdowns(organizationID int, breakdownID *int, projectID *int) ([]models.Breakdown, error) {
	log.Printf("Filtering breakdowns from Elasticsearch by OrganizationID: %d, BreakdownID: %v, ProjectID: %v", organizationID, breakdownID, projectID)

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

	// Add optional filters
	if breakdownID != nil {
		query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = append(
			query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"].([]map[string]interface{}),
			map[string]interface{}{
				"term": map[string]interface{}{
					"BreakdownId": *breakdownID,
				},
			},
		)
	}
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
	breakdowns, err := helpers.ParseResponse[models.Breakdown](res)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Elasticsearch response: %w", err)
	}

	// Sort Sections and Items within each Breakdown
	for bIdx := range breakdowns {
		sort.Slice(breakdowns[bIdx].Sections, func(i, j int) bool {
			return breakdowns[bIdx].Sections[i].Sort < breakdowns[bIdx].Sections[j].Sort
		})

		for sIdx := range breakdowns[bIdx].Sections {
			sort.Slice(breakdowns[bIdx].Sections[sIdx].Items, func(i, j int) bool {
				return breakdowns[bIdx].Sections[sIdx].Items[i].Sort < breakdowns[bIdx].Sections[sIdx].Items[j].Sort
			})
		}
	}

	log.Printf("Found %d breakdowns for OrganizationID: %d", len(breakdowns), organizationID)
	return breakdowns, nil
}

// GetByID fetches a Breakdown by ID (SQL)
func (r *breakdownRepositoryImpl) GetByID(breakdownID int) (*models.Breakdown, error) {
	var breakdown models.Breakdown
	if err := r.db.Preload("Sections.Items").First(&breakdown, "BreakdownId = ?", breakdownID).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch full breakdown: %w", err)
	}
	return &breakdown, nil
}

// Create inserts a new Breakdown into the database (SQL)
func (r *breakdownRepositoryImpl) Create(breakdown *models.Breakdown) error {
	if err := r.db.Create(breakdown).Error; err != nil {
		return fmt.Errorf("failed to create breakdown: %w", err)
	}
	return nil
}

// Update modifies an existing Breakdown (SQL)
func (r *breakdownRepositoryImpl) Update(breakdown *models.Breakdown) error {
	if err := r.db.Save(breakdown).Error; err != nil {
		return fmt.Errorf("failed to update breakdown: %w", err)
	}
	return nil
}

// Delete removes a Breakdown by ID (SQL)
func (r *breakdownRepositoryImpl) Delete(breakdownID int) error {
	if err := r.db.Delete(&models.Breakdown{}, "BreakdownId = ?", breakdownID).Error; err != nil {
		return fmt.Errorf("failed to delete breakdown: %w", err)
	}
	return nil
}
