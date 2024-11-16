package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/xuri/excelize/v2"
)

func main() {
	// Step 1: Extract text from the PDF
	pdfFile := "statement.pdf"
	extractedText, err := extractTextFromPDF(pdfFile)
	if err != nil {
		log.Fatalf("Error extracting text from PDF: %v", err)
	}

	// Debug: Print extracted text
	fmt.Println("Extracted Text:")
	fmt.Println(extractedText)

	// Step 2: Parse the extracted text
	records := parseTransactions(extractedText)

	// Debug: Print parsed transactions
	fmt.Println("Parsed Transactions:")
	for _, record := range records {
		fmt.Println(record)
	}

	// Step 3: Generate the Excel file
	outputFile := "statement.xlsx"
	err = generateExcel(records, outputFile)
	if err != nil {
		log.Fatalf("Error generating Excel file: %v", err)
	}

	fmt.Printf("Excel file successfully created: %s\n", outputFile)
}

// Step 1: Extract text from the PDF using pdfcpu
func extractTextFromPDF(pdfFile string) (string, error) {
	cmd := exec.Command("pdftotext", "-layout", pdfFile, "-") // Extract text with layout
	var out bytes.Buffer
	cmd.Stdout = &out
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("pdftotext error: %s - %v", stderr.String(), err)
	}

	return out.String(), nil
}

// Step 2: Parse transactions from the extracted text

func parseTransactions(text string) [][]string {
	lines := strings.Split(text, "\n")
	var records [][]string
	var currentRecord []string

	for lineNumber, line := range lines {
		line = strings.TrimSpace(line)

		// Skip irrelevant lines
		if line == "" || strings.Contains(line, "CATATAN") || strings.Contains(line, "BERSAMBUNG") ||
			strings.Contains(line, "REKENING GIRO") || strings.Contains(line, "HALAMAN") || strings.Contains(line, "MATA UANG") {
			continue
		}

		// Check if the line starts with a date (e.g., "01/01")
		if len(line) >= 5 && line[2] == '/' && line[5] == '/' {
			// Save the current record before starting a new one
			if len(currentRecord) > 0 {
				records = append(records, currentRecord)
				currentRecord = []string{}
			}

			// Split the line into fields
			fields := strings.Fields(line)
			if len(fields) > 1 {
				currentRecord = append(currentRecord, fields[0])                     // Date
				currentRecord = append(currentRecord, strings.Join(fields[1:], " ")) // Description
			} else {
				fmt.Printf("Warning: Malformed line at %d: %s\n", lineNumber, line)
			}
		} else {
			// Append to the previous record's description if no new date is detected
			if len(currentRecord) > 0 {
				currentRecord[1] += " " + line
			} else {
				fmt.Printf("Warning: Orphaned line at %d: %s\n", lineNumber, line)
			}
		}
	}

	// Add the last record if any
	if len(currentRecord) > 0 {
		records = append(records, currentRecord)
	}

	return records
}

// Step 3: Generate Excel file
func generateExcel(records [][]string, outputFile string) error {
	f := excelize.NewFile()
	sheetName := "Sheet1"
	f.NewSheet(sheetName)

	// Add headers
	headers := []string{"Tanggal", "Keterangan", "Debit", "Credit", "Saldo"}
	for col, header := range headers {
		cell := fmt.Sprintf("%c1", 'A'+col)
		f.SetCellValue(sheetName, cell, header)
	}

	// Add transaction data
	for rowIndex, record := range records {
		for colIndex, value := range record {
			cell := fmt.Sprintf("%c%d", 'A'+colIndex, rowIndex+2)
			f.SetCellValue(sheetName, cell, value)
		}
	}

	// Save the file
	return f.SaveAs(outputFile)
}
