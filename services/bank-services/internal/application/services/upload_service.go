package services

import (
	"AmanahPro/services/bank-services/internal/domain/models"
	"AmanahPro/services/bank-services/internal/domain/repositories"
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/NHadi/AmanahPro-common/messagebroker"
)

type UploadService struct {
	transactionRepo repositories.BankAccountTransactionRepository
	batchRepo       repositories.BatchRepository
	rabbitPublisher *messagebroker.RabbitMQPublisher
}

func NewUploadService(
	transactionRepo repositories.BankAccountTransactionRepository,
	batchRepo repositories.BatchRepository,
	rabbitPublisher *messagebroker.RabbitMQPublisher,
) *UploadService {
	return &UploadService{
		transactionRepo: transactionRepo,
		batchRepo:       batchRepo,
		rabbitPublisher: rabbitPublisher,
	}
}

func (s *UploadService) ParseAndSave(filePath string, organizationID, accountID, year, month uint, uploadedBy string) ([]models.BankAccountTransactions, error) {
	// Step 1: Parse CSV file into transactions
	transactions, err := s.parseCSVFile(filePath, accountID, organizationID)
	if err != nil {
		return nil, fmt.Errorf("failed to parse file: %v", err)
	}

	if len(transactions) == 0 {
		return nil, errors.New("no transactions found in the file")
	}

	// Step 2: Create a new UploadBatch
	batch := models.UploadBatch{
		AccountID:      accountID,
		FileName:       filePath,
		Year:           year,
		Month:          month,
		UploadedBy:     uploadedBy,
		OrganizationId: organizationID,
	}

	// Step 3: Save transactions in the database with rollback support
	err = s.transactionRepo.InsertWithRollback(&batch, transactions)
	if err != nil {
		return nil, fmt.Errorf("failed to save transactions: %v", err)
	}

	// Step 4: Retrieve the saved transactions to ensure IDs are populated
	savedTransactions, err := s.transactionRepo.GetByBatchID(batch.BatchID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve saved transactions: %v", err)
	}

	// Step 5: Publish all transactions as a batch to RabbitMQ
	err = s.publishBatchToRabbitMQ(savedTransactions)
	if err != nil {
		return nil, fmt.Errorf("failed to publish transactions to RabbitMQ: %v", err)
	}

	return savedTransactions, nil
}

func (s *UploadService) parseCSVFile(filePath string, accountID, organizationId uint) ([]models.BankAccountTransactions, error) {
	// Step 1: Open the CSV file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// Step 2: Read the file line by line
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	// Step 3: Locate the starting line for the transaction data
	startIndex := -1
	for i, line := range lines {
		if strings.Contains(line, "Tanggal Transaksi") {
			startIndex = i + 1 // Data starts from the next line
			break
		}
	}

	if startIndex == -1 {
		return nil, fmt.Errorf("could not find transaction data header in the CSV")
	}

	// Step 4: Extract header and transaction data
	headerLine := lines[startIndex-1]
	dataLines := lines[startIndex:]

	// Split header and data into columns
	header := splitCSVLine(headerLine)
	data := [][]string{}
	for _, line := range dataLines {
		if strings.TrimSpace(line) == "" || // Skip empty lines
			strings.HasPrefix(line, "Saldo Awal") || // Skip "Saldo Awal"
			strings.HasPrefix(line, "Mutasi Debet") || // Skip "Mutasi Debet"
			strings.HasPrefix(line, "Mutasi Kredit") || // Skip "Mutasi Kredit"
			strings.HasPrefix(line, "Saldo Akhir") { // Skip "Saldo Akhir"
			continue
		}
		row := splitCSVLine(line)
		// Split the first field (Tanggal Transaksi) into two fields
		if len(row) > 0 {
			tanggalTransaksi := strings.SplitN(row[0], ",", 2)
			if len(tanggalTransaksi) == 2 {
				row = append([]string{strings.TrimSpace(tanggalTransaksi[0]), strings.TrimSpace(tanggalTransaksi[1])}, row[1:]...)
			}
		}
		data = append(data, row)
	}
	// Step 5: Modify header to include Credit and Debit columns
	if len(header) > 0 {
		header = append([]string{"Tanggal", "Transaksi", "Cabang", "Credit", "Debit", "Saldo"})
	}

	var transactions []models.BankAccountTransactions

	// Parse and map credit, debit, and saldo into transactions
	for rowIndex, record := range data {
		credit := 0.0
		debit := 0.0
		saldo := 0.0

		// Ensure record has sufficient fields to process
		if len(record) < 5 {
			fmt.Printf("Skipping malformed row %d: %v\n", rowIndex, record)
			continue
		}

		// Check the "Jumlah" column (4th index) for "CR" or "DB" and extract the value
		jumlah := strings.TrimSpace(record[3])
		fmt.Printf("Raw jumlah: '%s'\n", jumlah) // Debugging step

		if strings.Contains(jumlah, "CR") {
			jumlah = strings.Replace(jumlah, "CR", "", -1)
			jumlah = strings.Replace(jumlah, ",", "", -1) // Remove commas for parsing
			jumlah = strings.TrimSpace(jumlah)            // Ensure no leading/trailing spaces
			credit, err = strconv.ParseFloat(jumlah, 64)
			if err != nil {
				fmt.Printf("Skipping malformed credit: '%s', error: %v\n", jumlah, err)
				continue
			}
		} else if strings.Contains(jumlah, "DB") {
			jumlah = strings.Replace(jumlah, "DB", "", -1)
			jumlah = strings.Replace(jumlah, ",", "", -1) // Remove commas for parsing
			jumlah = strings.TrimSpace(jumlah)            // Ensure no leading/trailing spaces
			debit, err = strconv.ParseFloat(jumlah, 64)
			if err != nil {
				fmt.Printf("Skipping malformed debit: '%s', error: %v\n", jumlah, err)
				continue
			}
		}

		// Parse saldo (5th index)
		saldoStr := strings.TrimSpace(record[4])
		saldoStr = strings.Replace(saldoStr, ",", "", -1) // Remove commas for parsing
		saldo, err = strconv.ParseFloat(saldoStr, 64)
		if err != nil {
			fmt.Printf("Skipping malformed saldo: %s\n", saldoStr)
			continue
		}

		// Parse date
		tanggalStr := strings.TrimSpace(record[0])
		fullDateStr := fmt.Sprintf("%s/%d", tanggalStr, time.Now().Year()) // Append the current year
		tanggal, err := time.Parse("02/01/2006", fullDateStr)
		if err != nil {
			fmt.Printf("Skipping malformed date: %s\n", tanggalStr)
			continue
		}

		// Parse other fields
		keterangan := strings.TrimSpace(strings.Trim(record[1], `"`))
		cabang := strings.TrimSpace(record[2])

		// Append to transactions slice
		transactions = append(transactions, models.BankAccountTransactions{
			AccountID:      accountID,
			Tanggal:        tanggal,
			Keterangan:     keterangan,
			Cabang:         cabang,
			Credit:         credit,
			Debit:          debit,
			Saldo:          saldo,
			OrganizationId: organizationId,
		})

		fmt.Printf("Processed row %d: Credit: %.2f, Debit: %.2f, Saldo: %.2f\n", rowIndex, credit, debit, saldo)
	}

	return transactions, nil
}

// splitCSVLine handles splitting of CSV lines with quoted fields
func splitCSVLine(line string) []string {
	line = strings.Trim(line, "\"")
	parts := strings.Split(line, "\",\"")
	for i := range parts {
		parts[i] = strings.Trim(parts[i], "\"")
	}
	return parts
}

// publishToRabbitMQ publishes a transaction to RabbitMQ for downstream consumers
func (s *UploadService) publishBatchToRabbitMQ(transactions []models.BankAccountTransactions) error {
	// Convert the batch of transactions into JSON
	message, err := json.Marshal(transactions)
	if err != nil {
		return fmt.Errorf("failed to marshal transactions batch: %v", err)
	}

	// Publish the batch message to RabbitMQ
	err = s.rabbitPublisher.Publish("transactions_queue", message)
	if err != nil {
		return fmt.Errorf("failed to publish transactions batch to RabbitMQ: %v", err)
	}

	return nil
}
