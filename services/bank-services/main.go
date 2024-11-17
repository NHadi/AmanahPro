package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/xuri/excelize/v2"
)

func extractTransactionsToExcel(inputFilePath, outputFilePath string) error {
	// Step 1: Open the CSV file
	file, err := os.Open(inputFilePath)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// Step 2: Read the file line by line
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %v", err)
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
		return fmt.Errorf("could not find transaction data header in the CSV")
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

	// Step 6: Create a new Excel file
	f := excelize.NewFile()
	sheetName := "Transactions"
	f.SetSheetName("Sheet1", sheetName)

	// Write header row
	for colIndex, headerCell := range header {
		cell, _ := excelize.CoordinatesToCellName(colIndex+1, 1)
		f.SetCellValue(sheetName, cell, strings.TrimSpace(headerCell))
	}

	// Write transaction data rows
	for rowIndex, record := range data {
		credit := ""
		debit := ""
		saldo := ""

		// Ensure record has sufficient fields to process
		if len(record) < 3 {
			fmt.Printf("Skipping malformed row %d: %v\n", rowIndex, record)
			continue
		}

		// Check the "Jumlah" column (4th index) for "CR" or "DB" and extract "Saldo"
		if len(record) > 3 {
			jumlah := record[3]
			if strings.Contains(jumlah, "CR") {
				credit = strings.TrimSpace(strings.Replace(jumlah, "CR", "", -1))
			} else if strings.Contains(jumlah, "DB") {
				debit = strings.TrimSpace(strings.Replace(jumlah, "DB", "", -1))
			}

			// Preserve the original "Saldo" value
			if len(record) > 4 {
				saldo = record[4]
			}
		}

		// Safely truncate record if it has more than 5 fields
		if len(record) > 5 {
			record = append(record[:3], record[5:]...)
		} else {
			record = record[:3] // Ensure it has at least 3 fields
		}

		// Create a new row with Credit, Debit, and Saldo columns added
		newRow := append(record[:3], credit, debit, saldo)

		// Write row to the Excel sheet
		for colIndex, cellValue := range newRow {
			cell, _ := excelize.CoordinatesToCellName(colIndex+1, rowIndex+2)
			f.SetCellValue(sheetName, cell, strings.Trim(cellValue, "\""))
		}
	}

	// Get the sheet index and set it as active
	sheetIndex, err := f.GetSheetIndex(sheetName)
	if err != nil {
		return fmt.Errorf("error getting sheet index: %v", err)
	}
	f.SetActiveSheet(sheetIndex)

	// Apply auto-filter
	if err := f.AutoFilter(sheetName, "A1:F1", nil); err != nil {
		return fmt.Errorf("error applying auto-filter: %v", err)
	}

	// Auto-resize columns based on content
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return fmt.Errorf("error getting rows: %v", err)
	}

	for colIdx := 0; colIdx < len(rows[0]); colIdx++ {
		maxWidth := 10.0 // Default minimum width
		for _, row := range rows {
			if colIdx < len(row) {
				cellContent := row[colIdx]
				contentWidth := float64(len(cellContent))
				if contentWidth > maxWidth {
					maxWidth = contentWidth
				}
			}
		}
		colName, _ := excelize.ColumnNumberToName(colIdx + 1)
		if err := f.SetColWidth(sheetName, colName, colName, maxWidth+2); err != nil { // Add padding
			return fmt.Errorf("error setting column width: %v", err)
		}
	}

	// Enable text wrapping for all cells
	style, err := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			WrapText: true,
		},
	})
	if err != nil {
		return fmt.Errorf("error creating style: %v", err)
	}

	// Apply the style to the entire sheet
	if err := f.SetCellStyle(sheetName, "A1", "F1000", style); err != nil { // Adjust range as needed
		return fmt.Errorf("error applying style: %v", err)
	}

	// Step 7: Save the Excel file
	if err := f.SaveAs(outputFilePath); err != nil {
		return fmt.Errorf("error saving Excel file: %v", err)
	}

	fmt.Printf("Transaction data successfully extracted to %s\n", outputFilePath)
	return nil
}

// splitCSVLine splits a single CSV line into fields, handling quoted values
func splitCSVLine(line string) []string {
	line = strings.Trim(line, "\"") // Trim outer quotes
	parts := strings.Split(line, "\",\"")
	for i := range parts {
		parts[i] = strings.Trim(parts[i], "\"") // Remove double quotes from individual fields
	}
	return parts
}

func main() {
	inputFilePath := "input.csv"    // Path to your CSV file
	outputFilePath := "output.xlsx" // Path to the output Excel file

	if err := extractTransactionsToExcel(inputFilePath, outputFilePath); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
