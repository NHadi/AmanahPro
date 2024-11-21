package services

import (
	"AmanahPro/services/bank-services/internal/domain/models"
	"AmanahPro/services/bank-services/internal/domain/repositories"
	"AmanahPro/services/bank-services/internal/infrastructure/messagebroker"
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type UploadService struct {
	transactionRepo repositories.BankAccountTransactionRepository
	batchRepo       repositories.BatchRepository
	rabbitPublisher *messagebroker.RabbitPublisher
}

func NewUploadService(
	transactionRepo repositories.BankAccountTransactionRepository,
	batchRepo repositories.BatchRepository,
	rabbitPublisher *messagebroker.RabbitPublisher,
) *UploadService {
	return &UploadService{
		transactionRepo: transactionRepo,
		batchRepo:       batchRepo,
		rabbitPublisher: rabbitPublisher,
	}
}

// ParseAndSave parses the uploaded CSV file, creates a batch, and saves transactions
func (s *UploadService) ParseAndSave(filePath string, accountID uint, periodeStart, periodeEnd time.Time, uploadedBy string) ([]models.BankAccountTransactions, error) {
	// Step 1: Parse CSV file into transactions
	transactions, err := s.parseCSVFile(filePath, accountID)
	if err != nil {
		return nil, fmt.Errorf("failed to parse file: %v", err)
	}

	if len(transactions) == 0 {
		return nil, errors.New("no transactions found in the file")
	}

	// Step 2: Create a new UploadBatch
	batch := models.UploadBatch{
		AccountID:    accountID,
		FileName:     filePath,
		PeriodeStart: periodeStart,
		PeriodeEnd:   periodeEnd,
		UploadedBy:   uploadedBy,
	}

	// Save the batch and get the generated BatchID
	err = s.batchRepo.Create(&batch)
	if err != nil {
		return nil, fmt.Errorf("failed to save batch: %v", err)
	}

	// Step 3: Assign the BatchID to each transaction
	for i := range transactions {
		transactions[i].BatchID = batch.BatchID
	}

	// Step 4: Save transactions in the database with rollback support
	err = s.transactionRepo.InsertWithRollback(&batch, transactions)
	if err != nil {
		return nil, fmt.Errorf("failed to save transactions: %v", err)
	}

	// Step 5: Publish all transactions as a batch to RabbitMQ
	err = s.publishBatchToRabbitMQ(transactions)
	if err != nil {
		return nil, fmt.Errorf("failed to publish transactions to RabbitMQ: %v", err)
	}

	return transactions, nil
}

// parseCSVFile parses a CSV file and returns a list of transactions
func (s *UploadService) parseCSVFile(filePath string, accountID uint) ([]models.BankAccountTransactions, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	var transactions []models.BankAccountTransactions
	scanner := bufio.NewScanner(file)

	// Get the current year
	currentYear := time.Now().Year()

	// Skip irrelevant lines until "Tanggal Transaksi" is found
	isTransactionSection := false

	for scanner.Scan() {
		line := scanner.Text()

		// Skip lines with metadata
		if !isTransactionSection {
			if strings.Contains(line, "Tanggal Transaksi") {
				isTransactionSection = true
			}
			continue
		}

		// Skip any remaining irrelevant lines
		if strings.TrimSpace(line) == "" || strings.Contains(line, "Saldo Awal") || strings.Contains(line, "Saldo Akhir") {
			continue
		}

		// Split the CSV line into parts
		parts := strings.Split(line, ",")

		// Validate the line structure
		if len(parts) < 6 {
			continue // Skip lines with insufficient fields
		}

		// Clean the date field by removing extra quotes
		tanggalStr := strings.Trim(strings.TrimSpace(parts[0]), "\"")
		fullDateStr := fmt.Sprintf("%s/%d", tanggalStr, currentYear)

		// Parse the cleaned date
		tanggal, err := time.Parse("02/01/2006", fullDateStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse date '%s': %v", fullDateStr, err)
		}

		// Clean and parse other fields
		keterangan := strings.Trim(strings.TrimSpace(parts[1]), "\"")
		cabang := strings.Trim(strings.TrimSpace(parts[2]), "\"")
		credit, _ := strconv.ParseFloat(strings.ReplaceAll(strings.Trim(strings.TrimSpace(parts[3]), "\""), "CR", ""), 64)
		debit, _ := strconv.ParseFloat(strings.ReplaceAll(strings.Trim(strings.TrimSpace(parts[4]), "\""), "DB", ""), 64)
		saldo, _ := strconv.ParseFloat(strings.Trim(strings.TrimSpace(parts[5]), "\""), 64)

		// Add to transactions
		transactions = append(transactions, models.BankAccountTransactions{
			AccountID:  accountID,
			Tanggal:    tanggal,
			Keterangan: keterangan,
			Cabang:     cabang,
			Credit:     credit,
			Debit:      debit,
			Saldo:      saldo,
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return transactions, nil
}

// publishToRabbitMQ publishes a transaction to RabbitMQ for downstream consumers
func (s *UploadService) publishBatchToRabbitMQ(transactions []models.BankAccountTransactions) error {
	// Convert the batch of transactions into JSON
	message, err := json.Marshal(transactions)
	if err != nil {
		return fmt.Errorf("failed to marshal transactions batch: %v", err)
	}

	// Publish the batch message to RabbitMQ
	err = s.rabbitPublisher.Publish(message)
	if err != nil {
		return fmt.Errorf("failed to publish transactions batch to RabbitMQ: %v", err)
	}

	return nil
}
