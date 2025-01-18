package services

import (
	"AmanahPro/services/ba-services/internal/domain/models"
	"AmanahPro/services/ba-services/internal/domain/repositories"
	"AmanahPro/services/ba-services/internal/dto"
	"bytes"
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/NHadi/AmanahPro-common/messagebroker"
	pb "github.com/NHadi/AmanahPro-common/protos"
	"github.com/xuri/excelize/v2"
)

type BAService struct {
	baRepo          repositories.BARepository
	sectionRepo     repositories.BASectionRepository
	detailRepo      repositories.BADetailRepository
	progressRepo    repositories.BAProgressRepository
	rabbitPublisher *messagebroker.RabbitMQPublisher
	baQueueName     string
	sphGrpcClient   pb.SphServiceClient // gRPC client for SPH
}

// NewBAService initializes the BAService
func NewBAService(
	baRepo repositories.BARepository,
	sectionRepo repositories.BASectionRepository,
	detailRepo repositories.BADetailRepository,
	progressRepo repositories.BAProgressRepository,
	rabbitPublisher *messagebroker.RabbitMQPublisher,
	baQueueName string,
	sphGrpcClient pb.SphServiceClient, // Inject SPH gRPC client
) *BAService {
	return &BAService{
		baRepo:          baRepo,
		sectionRepo:     sectionRepo,
		detailRepo:      detailRepo,
		progressRepo:    progressRepo,
		rabbitPublisher: rabbitPublisher,
		baQueueName:     baQueueName,
		sphGrpcClient:   sphGrpcClient,
	}
}

// CRUD Operations for BA

// FilterSPKs filters SPKs by organization ID and optional SPK ID or project ID
func (s *BAService) Filter(organizationID int, baID *int, projectID *int) ([]models.BA, error) {
	log.Printf("Filtering BAs for OrganizationID: %d, BAID: %v, ProjectID: %v", organizationID, baID, projectID)

	BAs, err := s.baRepo.Filter(organizationID, baID, projectID)
	if err != nil {
		log.Printf("Error filtering BAs: %v", err)
		return nil, fmt.Errorf("failed to filter SPKs: %w", err)
	}

	log.Printf("Found %d BAs for OrganizationID: %d", len(BAs), organizationID)
	return BAs, nil
}
func (s *BAService) GetBAByID(baID int) (*models.BA, error) {
	log.Printf("Fetching BA with ID: %d", baID)
	ba, err := s.baRepo.GetByID(baID, true)
	if err != nil {
		log.Printf("Error fetching BA: %v", err)
		return nil, fmt.Errorf("failed to fetch BA: %w", err)
	}
	return ba, nil
}

func (s *BAService) CreateBA(ba *models.BA, sphId int32) error {
	log.Printf("Creating BA: %+v with BA ID: %d", ba, sphId)

	// Call SPH gRPC service to get sections and details
	log.Printf("Calling SPH gRPC service for SPH ID: %d", sphId)
	sphDetailsResponse, err := s.sphGrpcClient.GetSphDetails(context.Background(), &pb.GetSphDetailsRequest{SphId: sphId})
	if err != nil {
		log.Printf("Failed to fetch SPH details from gRPC: %v", err)
		return fmt.Errorf("failed to fetch SPH details from gRPC: %w", err)
	}

	// Calculate the total sum of TotalPrice for all details
	var totalPriceSum float64
	for _, grpcSection := range sphDetailsResponse.Sections {
		for _, grpcDetail := range grpcSection.Details {
			totalPriceSum += grpcDetail.TotalPrice
		}
	}

	// Populate BA with sections and details from SPH
	for _, grpcSection := range sphDetailsResponse.Sections {
		section := models.BASection{
			SphSectionId:   int(grpcSection.SphSectionId),
			SectionName:    &grpcSection.SectionTitle,
			CreatedBy:      ba.CreatedBy,
			OrganizationId: ba.OrganizationId,
		}

		for _, grpcDetail := range grpcSection.Details {
			// Calculate WeightPercentage for the detail
			weightPercentage := 0.0
			if totalPriceSum > 0 {
				weightPercentage = (grpcDetail.TotalPrice / totalPriceSum) * 100
			}

			detail := models.BADetail{
				SphItemId: func(v int32) *int {
					value := int(v)
					return &value
				}(grpcDetail.SphDetailId),
				ItemName:         &grpcDetail.ItemDescription,
				Quantity:         grpcDetail.Quantity,
				Unit:             &grpcDetail.Unit,
				UnitPrice:        &grpcDetail.UnitPrice,
				DiscountPrice:    &grpcDetail.DiscountPrice,
				TotalPrice:       &grpcDetail.TotalPrice,
				WeightPercentage: &weightPercentage, // Assign calculated WeightPercentage
				CreatedBy:        ba.CreatedBy,
				OrganizationId:   ba.OrganizationId,
			}

			detailProgress := models.BAProgress{
				ProgressPreviousM2:         nil,
				ProgressPreviousPercentage: nil,
				ProgressCurrentM2:          nil,
				ProgressCurrentPercentage:  nil,
				CreatedBy:                  ba.CreatedBy,
				OrganizationId:             ba.OrganizationId,
			}

			detail.Progress = append(detail.Progress, detailProgress)
			section.Details = append(section.Details, detail)
		}
		ba.Sections = append(ba.Sections, section)
	}

	if err := s.baRepo.Create(ba); err != nil {
		log.Printf("Error creating BA: %v", err)
		return fmt.Errorf("error creating BA: %w", err)
	}

	// Publish event for centralized logging
	event := map[string]interface{}{
		"event":   "Created",
		"payload": ba,
		"meta": map[string]interface{}{
			"traceId": "",
			"action":  "CREATE",
			"idField": "BAId",
			"userId":  ba.CreatedBy,
		},
	}
	if err := s.rabbitPublisher.PublishEvent(s.baQueueName, event); err != nil {
		log.Printf("TraceID: %s - Error publishing create event for BA ID: %d, %v", "", ba.BAId, err)
	}

	log.Printf("Successfully created BA: %+v", ba)
	return nil
}

func (s *BAService) UpdateBA(ba *models.BA) error {
	log.Printf("Updating BA: %+v", ba)
	if err := s.baRepo.Update(ba); err != nil {
		log.Printf("Error updating BA: %v", err)
		return fmt.Errorf("error updating BA: %w", err)
	}
	s.PublishFullReindexEvent(ba.BAId, "", *ba.UpdatedBy)
	return nil
}

func (s *BAService) DeleteBA(baID int, userID int) error {
	log.Printf("Deleting BA ID: %d", baID)
	if err := s.baRepo.Delete(baID); err != nil {
		log.Printf("Error deleting BA: %v", err)
		return fmt.Errorf("error deleting BA: %w", err)
	}
	// Publish event for centralized logging
	event := map[string]interface{}{
		"event":   "Deleted",
		"payload": map[string]int{"BAId": baID},
		"meta": map[string]interface{}{
			"traceId": "",
			"action":  "DELETE",
			"idField": "BAId",
			"userId":  userID,
		},
	}
	if err := s.rabbitPublisher.PublishEvent(s.baQueueName, event); err != nil {
		log.Printf("TraceID: %s - Error publishing delete event for BA ID: %d, %v", "", baID, err)
	}
	return nil
}

// AddBAProgress adds default progress entries (all fields set to 0) for all details under a specific BA.
func (s *BAService) AddBAProgress(baID int, createdBy *int) error {
	log.Printf("Adding default progress to BA with ID: %d", baID)

	// Fetch the BA along with its hierarchy
	ba, err := s.GetBAByID(baID)
	if err != nil {
		log.Printf("Error fetching BA: %v", err)
		return fmt.Errorf("failed to fetch BA: %w", err)
	}

	if ba == nil {
		log.Printf("BA with ID %d not found", baID)
		return fmt.Errorf("BA not found")
	}

	// Iterate through sections and details to add progress
	for _, section := range ba.Sections {
		for _, detail := range section.Details {
			// Create a default progress entry for the detail
			progress := &models.BAProgress{
				DetailId:                   detail.DetailId,
				ProgressCurrentM2:          nil,
				ProgressCurrentPercentage:  nil,
				ProgressPreviousM2:         nil,
				ProgressPreviousPercentage: nil,
				CreatedBy:                  createdBy,
				OrganizationId:             detail.OrganizationId,
			}

			// Insert the progress into the database
			if err := s.progressRepo.Create(progress); err != nil {
				log.Printf("Error adding progress to detail ID %d: %v", detail.DetailId, err)
				return fmt.Errorf("failed to add progress to detail ID %d: %w", detail.DetailId, err)
			}
		}
	}
	s.PublishFullReindexEvent(baID, "", *createdBy)
	log.Printf("Default progress successfully added to BA with ID: %d", baID)
	return nil
}

// CRUD Operations for BA Section
func (s *BAService) GetBASectionByID(baSectionId int) (*models.BASection, error) {
	log.Printf("Fetching BASection with ID: %d", baSectionId)
	ba, err := s.sectionRepo.GetByID(baSectionId)
	if err != nil {
		log.Printf("Error fetching BASection: %v", err)
		return nil, fmt.Errorf("failed to fetch BASection: %w", err)
	}
	return ba, nil
}
func (s *BAService) CreateSection(section *models.BASection) error {
	log.Printf("Creating BA Section: %+v", section)
	if err := s.sectionRepo.Create(section); err != nil {
		log.Printf("Error creating BA Section: %v", err)
		return fmt.Errorf("error creating BA Section: %w", err)
	}
	s.PublishFullReindexEvent(*section.BAID, "", *section.CreatedBy)
	return nil
}

func (s *BAService) UpdateSection(section *models.BASection) error {
	log.Printf("Updating BA Section ID: %d", section.BASectionId)
	if err := s.sectionRepo.Update(section); err != nil {
		log.Printf("Error updating BA Section: %v", err)
		return fmt.Errorf("error updating BA Section: %w", err)
	}
	s.PublishFullReindexEvent(*section.BAID, "", *section.UpdatedBy)
	return nil
}

func (s *BAService) DeleteSection(sectionID, baID, userID int) error {
	log.Printf("Deleting BA Section ID: %d", sectionID)
	if err := s.sectionRepo.Delete(sectionID); err != nil {
		log.Printf("Error deleting BA Section: %v", err)
		return fmt.Errorf("error deleting BA Section: %w", err)
	}
	s.PublishFullReindexEvent(baID, "", userID)
	return nil
}

// CRUD Operations for BA Detail

func (s *BAService) GetBADetailByID(baDetailId int) (*models.BADetail, error) {
	log.Printf("Fetching baDetailId with ID: %d", baDetailId)
	ba, err := s.detailRepo.GetByID(baDetailId)
	if err != nil {
		log.Printf("Error fetching baDetailId: %v", err)
		return nil, fmt.Errorf("failed to fetch BASection: %w", err)
	}
	return ba, nil
}
func (s *BAService) CreateDetail(detail *models.BADetail, baId int) error {
	log.Printf("Creating BA Detail: %+v", detail)
	if err := s.detailRepo.Create(detail); err != nil {
		log.Printf("Error creating BA Detail: %v", err)
		return fmt.Errorf("error creating BA Detail: %w", err)
	}
	s.PublishFullReindexEvent(baId, "", *detail.CreatedBy)
	return nil
}

func (s *BAService) UpdateDetail(detail *models.BADetail, baId int) error {
	log.Printf("Updating BA Detail ID: %d", detail.DetailId)
	if err := s.detailRepo.Update(detail); err != nil {
		log.Printf("Error updating BA Detail: %v", err)
		return fmt.Errorf("error updating BA Detail: %w", err)
	}
	s.PublishFullReindexEvent(baId, "", *detail.CreatedBy)
	return nil
}

func (s *BAService) DeleteDetail(detailID, baId, userId int) error {
	log.Printf("Deleting BA Detail ID: %d", detailID)
	if err := s.detailRepo.Delete(detailID); err != nil {
		log.Printf("Error deleting BA Detail: %v", err)
		return fmt.Errorf("error deleting BA Detail: %w", err)
	}
	s.PublishFullReindexEvent(baId, "", userId)
	return nil
}

// CRUD Operations for BA Progress
func (s *BAService) GetBAProgressByID(baProgressId int) (*models.BAProgress, error) {
	log.Printf("Fetching baProgressId with ID: %d", baProgressId)
	ba, err := s.progressRepo.GetByID(baProgressId)
	if err != nil {
		log.Printf("Error fetching baProgressId: %v", err)
		return nil, fmt.Errorf("failed to fetch BAProgress: %w", err)
	}
	return ba, nil
}

func (s *BAService) CreateProgress(progress *models.BAProgress, baId int) error {
	log.Printf("Creating BA Progress: %+v", progress)
	if err := s.progressRepo.Create(progress); err != nil {
		log.Printf("Error creating BA Progress: %v", err)
		return fmt.Errorf("error creating BA Progress: %w", err)
	}
	s.PublishFullReindexEvent(baId, "", *progress.CreatedBy)
	return nil
}

func (s *BAService) UpdateProgress(progress *models.BAProgress, baId int) error {
	log.Printf("Updating BA Progress ID: %d", progress.BAProgressId)
	if err := s.progressRepo.Update(progress); err != nil {
		log.Printf("Error updating BA Progress: %v", err)
		return fmt.Errorf("error updating BA Progress: %w", err)
	}
	s.PublishFullReindexEvent(baId, "", *progress.UpdatedBy)
	return nil
}

func (s *BAService) DeleteProgress(progressID, baId, userId int) error {
	log.Printf("Deleting BA Progress ID: %d", progressID)
	if err := s.progressRepo.Delete(progressID); err != nil {
		log.Printf("Error deleting BA Progress: %v", err)
		return fmt.Errorf("error deleting BA Progress: %w", err)
	}
	s.PublishFullReindexEvent(baId, "", userId)
	return nil
}

// Helper: PublishFullReindexEvent sends a RabbitMQ event to re-index the full SPK
func (s *BAService) PublishFullReindexEvent(BAId int, traceID string, userID int) error {
	log.Printf("Triggering re-index for BAId: %d", BAId)

	// Retrieve the full SPK structure for re-indexing
	ba, err := s.baRepo.GetByID(BAId, true)
	if err != nil {
		log.Printf("Error retrieving BAId for reindexing: %v", err)
		return fmt.Errorf("error retrieving BA for reindexing: %w", err)
	}

	if ba == nil {
		log.Printf("Reindex failed because BA ID: %d Not Found", BAId)
		return nil
	}

	event := map[string]interface{}{
		"event":   "Reindexed",
		"payload": ba,
		"meta": map[string]interface{}{
			"traceId": traceID,
			"action":  "REINDEX",
			"idField": "BAId",
			"userId":  userID,
		},
	}
	if err := s.rabbitPublisher.PublishEvent(s.baQueueName, event); err != nil {
		log.Printf("Error publishing re-index event for BAID: %d, %v", BAId, err)
		return fmt.Errorf("error publishing re-index event: %w", err)
	}

	log.Printf("Successfully triggered re-index for BAId: %d", BAId)
	return nil
}

func (s *BAService) ImportFromExcel(metadata dto.BAImportDTO, fileBytes []byte, organizationID int, userID int) (models.BA, error) {
	// Load Excel from bytes
	f, err := excelize.OpenReader(bytes.NewReader(fileBytes))
	if err != nil {
		return models.BA{}, fmt.Errorf("failed to open Excel file: %v", err)
	}

	// Get all rows from the first sheet
	rows, err := f.GetRows(f.GetSheetName(0))
	if err != nil {
		return models.BA{}, fmt.Errorf("failed to read Excel rows: %v", err)
	}

	currentDate := time.Now().Format("2006-01-02")           // Get current date in yyyy-MM-dd format
	parsedDate, err := time.Parse("2006-01-02", currentDate) // Parse it into a time.Time object
	if err != nil {
		return models.BA{}, fmt.Errorf("failed to parse date: %v", err)
	}

	// Create BA record
	ba := models.BA{
		ProjectId:      metadata.ProjectId,
		BASubject:      *metadata.Subject,
		BADate:         models.CustomDate{Time: parsedDate}, // Set the parsed date
		RecepientName:  metadata.RecepientName,
		OrganizationId: &organizationID,
		CreatedBy:      &userID,
	}

	if err := s.baRepo.Create(&ba); err != nil {
		return ba, fmt.Errorf("failed to create BA: %v", err)
	}

	var currentSectionID int
	var currentDetailID int
	headerFound := false                      // Flag to indicate whether the header row has been found
	var grandTotal float64                    // Variable to accumulate the grand total
	var totalCurrentProgresPercentage float64 // Variable to accumulate the grand total
	var totalCurrentProgresM2 float64         // Variable to accumulate the grand total
	var totalPrevProgresPercentage float64    // Variable to accumulate the grand total
	var totalPrevProgresM2 float64            // Variable to accumulate the grand total

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
			section := models.BASection{
				BAID:           &ba.BAId,
				SectionName:    &row[1],
				OrganizationId: &organizationID,
				CreatedBy:      &userID,
			}
			if err := s.sectionRepo.Create(&section); err != nil {
				return ba, fmt.Errorf("failed to create BA section: %v", err)
			}
			currentSectionID = section.BASectionId // Get the inserted SectionId
			continue
		}

		// Parse rows as details
		if isNumeric(row[0]) {
			if len(row) < 7 { // Ensure detail row has enough columns
				log.Printf("Skipping invalid detail row at line %d: insufficient columns", rowIndex+1)
				continue
			}

			unitPrice := parseFloat(row[4])
			discountedPrice := parseFloat(row[5])
			discountPercentage := 0.0
			totalPrice := parseFloat(row[6])

			// Add to grand totals
			if totalPrice != nil {
				grandTotal += *totalPrice
			}

			totalPrevProgresM2 += parseFloatNonPointer(row[8])
			totalPrevProgresPercentage += parseFloatNonPointer(row[9])
			totalCurrentProgresM2 += parseFloatNonPointer(row[10])
			totalCurrentProgresPercentage += parseFloatNonPointer(row[11])

			if unitPrice != nil && discountedPrice != nil {
				discountPercentage = ((*unitPrice - *discountedPrice) / *unitPrice) * 100
			}

			detail := models.BADetail{
				SectionID:        &currentSectionID,
				ItemName:         &row[1],
				Quantity:         parseFloatNonPointer(row[2]),
				Unit:             &row[3],
				UnitPrice:        unitPrice,
				DiscountPrice:    &discountPercentage,
				TotalPrice:       totalPrice,
				OrganizationId:   &organizationID,
				WeightPercentage: parseFloat(row[7]),
				CreatedBy:        &userID,
			}

			if err := s.detailRepo.Create(&detail); err != nil {
				return ba, fmt.Errorf("failed to create BADetail: %v", err)
			}

			log.Printf("Successfully created BADetail with ID: %v", detail.DetailId)

			currentDetailID = detail.DetailId
			if currentDetailID == 0 {
				return ba, fmt.Errorf("BADetail creation failed: DetailId is 0")
			}

			progress := models.BAProgress{
				DetailId:                   currentDetailID,
				ProgressPreviousM2:         parseFloat(row[8]),
				ProgressPreviousPercentage: parseFloat(row[9]),
				ProgressCurrentM2:          parseFloat(row[10]),
				ProgressCurrentPercentage:  parseFloat(row[11]),
				OrganizationId:             &organizationID,
				CreatedBy:                  &userID,
			}

			if err := s.progressRepo.Create(&progress); err != nil {
				return ba, fmt.Errorf("failed to create BAProgress: %v", err)
			}

			log.Printf("Successfully created BAProgress for DetailId: %v", currentDetailID)
		}

	}

	ba.TotalPrice = &grandTotal
	ba.ProgressPreviousM2 = &totalPrevProgresM2
	ba.ProgressPreviousPercentage = &totalPrevProgresPercentage
	ba.ProgressCurrentM2 = &totalCurrentProgresM2
	ba.ProgressCurrentPercentage = &totalCurrentProgresPercentage

	if err := s.baRepo.Update(&ba); err != nil {
		log.Printf("Error updating BA: %v", err)
	}
	// Successfully processed the Excel and updated the project
	s.PublishFullReindexEvent(ba.BAId, "", userID)

	// Return the grand total along with nil error (indicating success)
	return ba, nil
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
