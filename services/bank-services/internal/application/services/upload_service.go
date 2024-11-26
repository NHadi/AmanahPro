package services

import (
	"AmanahPro/services/bank-services/internal/domain/models"
	"AmanahPro/services/bank-services/internal/domain/repositories"
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"AmanahPro/services/bank-services/common/messagebroker"
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
	log.Printf("Starting ParseAndSave for file: %s, AccountID: %d", filePath, accountID)

	// Step 1: Parse CSV file into transactions
	transactions, err := s.parseCSVFile(filePath, accountID, organizationID)
	if err != nil {
		log.Printf("Failed to parse file: %v", err)
		return nil, fmt.Errorf("failed to parse file: %v", err)
	}
	log.Printf("Parsed %d transactions from file: %s", len(transactions), filePath)

	if len(transactions) == 0 {
		log.Printf("No transactions found in the file: %s", filePath)
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
		log.Printf("Failed to save transactions: %v", err)
		return nil, fmt.Errorf("failed to save transactions: %v", err)
	}
	log.Printf("Saved transactions with BatchID: %d", batch.BatchID)

	// Step 4: Retrieve the saved transactions to ensure IDs are populated
	savedTransactions, err := s.transactionRepo.GetByBatchID(batch.BatchID)
	if err != nil {
		log.Printf("Failed to retrieve saved transactions: %v", err)
		return nil, fmt.Errorf("failed to retrieve saved transactions: %v", err)
	}

	// Step 5: Publish all transactions as a batch to RabbitMQ
	err = s.publishBatchToRabbitMQ(savedTransactions)
	if err != nil {
		log.Printf("Failed to publish transactions to RabbitMQ: %v", err)
		return nil, fmt.Errorf("failed to publish transactions to RabbitMQ: %v", err)
	}
	log.Printf("Successfully published transactions to RabbitMQ for BatchID: %d", batch.BatchID)

	return savedTransactions, nil
}

func (s *UploadService) parseCSVFile(filePath string, accountID, organizationID uint) ([]models.BankAccountTransactions, error) {
	log.Printf("Parsing CSV file: %s", filePath)

	// Step 1: Open the CSV file
	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("Error opening file: %v", err)
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
		log.Printf("Error reading file: %v", err)
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
		log.Printf("Could not find transaction data header in the CSV")
		return nil, fmt.Errorf("could not find transaction data header in the CSV")
	}

	log.Printf("Found transaction data header in the CSV at line: %d", startIndex)

	// Step 4: Extract header and transaction data
	headerLine := lines[startIndex-1]
	dataLines := lines[startIndex:]

	header := splitCSVLine(headerLine)
	log.Printf("Parsed header: %v", header)

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
		if len(row) > 0 {
			tanggalTransaksi := strings.SplitN(row[0], ",", 2)
			if len(tanggalTransaksi) == 2 {
				row = append([]string{strings.TrimSpace(tanggalTransaksi[0]), strings.TrimSpace(tanggalTransaksi[1])}, row[1:]...)
			}
		}
		data = append(data, row)
	}
	log.Printf("Processed %d rows of data", len(data))

	transactions := []models.BankAccountTransactions{}

	// Parse and map credit, debit, and saldo into transactions
	for rowIndex, record := range data {
		if len(record) < 5 {
			log.Printf("Skipping malformed row %d: %v", rowIndex, record)
			continue
		}

		jumlah := strings.TrimSpace(record[3])
		log.Printf("Raw jumlah: '%s'", jumlah)

		var credit, debit, saldo float64
		if strings.Contains(jumlah, "CR") {
			jumlah = strings.Replace(jumlah, "CR", "", -1)
			jumlah = strings.Replace(jumlah, ",", "", -1)
			jumlah = strings.TrimSpace(jumlah)
			credit, err = strconv.ParseFloat(jumlah, 64)
			if err != nil {
				log.Printf("Skipping malformed credit: '%s', error: %v", jumlah, err)
				continue
			}
		} else if strings.Contains(jumlah, "DB") {
			jumlah = strings.Replace(jumlah, "DB", "", -1)
			jumlah = strings.Replace(jumlah, ",", "", -1)
			jumlah = strings.TrimSpace(jumlah)
			debit, err = strconv.ParseFloat(jumlah, 64)
			if err != nil {
				log.Printf("Skipping malformed debit: '%s', error: %v", jumlah, err)
				continue
			}
		}

		saldoStr := strings.TrimSpace(record[4])
		saldoStr = strings.Replace(saldoStr, ",", "", -1)
		saldo, err = strconv.ParseFloat(saldoStr, 64)
		if err != nil {
			log.Printf("Skipping malformed saldo: %s", saldoStr)
			continue
		}

		tanggalStr := strings.TrimSpace(record[0])
		fullDateStr := fmt.Sprintf("%s/%d", tanggalStr, time.Now().Year())
		tanggal, err := time.Parse("02/01/2006", fullDateStr)
		if err != nil {
			log.Printf("Skipping malformed date: %s", tanggalStr)
			continue
		}

		keterangan := strings.TrimSpace(strings.Trim(record[1], `"`))
		cabang := strings.TrimSpace(record[2])

		transactions = append(transactions, models.BankAccountTransactions{
			AccountID:      accountID,
			Tanggal:        tanggal,
			Keterangan:     keterangan,
			Cabang:         cabang,
			Credit:         credit,
			Debit:          debit,
			Saldo:          saldo,
			OrganizationId: organizationID,
		})

		log.Printf("Processed row %d: %+v", rowIndex, transactions[len(transactions)-1])
	}

	log.Printf("Successfully parsed %d transactions", len(transactions))
	return transactions, nil
}

func splitCSVLine(line string) []string {
	line = strings.Trim(line, "\"")
	parts := strings.Split(line, "\",\"")
	for i := range parts {
		parts[i] = strings.Trim(parts[i], "\"")
	}
	return parts
}

func (s *UploadService) publishBatchToRabbitMQ(transactions []models.BankAccountTransactions) error {
	log.Printf("Publishing batch of %d transactions to RabbitMQ", len(transactions))

	message, err := json.Marshal(transactions)
	if err != nil {
		log.Printf("Failed to marshal transactions batch: %v", err)
		return fmt.Errorf("failed to marshal transactions batch: %v", err)
	}

	err = s.rabbitPublisher.Publish("transactions_queue", message)
	if err != nil {
		log.Printf("Failed to publish transactions batch to RabbitMQ: %v", err)
		return fmt.Errorf("failed to publish transactions batch to RabbitMQ: %v", err)
	}

	log.Printf("Successfully published transactions batch to RabbitMQ")
	return nil
}
