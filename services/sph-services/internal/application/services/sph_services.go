package services

import (
	"AmanahPro/services/sph-services/internal/domain/models"
	"AmanahPro/services/sph-services/internal/domain/repositories"
	"AmanahPro/services/sph-services/internal/dto"
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/NHadi/AmanahPro-common/messagebroker"
	"github.com/xuri/excelize/v2"
)

type SphService struct {
	sphRepo         repositories.SphRepository
	sectionRepo     repositories.SphSectionRepository
	detailRepo      repositories.SphDetailRepository
	rabbitPublisher *messagebroker.RabbitMQPublisher
	sphQueueName    string
}

// NewSphService initializes the SphService
func NewSphService(
	sphRepo repositories.SphRepository,
	sectionRepo repositories.SphSectionRepository,
	detailRepo repositories.SphDetailRepository,
	rabbitPublisher *messagebroker.RabbitMQPublisher,
	sphQueueName string,
) *SphService {
	return &SphService{
		sphRepo:         sphRepo,
		sectionRepo:     sectionRepo,
		detailRepo:      detailRepo,
		rabbitPublisher: rabbitPublisher,
		sphQueueName:    sphQueueName,
	}
}

// FilterSPHs filters SPHs by organization ID and optional SPH ID or project ID
func (s *SphService) Filter(organizationID int, sphID *int, projectID *int) ([]models.Sph, error) {
	log.Printf("Filtering SPHs for OrganizationID: %d, SphID: %v, ProjectID: %v", organizationID, sphID, projectID)

	sphs, err := s.sphRepo.Filter(organizationID, sphID, projectID)
	if err != nil {
		log.Printf("Error filtering SPHs: %v", err)
		return nil, fmt.Errorf("failed to filter SPHs: %w", err)
	}

	log.Printf("Found %d SPHs for OrganizationID: %d", len(sphs), organizationID)
	return sphs, nil
}

// Helper: PublishFullReindexEvent sends a RabbitMQ event to re-index the full SPH
func (s *SphService) PublishFullReindexEvent(sphID int) error {
	log.Printf("Triggering re-index for SphID: %d", sphID)

	// Retrieve the full SPH structure for re-indexing
	sph, err := s.sphRepo.GetByID(sphID)
	if err != nil {
		log.Printf("Error retrieving SPH for reindexing: %v", err)
		return fmt.Errorf("error retrieving SPH for reindexing: %w", err)
	}

	if sph == nil {
		log.Printf("Reindex failed because sph id: %d Not Found", sphID)
		return nil
	}

	event := map[string]interface{}{
		"event":   "Reindexed",
		"payload": sph,
		"meta": map[string]interface{}{
			"idField": "SphId",
		},
	}
	if err := s.rabbitPublisher.PublishEvent(s.sphQueueName, event); err != nil {
		log.Printf("Error publishing re-index event for SphID: %d, %v", sphID, err)
		return nil
	}

	log.Printf("Successfully triggered re-index for SphID: %d", sphID)
	return nil
}

// SPH CRUD

func (s *SphService) GetSphByID(sphId int) (*models.Sph, error) {
	log.Printf("Fetching SphSection with ID: %d", sphId)

	sph, err := s.sphRepo.GetByID(sphId)
	if err != nil {
		log.Printf("Error fetching sph: %v", err)
		return nil, fmt.Errorf("failed to fetch sph: %w", err)
	}

	return sph, nil
}

func (s *SphService) CreateSph(sph *models.Sph) error {
	log.Printf("Creating SPH: %+v", sph)
	if err := s.sphRepo.Create(sph); err != nil {
		log.Printf("Error creating SPH: %v", err)
		return fmt.Errorf("error creating SPH: %w", err)
	}

	event := map[string]interface{}{
		"event":   "Created",
		"payload": sph,
		"meta": map[string]interface{}{
			"idField": "SphId",
		},
	}
	if err := s.rabbitPublisher.PublishEvent(s.sphQueueName, event); err != nil {
		log.Printf("Error publishing index event for SphID: %d, %v", sph.SphId, err)
		return nil
	}

	log.Printf("Successfully triggered index for SphID: %d", sph.SphId)
	return nil
}

func (s *SphService) UpdateSph(sph *models.Sph) error {
	log.Printf("Updating SPH: %+v", sph)
	if err := s.sphRepo.Update(sph); err != nil {
		log.Printf("Error updating SPH: %v", err)
		return fmt.Errorf("error updating SPH: %w", err)
	}

	return s.PublishFullReindexEvent(sph.SphId)
}

func (s *SphService) DeleteSph(sphID int) error {
	log.Printf("Deleting SPH ID: %d", sphID)
	if err := s.sphRepo.Delete(sphID); err != nil {
		log.Printf("Error deleting SPH: %v", err)
		return fmt.Errorf("error deleting SPH: %w", err)
	}

	// Send a delete event to RabbitMQ
	event := map[string]interface{}{
		"event":   "Deleted",
		"payload": map[string]int{"SphId": sphID},
		"meta": map[string]interface{}{
			"idField": "SphId",
		},
	}
	if err := s.rabbitPublisher.PublishEvent(s.sphQueueName, event); err != nil {
		log.Printf("Error publishing delete event for SphID: %d, %v", sphID, err)
		return fmt.Errorf("error publishing delete event: %w", err)
	}

	log.Printf("Successfully deleted SPH ID: %d", sphID)
	return nil
}

// Section CRUD
func (s *SphService) CreateSphSection(section *models.SphSection) error {
	log.Printf("Creating SphSection: %+v", section)
	if err := s.sectionRepo.Create(section); err != nil {
		log.Printf("Error creating SphSection: %v", err)
		return fmt.Errorf("error creating SphSection: %w", err)
	}

	return s.PublishFullReindexEvent(section.SphId)
}

func (s *SphService) UpdateSphSection(section *models.SphSection) error {
	log.Printf("Updating SphSection: %+v", section)
	if err := s.sectionRepo.Update(section); err != nil {
		log.Printf("Error updating SphSection: %v", err)
		return fmt.Errorf("error updating SphSection: %w", err)
	}

	return s.PublishFullReindexEvent(section.SphId)
}

func (s *SphService) DeleteSphSection(sectionID int, sphID int) error {
	log.Printf("Deleting SphSection ID: %d", sectionID)
	if err := s.sectionRepo.Delete(sectionID); err != nil {
		log.Printf("Error deleting SphSection: %v", err)
		return fmt.Errorf("error deleting SphSection: %w", err)
	}

	return s.PublishFullReindexEvent(sphID)
}

// Detail CRUD
func (s *SphService) CreateSphDetail(detail *models.SphDetail, sphId int) error {
	log.Printf("Creating SphDetail: %+v", detail)
	if err := s.detailRepo.Create(detail); err != nil {
		log.Printf("Error creating SphDetail: %v", err)
		return fmt.Errorf("error creating SphDetail: %w", err)
	}

	return s.PublishFullReindexEvent(sphId)
}

func (s *SphService) UpdateSphDetail(detail *models.SphDetail, sphId int) error {
	log.Printf("Updating SphDetail: %+v", detail)
	if err := s.detailRepo.Update(detail); err != nil {
		log.Printf("Error updating SphDetail: %v", err)
		return fmt.Errorf("error updating SphDetail: %w", err)
	}

	return s.PublishFullReindexEvent(sphId)
}

func (s *SphService) DeleteSphDetail(detailID int, sphID int) error {
	log.Printf("Deleting SphDetail ID: %d", detailID)
	if err := s.detailRepo.Delete(detailID); err != nil {
		log.Printf("Error deleting SphDetail: %v", err)
		return fmt.Errorf("error deleting SphDetail: %w", err)
	}

	return s.PublishFullReindexEvent(sphID)
}

// GetSphSectionByID retrieves a specific SPH Section by its ID
func (s *SphService) GetSphSectionByID(sectionID int) (*models.SphSection, error) {
	log.Printf("Fetching SphSection with ID: %d", sectionID)

	section, err := s.sectionRepo.GetByID(sectionID)
	if err != nil {
		log.Printf("Error fetching SphSection: %v", err)
		return nil, fmt.Errorf("failed to fetch SphSection: %w", err)
	}

	return section, nil
}

// GetSphDetailByID retrieves a specific SPH Detail by its ID
func (s *SphService) GetSphDetailByID(detailID int) (*models.SphDetail, error) {
	log.Printf("Fetching SphDetail with ID: %d", detailID)

	detail, err := s.detailRepo.GetByID(detailID)
	if err != nil {
		log.Printf("Error fetching SphDetail: %v", err)
		return nil, fmt.Errorf("failed to fetch SphDetail: %w", err)
	}

	return detail, nil
}

func (s *SphService) ImportSphFromExcel(metadata dto.SphImportDTO, fileBytes []byte, organizationID int, userID int) (float64, error) {
	// Load Excel from bytes
	f, err := excelize.OpenReader(bytes.NewReader(fileBytes))
	if err != nil {
		return 0, fmt.Errorf("failed to open Excel file: %v", err)
	}

	// Get all rows from the first sheet
	rows, err := f.GetRows(f.GetSheetName(0))
	if err != nil {
		return 0, fmt.Errorf("failed to read Excel rows: %v", err)
	}

	currentDate := time.Now().Format("2006-01-02")           // Get current date in yyyy-MM-dd format
	parsedDate, err := time.Parse("2006-01-02", currentDate) // Parse it into a time.Time object
	if err != nil {
		return 0, fmt.Errorf("failed to parse date: %v", err)
	}

	// Create SPH record
	sph := models.Sph{
		ProjectId:      metadata.ProjectId,
		Location:       metadata.Location,
		Subject:        metadata.Subject,
		Date:           &models.CustomDate{Time: parsedDate}, // Set the parsed date
		RecepientName:  metadata.RecepientName,
		OrganizationId: &organizationID,
		CreatedBy:      &userID,
	}

	if err := s.sphRepo.Create(&sph); err != nil {
		return 0, fmt.Errorf("failed to create SPH: %v", err)
	}

	var currentSectionID int
	headerFound := false   // Flag to indicate whether the header row has been found
	var grandTotal float64 // Variable to accumulate the grand total

	for rowIndex, row := range rows {
		if len(row) == 0 {
			continue // Skip empty rows
		}

		// Detect header row
		if !headerFound && len(row) >= 6 &&
			strings.EqualFold(strings.TrimSpace(row[0]), "No") &&
			strings.EqualFold(strings.TrimSpace(row[1]), "Area Pemasangan") &&
			strings.EqualFold(strings.TrimSpace(row[2]), "Qty") {
			log.Printf("Header row found at line %d", rowIndex+1)
			headerFound = true
			continue // Skip the header row and proceed to the next row
		}

		// Skip rows until the header row is found
		if !headerFound {
			log.Printf("Skipping row %d: header not found yet", rowIndex+1)
			continue
		}

		// Check if it's a new section based on "No" column
		if isAlphabet(row[0]) {
			if len(row) < 2 { // Ensure section row has enough columns
				log.Printf("Skipping invalid section row at line %d: insufficient columns", rowIndex+1)
				continue
			}
			section := models.SphSection{
				SphId:          sph.SphId,
				SectionTitle:   &row[1],
				OrganizationId: &organizationID,
				CreatedBy:      &userID,
			}
			if err := s.sectionRepo.Create(&section); err != nil {
				return 0, fmt.Errorf("failed to create SPH section: %v", err)
			}
			currentSectionID = section.SphSectionId // Get the inserted SectionId
			continue
		}

		// Parse rows as details
		if isNumeric(row[0]) {
			if len(row) < 7 { // Ensure detail row has enough columns
				log.Printf("Skipping invalid detail row at line %d: insufficient columns", rowIndex+1)
				continue
			}

			// Parse TotalPrice
			totalPrice := parseFloat(row[6])
			if totalPrice != nil {
				grandTotal += *totalPrice // Add to grand total
			}

			unitPrice := parseFloat(row[4])
			discountedPrice := parseFloat(row[5])
			discountPercentage := 0.0

			if unitPrice != nil && discountedPrice != nil {
				// Dereference the pointers to get the float64 values
				discountPercentage = ((*unitPrice - *discountedPrice) / *unitPrice) * 100

			}

			detail := models.SphDetail{
				SectionId:       currentSectionID, // Use the actual SectionId
				ItemDescription: &row[1],
				Quantity:        parseFloat(row[2]),
				Unit:            &row[3],
				UnitPrice:       parseFloat(row[4]),
				DiscountPrice:   &discountPercentage,
				TotalPrice:      totalPrice,
				OrganizationId:  &organizationID,
				CreatedBy:       &userID,
			}
			if err := s.detailRepo.Create(&detail); err != nil {
				return 0, fmt.Errorf("failed to create SPH detail: %v", err)
			}
		}
	}

	// Log the grand total for debugging
	log.Printf("Grand Total: %.2f", grandTotal)
	sph.Total = &grandTotal
	if err := s.sphRepo.Update(&sph); err != nil {
		log.Printf("Error updating SPH: %v", err)
	}
	// Successfully processed the Excel and updated the project
	s.PublishFullReindexEvent(sph.SphId)

	// Return the grand total along with nil error (indicating success)
	return grandTotal, nil
}

// Helper function to check if a string is an alphabet
func isAlphabet(input string) bool {
	if len(input) == 0 {
		return false
	}
	return unicode.IsLetter(rune(input[0]))
}

// Helper function to check if a string is numeric
func isNumeric(input string) bool {
	_, err := strconv.Atoi(input)
	return err == nil
}

// Helper function to parse a string into a float pointer
func parseFloat(input string) *float64 {
	// Trim spaces and unwanted characters
	cleaned := strings.TrimSpace(strings.ReplaceAll(input, ",", ""))

	// Attempt to parse the cleaned string
	value, err := strconv.ParseFloat(cleaned, 64)
	if err != nil {
		return nil // Return nil if parsing fails
	}
	return &value
}
