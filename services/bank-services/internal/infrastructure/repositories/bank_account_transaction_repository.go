package repositories

import (
	"AmanahPro/services/bank-services/internal/application/dto"
	"AmanahPro/services/bank-services/internal/domain/models"
	"AmanahPro/services/bank-services/internal/domain/repositories"
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"gorm.io/gorm"
)

type bankAccountTransactionRepository struct {
	db       *gorm.DB
	esClient *elasticsearch.Client
	esIndex  string
}

// Ensure the struct implements the interface
var _ repositories.BankAccountTransactionRepository = &bankAccountTransactionRepository{}

func NewBankAccountTransactionRepository(db *gorm.DB, esClient *elasticsearch.Client, esIndex string) repositories.BankAccountTransactionRepository {
	return &bankAccountTransactionRepository{
		db:       db,
		esClient: esClient,
		esIndex:  esIndex,
	}
}

// InsertWithRollback saves the batch and transactions in the SQL database
func (r *bankAccountTransactionRepository) InsertWithRollback(batch *models.UploadBatch, transactions []models.BankAccountTransactions) error {
	// Begin a transaction
	tx := r.db.Begin()

	// Save the batch
	if err := tx.Create(batch).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to save batch: %v", err)
	}

	// Assign the generated BatchID to each transaction
	for i := range transactions {
		transactions[i].BatchID = batch.BatchID
	}

	// Save transactions
	for _, transaction := range transactions {
		if err := tx.Create(&transaction).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to save transaction: %v", err)
		}
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}

func (r *bankAccountTransactionRepository) GetTransactionsByBankAndPeriod(organizationID, bankID uint, year *int) ([]dto.BankAccountTransactionDTO, error) {
	// Build the base query for filtering by AccountID and OrganizationID
	query := map[string]interface{}{
		"size": 10000, // Adjust this size as needed for your use case
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{
					{
						"term": map[string]interface{}{
							"AccountID": bankID,
						},
					},
					{
						"term": map[string]interface{}{
							"OrganizationId": organizationID,
						},
					},
				},
			},
		},
		"sort": []map[string]interface{}{ // Add sorting parameters
			{
				"Tanggal": map[string]interface{}{
					"order": "asc", // Sort by Tanggal in ascending order
				},
			},
			{
				"ID": map[string]interface{}{
					"order": "asc", // Sort by ID in ascending order
				},
			},
		},
	}

	// If year is provided, add a range filter for the Tanggal field
	if year != nil {
		periodeStart := time.Date(*year, time.January, 1, 0, 0, 0, 0, time.UTC).Format(time.RFC3339)
		periodeEnd := time.Date(*year, time.December, 31, 23, 59, 59, 0, time.UTC).Format(time.RFC3339)

		rangeFilter := map[string]interface{}{
			"range": map[string]interface{}{
				"Tanggal": map[string]interface{}{
					"gte": periodeStart,
					"lte": periodeEnd,
				},
			},
		}

		// Append the range filter to the must clause
		query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = append(
			query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"].([]map[string]interface{}),
			rangeFilter,
		)
	}

	// Execute the search query
	res, err := r.esClient.Search(
		r.esClient.Search.WithIndex(r.esIndex),
		r.esClient.Search.WithBody(mapToReader(query)),
		r.esClient.Search.WithContext(context.Background()),
	)
	if err != nil {
		return nil, fmt.Errorf("error executing search query: %v", err)
	}
	defer res.Body.Close()

	// Parse the Elasticsearch response
	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error parsing Elasticsearch response: %v", err)
	}

	// Extract hits
	hitsData, ok := result["hits"].(map[string]interface{})
	if !ok {
		// Return an empty list instead of an error
		return []dto.BankAccountTransactionDTO{}, nil
	}

	hits, ok := hitsData["hits"].([]interface{})
	if !ok || len(hits) == 0 {
		// Return an empty list if no transactions are found
		return []dto.BankAccountTransactionDTO{}, nil
	}

	var transactions []dto.BankAccountTransactionDTO
	for _, hit := range hits {
		hitMap, ok := hit.(map[string]interface{})
		if !ok {
			continue // Skip invalid hits
		}

		source, ok := hitMap["_source"].(map[string]interface{})
		if !ok {
			continue // Skip hits without a valid _source
		}

		// Map Elasticsearch fields to DTO, with safe type assertions
		transaction := dto.BankAccountTransactionDTO{}

		if id, ok := source["ID"].(float64); ok {
			transaction.ID = uint(id)
		}
		if accountID, ok := source["AccountID"].(float64); ok {
			transaction.AccountID = uint(accountID)
		}
		if batchID, ok := source["BatchID"].(float64); ok {
			transaction.BatchID = uint(batchID)
		}
		if orgID, ok := source["OrganizationId"].(float64); ok {
			transaction.OrganizationId = uint(orgID)
		}
		if tanggal, ok := source["Tanggal"].(string); ok {
			transaction.Tanggal = tanggal
		}
		if keterangan, ok := source["Keterangan"].(string); ok {
			transaction.Keterangan = keterangan
		}
		if cabang, ok := source["Cabang"].(string); ok {
			transaction.Cabang = cabang
		}
		if credit, ok := source["Credit"].(float64); ok {
			transaction.Credit = credit
		}
		if debit, ok := source["Debit"].(float64); ok {
			transaction.Debit = debit
		}
		if saldo, ok := source["Saldo"].(float64); ok {
			transaction.Saldo = saldo
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

// Helper to convert a map to a JSON reader
func mapToReader(m map[string]interface{}) *strings.Reader {
	body, _ := json.Marshal(m)
	return strings.NewReader(string(body))
}

// Helper to map Elasticsearch source data to a struct
func mapToStruct(source interface{}, dest interface{}) error {
	sourceBytes, err := json.Marshal(source)
	if err != nil {
		return err
	}
	return json.Unmarshal(sourceBytes, dest)
}

func (r *bankAccountTransactionRepository) GetByBatchID(batchID uint) ([]models.BankAccountTransactions, error) {
	var transactions []models.BankAccountTransactions
	err := r.db.Where("BatchID = ?", batchID).Find(&transactions).Error
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve transactions for batch %d: %v", batchID, err)
	}
	return transactions, nil
}
