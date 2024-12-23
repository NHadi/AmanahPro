package services

import (
	"AmanahPro/services/spk-services/internal/domain/models"
	"AmanahPro/services/spk-services/internal/domain/repositories"
	"AmanahPro/services/spk-services/internal/dto"
	"bytes"
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
	"unicode"

	commonServices "github.com/NHadi/AmanahPro-common/services"
	"github.com/xuri/excelize/v2"

	"github.com/NHadi/AmanahPro-common/messagebroker"
	pb "github.com/NHadi/AmanahPro-common/protos"
)

type SpkService struct {
	spkRepo         repositories.SPKRepository
	spkSectionRepo  repositories.SPKSectionRepository
	detailRepo      repositories.SPKDetailRepository
	rabbitPublisher *messagebroker.RabbitMQPublisher
	spkQueueName    string
	sphGrpcClient   pb.SphServiceClient // gRPC client for SPH,
	auditTrail      *commonServices.AuditTrailService
}

// NewSpkService initializes the SpkService
func NewSpkService(
	spkRepo repositories.SPKRepository,
	spkSectionRepo repositories.SPKSectionRepository,
	detailRepo repositories.SPKDetailRepository,
	rabbitPublisher *messagebroker.RabbitMQPublisher,
	spkQueueName string,
	sphGrpcClient pb.SphServiceClient, // Inject SPH gRPC client
	auditTrail *commonServices.AuditTrailService,
) *SpkService {
	return &SpkService{
		spkRepo:         spkRepo,
		spkSectionRepo:  spkSectionRepo,
		detailRepo:      detailRepo,
		rabbitPublisher: rabbitPublisher,
		spkQueueName:    spkQueueName,
		sphGrpcClient:   sphGrpcClient,
		auditTrail:      auditTrail,
	}
}

// FilterSPKs filters SPKs by organization ID and optional SPK ID or project ID
func (s *SpkService) Filter(organizationID int, spkID *int, projectID *int) ([]models.SPK, error) {
	log.Printf("Filtering SPKs for OrganizationID: %d, SpkID: %v, ProjectID: %v", organizationID, spkID, projectID)

	spks, err := s.spkRepo.Filter(organizationID, spkID, projectID)
	if err != nil {
		log.Printf("Error filtering SPKs: %v", err)
		return nil, fmt.Errorf("failed to filter SPKs: %w", err)
	}

	log.Printf("Found %d SPKs for OrganizationID: %d", len(spks), organizationID)
	return spks, nil
}

// Helper: PublishFullReindexEvent sends a RabbitMQ event to re-index the full SPK
func (s *SpkService) PublishFullReindexEvent(spkID int, traceID string, userID int) error {
	log.Printf("Triggering re-index for SpkID: %d", spkID)

	// Retrieve the full SPK structure for re-indexing
	spk, err := s.spkRepo.GetByID(spkID, true)
	if err != nil {
		log.Printf("Error retrieving SPK for reindexing: %v", err)
		return fmt.Errorf("error retrieving SPK for reindexing: %w", err)
	}

	if spk == nil {
		log.Printf("Reindex failed because spk ID: %d Not Found", spkID)
		return nil
	}

	event := map[string]interface{}{
		"event":   "Reindexed",
		"payload": spk,
		"meta": map[string]interface{}{
			"traceId": traceID,
			"action":  "REINDEX",
			"idField": "SpkId",
			"userId":  userID,
		},
	}
	if err := s.rabbitPublisher.PublishEvent(s.spkQueueName, event); err != nil {
		log.Printf("Error publishing re-index event for SpkID: %d, %v", spkID, err)
		return fmt.Errorf("error publishing re-index event: %w", err)
	}

	log.Printf("Successfully triggered re-index for SpkID: %d", spkID)
	return nil
}

// CRUD Operations for SPK
func (s *SpkService) GetSpkByID(spkID int) (*models.SPK, error) {
	log.Printf("Fetching SPK with ID: %d", spkID)

	spk, err := s.spkRepo.GetByID(spkID, false)
	if err != nil {
		log.Printf("Error fetching SPK: %v", err)
		return nil, fmt.Errorf("failed to fetch SPK: %w", err)
	}

	return spk, nil
}

// CreateSpk creates an SPK and populates sections and details from SPH
func (s *SpkService) CreateSpk(spk *models.SPK, sphId int32, traceID string) error {
	log.Printf("Creating SPK: %+v with SPH ID: %d", spk, sphId)

	// Call SPH gRPC service to get sections and details
	log.Printf("Calling SPH gRPC service for SPH ID: %d", sphId)
	sphDetailsResponse, err := s.sphGrpcClient.GetSphDetails(context.Background(), &pb.GetSphDetailsRequest{SphId: sphId})
	if err != nil {
		log.Printf("Failed to fetch SPH details from gRPC: %v", err)
		return fmt.Errorf("failed to fetch SPH details from gRPC: %w", err)
	}

	// Populate SPK with sections and details from SPH
	for _, grpcSection := range sphDetailsResponse.Sections {
		section := models.SPKSection{
			SphSectionId:   int(grpcSection.SphSectionId),
			SectionTitle:   &grpcSection.SectionTitle,
			CreatedBy:      spk.CreatedBy,
			OrganizationId: spk.OrganizationId,
		}

		for _, grpcDetail := range grpcSection.Details {
			detail := models.SPKDetail{
				SphItemId: func(v int32) *int {
					value := int(v)
					return &value
				}(grpcDetail.SphDetailId),
				Description:    &grpcDetail.ItemDescription,
				Quantity:       grpcDetail.Quantity,
				Unit:           &grpcDetail.Unit,
				UnitPriceJasa:  0,
				TotalJasa:      0,
				CreatedBy:      spk.CreatedBy,
				OrganizationId: spk.OrganizationId,
			}
			section.Details = append(section.Details, detail)
		}
		spk.Sections = append(spk.Sections, section)
	}

	if err := s.spkRepo.Create(spk); err != nil {
		log.Printf("Error creating SPK: %v", err)
		return fmt.Errorf("error creating SPK: %w", err)
	}

	// Publish event for centralized logging
	event := map[string]interface{}{
		"event":   "Created",
		"payload": spk,
		"meta": map[string]interface{}{
			"traceId": traceID,
			"action":  "CREATE",
			"idField": "SpkId",
			"userId":  spk.CreatedBy,
		},
	}
	if err := s.rabbitPublisher.PublishEvent(s.spkQueueName, event); err != nil {
		log.Printf("TraceID: %s - Error publishing create event for SPK ID: %d, %v", traceID, spk.SpkId, err)
	}

	log.Printf("Successfully created SPK: %+v", spk)
	return nil
}

// UpdateSpk updates an SPK
func (s *SpkService) UpdateSpk(spk *models.SPK, traceID string) error {

	if err := s.spkRepo.Update(spk); err != nil {
		log.Printf("TraceID: %s - Error updating SPK: %v", traceID, err)
		return fmt.Errorf("error updating SPK: %w", err)
	}

	// Publish event for centralized logging
	event := map[string]interface{}{
		"event":   "Updated",
		"payload": spk,
		"meta": map[string]interface{}{
			"traceId": traceID,
			"action":  "UPDATE",
			"idField": "SpkId",
			"userId":  spk.UpdatedBy,
		},
	}
	if err := s.rabbitPublisher.PublishEvent(s.spkQueueName, event); err != nil {
		log.Printf("TraceID: %s - Error publishing update event for SPK ID: %d, %v", traceID, spk.SpkId, err)
	}

	return s.PublishFullReindexEvent(spk.SpkId, traceID, *spk.UpdatedBy)
}

// DeleteSpk deletes an SPK
func (s *SpkService) DeleteSpk(spkID int, traceID string, userID int) error {
	if err := s.spkRepo.Delete(spkID); err != nil {
		log.Printf("TraceID: %s - Error deleting SPK: %v", traceID, err)
		return fmt.Errorf("error deleting SPK: %w", err)
	}

	// Publish event for centralized logging
	event := map[string]interface{}{
		"event":   "Deleted",
		"payload": map[string]int{"SpkId": spkID},
		"meta": map[string]interface{}{
			"traceId": traceID,
			"action":  "DELETE",
			"idField": "SpkId",
			"userId":  userID,
		},
	}
	if err := s.rabbitPublisher.PublishEvent(s.spkQueueName, event); err != nil {
		log.Printf("TraceID: %s - Error publishing delete event for SPK ID: %d, %v", traceID, spkID, err)
	}

	log.Printf("TraceID: %s - Successfully deleted SPK ID: %d", traceID, spkID)
	return nil
}

// CRUD Operations for SPK Sections

func (s *SpkService) CreateSpkSection(section *models.SPKSection, spkID int) error {
	log.Printf("Creating SPK Section for SPK ID: %d", spkID)

	// Assign the SPK ID
	section.SpkId = spkID

	// Save the section to the repository
	if err := s.spkSectionRepo.Create(section); err != nil {
		log.Printf("Error creating SPK Section: %v", err)
		return fmt.Errorf("failed to create SPK Section: %w", err)
	}

	s.PublishFullReindexEvent(spkID, "", 0)

	log.Printf("Successfully created SPK Section with ID: %d", section.SectionId)
	return nil
}

func (s *SpkService) UpdateSpkSection(updatedSection *models.SPKSection) error {
	log.Printf("Updating SPK Section ID: %d", updatedSection.SectionId)

	// Update the section in the repository
	if err := s.spkSectionRepo.Update(updatedSection); err != nil {
		log.Printf("Error updating SPK Section: %v", err)
		return fmt.Errorf("failed to update SPK Section: %w", err)
	}

	s.PublishFullReindexEvent(updatedSection.SpkId, "", *updatedSection.UpdatedBy)

	log.Printf("Successfully updated SPK Section with ID: %d", updatedSection.SectionId)
	return nil
}

func (s *SpkService) DeleteSpkSection(sectionID, SPKId int) error {
	log.Printf("Deleting SPK Section with ID: %d", sectionID)

	// Delete the section from the repository
	if err := s.spkSectionRepo.Delete(sectionID); err != nil {
		log.Printf("Error deleting SPK Section: %v", err)
		return fmt.Errorf("failed to delete SPK Section: %w", err)
	}

	s.PublishFullReindexEvent(SPKId, "", 0)

	log.Printf("Successfully deleted SPK Section with ID: %d", sectionID)
	return nil
}

func (s *SpkService) GetSpkSectionByID(sectionID int) (*models.SPKSection, error) {
	log.Printf("Fetching SPK Section with ID: %d", sectionID)

	// Retrieve the section from the repository
	section, err := s.spkSectionRepo.GetByID(sectionID)
	if err != nil {
		log.Printf("Error fetching SPK Section: %v", err)
		return nil, fmt.Errorf("failed to fetch SPK Section: %w", err)
	}

	return section, nil
}

// CRUD Operations for SPK Details

func (s *SpkService) CreateSpkDetail(detail *models.SPKDetail, sectionID, SPKId int) error {
	log.Printf("Creating SPK Detail for Section ID: %d", sectionID)

	// Assign the Section ID
	detail.SectionId = sectionID

	// Save the detail to the repository
	if err := s.detailRepo.Create(detail); err != nil {
		log.Printf("Error creating SPK Detail: %v", err)
		return fmt.Errorf("failed to create SPK Detail: %w", err)
	}

	s.PublishFullReindexEvent(SPKId, "", *detail.CreatedBy)

	log.Printf("Successfully created SPK Detail with ID: %d", detail.DetailId)
	return nil
}

func (s *SpkService) UpdateSpkDetail(updatedDetail *models.SPKDetail, SPKId int) error {
	log.Printf("Updating SPK Detail ID: %d", updatedDetail.DetailId)

	// Update the detail in the repository
	if err := s.detailRepo.Update(updatedDetail); err != nil {
		log.Printf("Error updating SPK Detail: %v", err)
		return fmt.Errorf("failed to update SPK Detail: %w", err)
	}

	s.PublishFullReindexEvent(SPKId, "", *updatedDetail.UpdatedBy)

	log.Printf("Successfully updated SPK Detail with ID: %d", updatedDetail.DetailId)
	return nil
}

func (s *SpkService) DeleteSpkDetail(detailID, SPKId int) error {
	log.Printf("Deleting SPK Detail with ID: %d", detailID)

	// Delete the detail from the repository
	if err := s.detailRepo.Delete(detailID); err != nil {
		log.Printf("Error deleting SPK Detail: %v", err)
		return fmt.Errorf("failed to delete SPK Detail: %w", err)
	}

	s.PublishFullReindexEvent(SPKId, "", detailID)

	log.Printf("Successfully deleted SPK Detail with ID: %d", detailID)
	return nil
}

func (s *SpkService) GetSpkDetailByID(detailID int) (*models.SPKDetail, error) {
	log.Printf("Fetching SPK Detail with ID: %d", detailID)

	// Retrieve the detail from the repository
	detail, err := s.detailRepo.GetByID(detailID)
	if err != nil {
		log.Printf("Error fetching SPK Detail: %v", err)
		return nil, fmt.Errorf("failed to fetch SPK Detail: %w", err)
	}

	return detail, nil
}

func (s *SpkService) ImportSpkFromExcel(metadata dto.SpkImportDTO, fileBytes []byte, organizationID int, userID int) (models.SPK, error) {
	// Load Excel from bytes
	f, err := excelize.OpenReader(bytes.NewReader(fileBytes))
	if err != nil {
		return models.SPK{}, fmt.Errorf("failed to open Excel file: %v", err)
	}

	// Get all rows from the first sheet
	rows, err := f.GetRows(f.GetSheetName(0))
	if err != nil {
		return models.SPK{}, fmt.Errorf("failed to read Excel rows: %v", err)
	}

	currentDate := time.Now().Format("2006-01-02")           // Get current date in yyyy-MM-dd format
	parsedDate, err := time.Parse("2006-01-02", currentDate) // Parse it into a time.Time object
	if err != nil {
		return models.SPK{}, fmt.Errorf("failed to parse date: %v", err)
	}

	// Create SPK record
	spk := models.SPK{
		ProjectId:      metadata.ProjectId,
		Subject:        metadata.Subject,
		Date:           &models.CustomDate{Time: parsedDate}, // Set the parsed date
		Mandor:         metadata.Mandor,
		OrganizationId: &organizationID,
		CreatedBy:      &userID,
	}

	if err := s.spkRepo.Create(&spk); err != nil {
		return spk, fmt.Errorf("failed to create SPK: %v", err)
	}

	var currentSectionID int
	headerFound := false           // Flag to indicate whether the header row has been found
	var grandTotal float64         // Variable to accumulate the grand total
	var grandTotalMaterial float64 // Variable to accumulate the grand total

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
			section := models.SPKSection{
				SpkId:          spk.SpkId,
				SectionTitle:   &row[1],
				OrganizationId: &organizationID,
				CreatedBy:      &userID,
			}
			if err := s.spkSectionRepo.Create(&section); err != nil {
				return spk, fmt.Errorf("failed to create SPK section: %v", err)
			}
			currentSectionID = section.SectionId // Get the inserted SectionId
			continue
		}

		// Parse rows as details
		if isNumeric(row[0]) {
			if len(row) < 7 { // Ensure detail row has enough columns
				log.Printf("Skipping invalid detail row at line %d: insufficient columns", rowIndex+1)
				continue
			}

			// Parse TotalPrice
			totalPrice := parseFloatNonPointer(row[5])
			if totalPrice != 0 {
				grandTotal += totalPrice // Add to grand total
			}

			totalPriceMaterial := parseFloatNonPointer(row[8])
			if totalPriceMaterial != 0 {
				grandTotalMaterial += totalPriceMaterial // Add to grand total
			}

			detail := models.SPKDetail{
				SectionId:         currentSectionID, // Use the actual SectionId
				Description:       &row[1],
				Quantity:          parseFloatNonPointer(row[2]),
				Unit:              &row[3],
				UnitPriceJasa:     parseFloatNonPointer(row[4]),
				TotalJasa:         totalPrice,
				OrganizationId:    &organizationID,
				UnitPriceMaterial: parseFloatNonPointer(row[7]),
				TotalMaterial:     totalPriceMaterial,
				CreatedBy:         &userID,
			}
			if err := s.detailRepo.Create(&detail); err != nil {
				return spk, fmt.Errorf("failed to create SPH detail: %v", err)
			}
		}
	}

	spk.TotalJasa = &grandTotal
	spk.TotalMaterial = &grandTotalMaterial

	if err := s.spkRepo.Update(&spk); err != nil {
		log.Printf("Error updating SPK: %v", err)
	}
	// Successfully processed the Excel and updated the project
	s.PublishFullReindexEvent(spk.SpkId, "", userID)

	// Return the grand total along with nil error (indicating success)
	return spk, nil
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

// Helper function to parse a string into a float64 value
func parseFloatNonPointer(input string) float64 {
	// Trim spaces and unwanted characters
	cleaned := strings.TrimSpace(strings.ReplaceAll(input, ",", ""))

	// Attempt to parse the cleaned string
	value, err := strconv.ParseFloat(cleaned, 64)
	if err != nil {
		return 0 // Return 0 if parsing fails (or choose another default value)
	}
	return value
}
