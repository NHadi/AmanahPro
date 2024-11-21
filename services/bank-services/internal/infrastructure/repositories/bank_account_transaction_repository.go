package repositories

import (
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

// GetTransactionsByBankAndPeriod fetches transactions from Elasticsearch
func (r *bankAccountTransactionRepository) GetTransactionsByBankAndPeriod(bankID uint, periodeStart, periodeEnd time.Time) ([]models.BankAccountTransactions, error) {
	// Build the Elasticsearch query
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{
					{
						"term": map[string]interface{}{
							"bank_id": bankID,
						},
					},
					{
						"range": map[string]interface{}{
							"transaction_date": map[string]interface{}{
								"gte": periodeStart.Format(time.RFC3339),
								"lte": periodeEnd.Format(time.RFC3339),
							},
						},
					},
				},
			},
		},
	}

	// Perform the search
	res, err := r.esClient.Search(
		r.esClient.Search.WithIndex(r.esIndex),
		r.esClient.Search.WithBody(mapToReader(query)),
		r.esClient.Search.WithContext(context.Background()),
	)
	if err != nil {
		return nil, fmt.Errorf("error executing search query: %v", err)
	}
	defer res.Body.Close()

	// Parse the response
	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error parsing Elasticsearch response: %v", err)
	}

	// Extract hits
	hits := result["hits"].(map[string]interface{})["hits"].([]interface{})
	var transactions []models.BankAccountTransactions
	for _, hit := range hits {
		source := hit.(map[string]interface{})["_source"]
		transaction := models.BankAccountTransactions{}
		if err := mapToStruct(source, &transaction); err != nil {
			return nil, fmt.Errorf("error converting source to transaction: %v", err)
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
