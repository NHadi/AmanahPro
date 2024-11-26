package services

import (
	"AmanahPro/services/bank-services/internal/application/dto"
	"AmanahPro/services/bank-services/internal/domain/models"
	"AmanahPro/services/bank-services/internal/domain/repositories"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-redis/redis/v8"
)

type ReconciliationService struct {
	esClient                         *elasticsearch.Client
	esIndex                          string
	redisClient                      *redis.Client
	bankAccountTransactionRepository repositories.BankAccountTransactionRepository
}

const (
	batchSize  = 1000        // Number of documents per batch
	maxRetries = 3           // Maximum retries for failed batches
	retryDelay = time.Second // Initial delay between retries
)

// NewReconciliationService creates a new instance of the ReconciliationService
func NewReconciliationService(
	esClient *elasticsearch.Client,
	esIndex string,
	redisClient *redis.Client,
	bankAccountTransactionRepository repositories.BankAccountTransactionRepository,
) *ReconciliationService {
	return &ReconciliationService{
		esClient:                         esClient,
		esIndex:                          esIndex,
		redisClient:                      redisClient,
		bankAccountTransactionRepository: bankAccountTransactionRepository,
	}
}

// PerformReconciliation executes incremental reconciliation
func (r *ReconciliationService) PerformReconciliation() error {
	log.Println("Starting reconciliation process...")
	ctx := context.Background()

	// Get the last processed timestamp from Redis
	lastProcessed, err := r.GetLastProcessedTimestamp(ctx, "last_reconciliation_timestamp")
	if err != nil {
		log.Printf("Error retrieving last processed timestamp: %v", err)
		return fmt.Errorf("failed to get last processed timestamp: %w", err)
	}
	log.Printf("Last processed timestamp: %v", lastProcessed)

	offset := 0
	for {
		// Fetch a batch of updated transactions
		updatedTransactions, err := r.bankAccountTransactionRepository.FindUpdatedAfter(lastProcessed, batchSize, offset)
		if err != nil {
			log.Printf("Error fetching updated transactions (offset=%d): %v", offset, err)
			return fmt.Errorf("failed to fetch updated transactions: %w", err)
		}
		log.Printf("Fetched %d transactions for reconciliation (offset=%d)", len(updatedTransactions), offset)

		// Exit if no more transactions
		if len(updatedTransactions) == 0 {
			log.Println("No more transactions to process.")
			break
		}

		// Bulk index the transactions into Elasticsearch
		err = r.BulkIndexTransactions(updatedTransactions)
		if err != nil {
			log.Printf("Bulk indexing failed for batch (offset=%d): %v", offset, err)
			return err
		}
		log.Printf("Successfully indexed %d transactions (offset=%d)", len(updatedTransactions), offset)

		// Update offset and the last processed timestamp
		offset += len(updatedTransactions)
		lastProcessed = updatedTransactions[len(updatedTransactions)-1].UpdatedAt
		err = r.SetLastProcessedTimestamp(ctx, "last_reconciliation_timestamp", lastProcessed)
		if err != nil {
			log.Printf("Failed to update last processed timestamp: %v", err)
		}
	}

	// Final update of the last processed timestamp
	finalTimestamp := time.Now()
	err = r.SetLastProcessedTimestamp(ctx, "last_reconciliation_timestamp", finalTimestamp)
	if err != nil {
		log.Printf("Failed to update final timestamp: %v", err)
	} else {
		log.Printf("Reconciliation completed. Final timestamp set to: %v", finalTimestamp)
	}
	return err
}

// GetLastProcessedTimestamp fetches the last processed timestamp from Redis
func (r *ReconciliationService) GetLastProcessedTimestamp(ctx context.Context, key string) (time.Time, error) {
	timestampStr, err := r.redisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		now := time.Now()
		log.Printf("No last reconciliation timestamp found. Defaulting to now: %v", now)
		return now, nil
	} else if err != nil {
		log.Printf("Error retrieving last processed timestamp from Redis: %v", err)
		return time.Time{}, err
	}

	timestamp, err := time.Parse(time.RFC3339, timestampStr)
	if err != nil {
		log.Printf("Invalid timestamp format in Redis: %s, error: %v", timestampStr, err)
		return time.Time{}, fmt.Errorf("invalid timestamp format: %w", err)
	}
	return timestamp, nil
}

// SetLastProcessedTimestamp updates the last processed timestamp in Redis
func (r *ReconciliationService) SetLastProcessedTimestamp(ctx context.Context, key string, timestamp time.Time) error {
	err := r.redisClient.Set(ctx, key, timestamp, 0).Err()
	if err != nil {
		log.Printf("Error setting last processed timestamp in Redis: %v", err)
		return fmt.Errorf("failed to set last processed timestamp: %w", err)
	}
	log.Printf("Last processed timestamp updated in Redis: %v", timestamp)
	return nil
}

// BulkIndexTransactions indexes multiple transactions in batches
func (r *ReconciliationService) BulkIndexTransactions(transactions []models.BankAccountTransactions) error {
	for i := 0; i < len(transactions); i += batchSize {
		end := i + batchSize
		if end > len(transactions) {
			end = len(transactions)
		}

		batch := transactions[i:end]
		log.Printf("Processing batch %d-%d", i, end)
		err := r.processBatchWithRetry(batch)
		if err != nil {
			log.Printf("Failed to process batch %d-%d: %v", i, end, err)
			return err
		}
		log.Printf("Batch %d-%d processed successfully", i, end)
	}
	return nil
}

// processBatchWithRetry handles a single batch with retry logic
func (r *ReconciliationService) processBatchWithRetry(batch []models.BankAccountTransactions) error {
	for attempt := 1; attempt <= maxRetries; attempt++ {
		err := r.processBatch(batch)
		if err == nil {
			log.Printf("Batch processed successfully on attempt %d", attempt)
			return nil
		}

		log.Printf("Batch failed (attempt %d/%d): %v", attempt, maxRetries, err)
		time.Sleep(retryDelay * time.Duration(attempt)) // Exponential backoff
	}
	return fmt.Errorf("batch failed after %d attempts", maxRetries)
}

// processBatch performs the actual bulk indexing for a batch
func (r *ReconciliationService) processBatch(batch []models.BankAccountTransactions) error {
	var bulkPayload strings.Builder
	failedDocs := []string{}

	for _, transaction := range batch {
		dtoTransaction := dto.BankAccountTransactionDTO{
			ID:             transaction.ID,
			AccountID:      transaction.AccountID,
			BatchID:        transaction.BatchID,
			Tanggal:        transaction.Tanggal.Format("2006-01-02"),
			Keterangan:     strings.Trim(transaction.Keterangan, `"`),
			Cabang:         transaction.Cabang,
			Credit:         transaction.Credit,
			Debit:          transaction.Debit,
			Saldo:          transaction.Saldo,
			OrganizationId: transaction.OrganizationId,
			UpdatedAt:      transaction.UpdatedAt,
		}

		data, err := json.Marshal(dtoTransaction)
		if err != nil {
			log.Printf("Failed to marshal transaction ID %d: %v", transaction.ID, err)
			failedDocs = append(failedDocs, fmt.Sprintf("%d", transaction.ID))
			continue
		}

		bulkMetadata := fmt.Sprintf(`{ "index": { "_index": "%s", "_id": "%d" } }`, r.esIndex, transaction.ID)
		bulkPayload.WriteString(bulkMetadata + "\n")
		bulkPayload.WriteString(string(data) + "\n")
	}

	res, err := r.esClient.Bulk(strings.NewReader(bulkPayload.String()), r.esClient.Bulk.WithRefresh("true"))
	if err != nil {
		log.Printf("Bulk request failed: %v", err)
		return fmt.Errorf("bulk indexing failed: %w", err)
	}

	failures, err := parseBulkResponse(res.Body)
	if err != nil {
		log.Printf("Error parsing bulk response: %v", err)
		return fmt.Errorf("failed to parse bulk response: %w", err)
	}

	for _, failure := range failures {
		log.Printf("Failed document: ID=%s, Reason=%s", failure.ID, failure.Reason)
	}

	if len(failures) > 0 {
		log.Printf("Some documents failed in bulk request: %d failures", len(failures))
		return fmt.Errorf("some documents failed in bulk request")
	}
	return nil
}

// parseBulkResponse parses the Elasticsearch bulk API response
func parseBulkResponse(body io.ReadCloser) ([]struct {
	ID     string
	Reason string
}, error) {
	defer body.Close()
	var bulkResponse struct {
		Items []struct {
			Index struct {
				Status int    `json:"status"`
				ID     string `json:"_id"`
				Error  struct {
					Type   string `json:"type"`
					Reason string `json:"reason"`
				} `json:"error,omitempty"`
			} `json:"index"`
		} `json:"items"`
	}
	if err := json.NewDecoder(body).Decode(&bulkResponse); err != nil {
		return nil, fmt.Errorf("failed to decode bulk response: %w", err)
	}

	failures := []struct {
		ID     string
		Reason string
	}{}

	for _, item := range bulkResponse.Items {
		if item.Index.Status >= 400 {
			failures = append(failures, struct {
				ID     string
				Reason string
			}{
				ID:     item.Index.ID,
				Reason: item.Index.Error.Reason,
			})
		}
	}
	return failures, nil
}
