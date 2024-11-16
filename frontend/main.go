package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/xuri/excelize/v2"
)

func main() {
	// Step 1: Open the CSV file
	csvFile, err := os.Open("transactions.csv")
	if err != nil {
		log.Fatalf("Failed to open CSV file: %v", err)
	}
	defer csvFile.Close()

	// Step 2: Read and parse the file
	records, err := parseCSV(csvFile)
	if err != nil {
		log.Fatalf("Failed to parse CSV file: %v", err)
	}

	// Step 3: Generate the Excel file
	outputFile := "transactions.xlsx"
	err = generateExcel(records, outputFile)
	if err != nil {
		log.Fatalf("Failed to generate Excel file: %v", err)
	}

	fmt.Printf("Excel file successfully created: %s\n", outputFile)
}

func parseCSV(file *os.File) ([][]string, error) {
	scanner := bufio.NewScanner(file)
	var records [][]string

	// Read each line
	for scanner.Scan() {
		line := scanner.Text()

		// Split by comma while handling fields wrapped in single quotes
		fields := strings.Split(line, ",")
		cleanedFields := []string{}
		for _, field := range fields {
			cleanedFields = append(cleanedFields, strings.TrimSpace(strings.Trim(field, "'")))
		}

		// Ensure the line has at least 6 fields (based on the file's structure)
		if len(cleanedFields) >= 6 {
			records = append(records, cleanedFields)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return records, nil
}

func generateExcel(records [][]string, outputFile string) error {
	f := excelize.NewFile()
	sheetName := "Sheet1"
	f.NewSheet(sheetName)

	// Step 4: Write Headers
	headers := []string{"Tanggal", "Keterangan", "Debit", "Credit"}
	for col, header := range headers {
		cell := fmt.Sprintf("%c1", 'A'+col)
		f.SetCellValue(sheetName, cell, header)
	}

	// Step 5: Write Data Rows
	for rowIndex, record := range records[1:] { // Skip the header row
		date := record[0]
		description := record[1]
		debit := ""
		credit := ""
		if record[4] == "DB" {
			debit = record[3]
		} else if record[4] == "CR" {
			credit = record[3]
		}
		balance := record[5]

		// Write each cell
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", rowIndex+2), date)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", rowIndex+2), description)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", rowIndex+2), debit)
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", rowIndex+2), credit)
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", rowIndex+2), balance)
	}

	// Save the file
	return f.SaveAs(outputFile)
}
