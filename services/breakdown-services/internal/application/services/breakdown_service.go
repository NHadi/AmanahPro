package services

import (
	"AmanahPro/services/breakdown-services/common/messagebroker"
	"AmanahPro/services/breakdown-services/internal/domain/models"
	"AmanahPro/services/breakdown-services/internal/domain/repositories"
	"AmanahPro/services/breakdown-services/internal/dto"
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/xuri/excelize/v2"
)

type BreakdownService struct {
	breakdownRepo      repositories.BreakdownRepository
	sectionRepo        repositories.BreakdownSectionRepository
	itemRepo           repositories.BreakdownItemRepository
	mstSectionRepo     repositories.MstBreakdownSectionRepository
	mstItemRepo        repositories.MstBreakdownItemRepository
	rabbitPublisher    *messagebroker.RabbitMQPublisher
	breakdownQueueName string
}

// NewBreakdownService initializes the BreakdownService
func NewBreakdownService(
	breakdownRepo repositories.BreakdownRepository,
	sectionRepo repositories.BreakdownSectionRepository,
	itemRepo repositories.BreakdownItemRepository,
	mstSectionRepo repositories.MstBreakdownSectionRepository,
	mstItemRepo repositories.MstBreakdownItemRepository,
	rabbitPublisher *messagebroker.RabbitMQPublisher,
	breakdownQueueName string,
) *BreakdownService {
	return &BreakdownService{
		breakdownRepo:      breakdownRepo,
		sectionRepo:        sectionRepo,
		itemRepo:           itemRepo,
		mstSectionRepo:     mstSectionRepo,
		mstItemRepo:        mstItemRepo,
		rabbitPublisher:    rabbitPublisher,
		breakdownQueueName: breakdownQueueName,
	}
}

// GetBreakdownByID retrieves a breakdown by its ID
func (s *BreakdownService) GetBreakdownByID(id int) (*models.Breakdown, error) {
	item, err := s.breakdownRepo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("error retrieving breakdown item: %w", err)
	}
	return item, nil
}

// FilterBreakdowns filters breakdowns by organization ID and optional breakdown ID or project ID
func (s *BreakdownService) FilterBreakdowns(organizationID int, breakdownID *int, projectID *int) ([]models.Breakdown, error) {
	log.Printf("Filtering breakdowns for OrganizationID: %d, BreakdownID: %v, ProjectID: %v", organizationID, breakdownID, projectID)

	breakdowns, err := s.breakdownRepo.FilterBreakdowns(organizationID, breakdownID, projectID)
	if err != nil {
		log.Printf("Error filtering breakdowns: %v", err)
		return nil, fmt.Errorf("failed to filter breakdowns: %w", err)
	}

	log.Printf("Found %d breakdowns for OrganizationID: %d", len(breakdowns), organizationID)
	return breakdowns, nil
}

// Helper: PublishFullReindexEvent sends a RabbitMQ event to re-index the full Breakdown
func (s *BreakdownService) PublishFullReindexEvent(breakdownID int) error {
	log.Printf("Triggering re-index for BreakdownID: %d", breakdownID)

	// Retrieve the full breakdown structure for re-indexing
	breakdown, err := s.breakdownRepo.GetByID(breakdownID)
	if err != nil {
		log.Printf("Error retrieving breakdown for reindexing: %v", err)
		return fmt.Errorf("error retrieving breakdown for reindexing: %w", err)
	}

	event := map[string]interface{}{
		"event":   "Reindexed",
		"payload": breakdown,
		"meta": map[string]interface{}{
			"idField": "BreakdownId",
		},
	}
	if err := s.rabbitPublisher.PublishEvent(s.breakdownQueueName, event); err != nil {
		log.Printf("Error publishing re-index event for BreakdownID: %d, %v", breakdownID, err)
		return nil
	}

	log.Printf("Successfully triggered re-index for BreakdownID: %d", breakdownID)
	return nil
}

// Breakdown CRUD
func (s *BreakdownService) CreateBreakdown(breakdown *models.Breakdown) error {
	log.Printf("Creating Breakdown: %+v", breakdown)
	if err := s.breakdownRepo.Create(breakdown); err != nil {
		log.Printf("Error creating Breakdown: %v", err)
		return fmt.Errorf("error creating Breakdown: %w", err)
	}

	event := map[string]interface{}{
		"event":   "Created",
		"payload": breakdown,
		"meta": map[string]interface{}{
			"idField": "BreakdownId",
		},
	}
	if err := s.rabbitPublisher.PublishEvent(s.breakdownQueueName, event); err != nil {
		log.Printf("Error publishing index event for BreakdownID: %d, %v", breakdown.BreakdownId, err)
		return nil
	}

	log.Printf("Successfully triggered index for BreakdownID: %d", breakdown.BreakdownId)
	return nil
}

func (s *BreakdownService) UpdateBreakdown(breakdown *models.Breakdown) error {
	log.Printf("Updating Breakdown: %+v", breakdown)
	if err := s.breakdownRepo.Update(breakdown); err != nil {
		log.Printf("Error updating Breakdown: %v", err)
		return fmt.Errorf("error updating Breakdown: %w", err)
	}

	return s.PublishFullReindexEvent(breakdown.BreakdownId)
}

func (s *BreakdownService) DeleteBreakdown(breakdownID int) error {
	log.Printf("Deleting Breakdown ID: %d", breakdownID)
	if err := s.breakdownRepo.Delete(breakdownID); err != nil {
		log.Printf("Error deleting Breakdown: %v", err)
		return fmt.Errorf("error deleting Breakdown: %w", err)
	}

	// Send a delete event to RabbitMQ
	event := map[string]interface{}{
		"event":   "Deleted",
		"payload": map[string]int{"BreakdownId": breakdownID},
		"meta": map[string]interface{}{
			"idField": "BreakdownId",
		},
	}
	if err := s.rabbitPublisher.PublishEvent(s.breakdownQueueName, event); err != nil {
		log.Printf("Error publishing delete event for BreakdownID: %d, %v", breakdownID, err)
		return fmt.Errorf("error publishing delete event: %w", err)
	}

	log.Printf("Successfully deleted Breakdown ID: %d", breakdownID)
	return nil
}

// Section CRUD

func (s *BreakdownService) GetBreakdownSectionByID(sectionID int, breakdownID int) (*models.BreakdownSection, error) {
	log.Printf("Fetching BreakdownSection with ID: %d and BreakdownID: %d", sectionID, breakdownID)

	section, err := s.sectionRepo.GetByIDAndBreakdownID(sectionID, breakdownID)
	if err != nil {
		log.Printf("Error fetching BreakdownSection: %v", err)
		return nil, err
	}

	return section, nil
}

func (s *BreakdownService) CreateBreakdownSection(section *models.BreakdownSection) error {
	log.Printf("Creating BreakdownSection: %+v", section)
	if err := s.sectionRepo.Create(section); err != nil {
		log.Printf("Error creating BreakdownSection: %v", err)
		return fmt.Errorf("error creating BreakdownSection: %w", err)
	}

	return s.PublishFullReindexEvent(section.BreakdownId)
}

func (s *BreakdownService) UpdateBreakdownSection(section *models.BreakdownSection) error {
	log.Printf("Updating BreakdownSection: %+v", section)
	if err := s.sectionRepo.Update(section); err != nil {
		log.Printf("Error updating BreakdownSection: %v", err)
		return fmt.Errorf("error updating BreakdownSection: %w", err)
	}

	return s.PublishFullReindexEvent(section.BreakdownId)
}

func (s *BreakdownService) DeleteBreakdownSection(sectionID int, breakdownID int) error {
	log.Printf("Deleting BreakdownSection ID: %d", sectionID)
	if err := s.sectionRepo.Delete(sectionID); err != nil {
		log.Printf("Error deleting BreakdownSection: %v", err)
		return fmt.Errorf("error deleting BreakdownSection: %w", err)
	}

	return s.PublishFullReindexEvent(breakdownID)
}

// Item CRUD
// GetBreakdownItemByID retrieves a BreakdownItem by its ID
func (s *BreakdownService) GetBreakdownItemByID(id int) (*models.BreakdownItem, error) {
	item, err := s.itemRepo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("error retrieving breakdown item: %w", err)
	}
	return item, nil
}

func (s *BreakdownService) CreateBreakdownItem(item *models.BreakdownItem, breakdownId int) error {
	log.Printf("Creating BreakdownItem: %+v", item)
	if err := s.itemRepo.Create(item); err != nil {
		log.Printf("Error creating BreakdownItem: %v", err)
		return fmt.Errorf("error creating BreakdownItem: %w", err)
	}

	return s.PublishFullReindexEvent(breakdownId)
}

func (s *BreakdownService) UpdateBreakdownItem(item *models.BreakdownItem, breakdownId int) error {
	log.Printf("Updating BreakdownItem: %+v", item)
	if err := s.itemRepo.Update(item); err != nil {
		log.Printf("Error updating BreakdownItem: %v", err)
		return fmt.Errorf("error updating BreakdownItem: %w", err)
	}

	return s.PublishFullReindexEvent(breakdownId)
}

func (s *BreakdownService) DeleteBreakdownItem(itemID int, breakdownID int) error {
	log.Printf("Deleting BreakdownItem ID: %d", itemID)
	if err := s.itemRepo.Delete(itemID); err != nil {
		log.Printf("Error deleting BreakdownItem: %v", err)
		return fmt.Errorf("error deleting BreakdownItem: %w", err)
	}

	return s.PublishFullReindexEvent(breakdownID)
}

// MstBreakdownSection CRUD Operations

func (s *BreakdownService) GetMstBreakdownSectionyID(itemID int) (*models.MstBreakdownSection, error) {
	log.Printf("Fetching MstBreakdownItem with ID: %d", itemID)

	item, err := s.mstSectionRepo.GetByID(itemID)
	if err != nil {
		log.Printf("Error fetching MstBreakdownItem: %v", err)
		return nil, fmt.Errorf("error retrieving MstBreakdownItem: %w", err)
	}

	return item, nil
}

// FilterMstBreakdownSections retrieves all MstBreakdownSections for an organization
func (s *BreakdownService) FilterMstBreakdownSections(organizationID *int) ([]models.MstBreakdownSection, error) {
	log.Printf("Fetching MstBreakdownSections for OrganizationID: %v", organizationID)

	sections, err := s.mstSectionRepo.FilterBreakdowns(organizationID)
	if err != nil {
		log.Printf("Error fetching MstBreakdownSections: %v", err)
		return nil, fmt.Errorf("error retrieving MstBreakdownSections: %w", err)
	}

	log.Printf("Successfully fetched %d MstBreakdownSections", len(sections))
	return sections, nil
}

// CreateMstBreakdownSection creates a new MstBreakdownSection
func (s *BreakdownService) CreateMstBreakdownSection(section *models.MstBreakdownSection) error {
	log.Printf("Creating MstBreakdownSection: %+v", section)

	if err := s.mstSectionRepo.Create(section); err != nil {
		log.Printf("Error creating MstBreakdownSection: %v", err)
		return fmt.Errorf("error creating MstBreakdownSection: %w", err)
	}

	return nil
}

// UpdateMstBreakdownSection updates an existing MstBreakdownSection
func (s *BreakdownService) UpdateMstBreakdownSection(section *models.MstBreakdownSection) error {
	log.Printf("Updating MstBreakdownSection: %+v", section)

	if err := s.mstSectionRepo.Update(section); err != nil {
		log.Printf("Error updating MstBreakdownSection: %v", err)
		return fmt.Errorf("error updating MstBreakdownSection: %w", err)
	}

	return nil
}

// DeleteMstBreakdownSection deletes a MstBreakdownSection by its ID
func (s *BreakdownService) DeleteMstBreakdownSection(sectionID int) error {
	log.Printf("Deleting MstBreakdownSection ID: %d", sectionID)

	if err := s.mstSectionRepo.Delete(sectionID); err != nil {
		log.Printf("Error deleting MstBreakdownSection: %v", err)
		return fmt.Errorf("error deleting MstBreakdownSection: %w", err)
	}

	return nil
}

// MstBreakdownItem CRUD Operations

// GetMstBreakdownItemByID retrieves a MstBreakdownItem by its ID
func (s *BreakdownService) GetMstBreakdownItemByID(itemID int) (*models.MstBreakdownItem, error) {
	log.Printf("Fetching MstBreakdownItem with ID: %d", itemID)

	item, err := s.mstItemRepo.GetByID(itemID)
	if err != nil {
		log.Printf("Error fetching MstBreakdownItem: %v", err)
		return nil, fmt.Errorf("error retrieving MstBreakdownItem: %w", err)
	}

	return item, nil
}

// CreateMstBreakdownItem creates a new MstBreakdownItem
func (s *BreakdownService) CreateMstBreakdownItem(item *models.MstBreakdownItem) error {
	log.Printf("Creating MstBreakdownItem: %+v", item)

	if err := s.mstItemRepo.Create(item); err != nil {
		log.Printf("Error creating MstBreakdownItem: %v", err)
		return fmt.Errorf("error creating MstBreakdownItem: %w", err)
	}

	return nil
}

// UpdateMstBreakdownItem updates an existing MstBreakdownItem
func (s *BreakdownService) UpdateMstBreakdownItem(item *models.MstBreakdownItem) error {
	log.Printf("Updating MstBreakdownItem: %+v", item)

	if err := s.mstItemRepo.Update(item); err != nil {
		log.Printf("Error updating MstBreakdownItem: %v", err)
		return fmt.Errorf("error updating MstBreakdownItem: %w", err)
	}

	return nil
}

// DeleteMstBreakdownItem deletes a MstBreakdownItem by its ID
func (s *BreakdownService) DeleteMstBreakdownItem(itemID int) error {
	log.Printf("Deleting MstBreakdownItem ID: %d", itemID)

	if err := s.mstItemRepo.Delete(itemID); err != nil {
		log.Printf("Error deleting MstBreakdownItem: %v", err)
		return fmt.Errorf("error deleting MstBreakdownItem: %w", err)
	}

	return nil
}

// SyncBreakdown synchronizes BreakdownSection and BreakdownItem with master data.
func (s *BreakdownService) SyncBreakdown(breakdownId int, organizationId *int) error {
	// Fetch master sections and their items
	masterSections, err := s.mstSectionRepo.FilterBreakdowns(organizationId)
	if err != nil {
		return fmt.Errorf("failed to fetch master sections: %w", err)
	}

	// Fetch existing sections for the given breakdownId
	existingSections, err := s.sectionRepo.GetBreakdownSectionsByBreakdownId(breakdownId)
	if err != nil {
		return fmt.Errorf("failed to fetch existing sections: %w", err)
	}

	// Create maps for easier lookup
	existingSectionMap := make(map[int]models.BreakdownSection)
	for _, section := range existingSections {
		existingSectionMap[section.BreakdownSectionId] = section
	}

	masterSectionMap := make(map[int]models.MstBreakdownSection)
	for _, masterSection := range masterSections {
		masterSectionMap[masterSection.MstBreakdownSectionId] = masterSection
	}

	// Sync sections
	for _, masterSection := range masterSections {
		if existingSection, exists := existingSectionMap[masterSection.MstBreakdownSectionId]; exists {
			// Update section if necessary
			if existingSection.SectionTitle != masterSection.SectionTitle || existingSection.Sort != masterSection.Sort {
				existingSection.SectionTitle = masterSection.SectionTitle
				existingSection.Sort = masterSection.Sort
				if err := s.sectionRepo.Update(&existingSection); err != nil {
					return fmt.Errorf("failed to update section: %w", err)
				}
			}

			// Sync items within the section
			masterItems, err := s.mstItemRepo.GetMstBreakdownItemsBySectionId(masterSection.MstBreakdownSectionId)
			if err != nil {
				return fmt.Errorf("failed to fetch master items for section %d: %w", masterSection.MstBreakdownSectionId, err)
			}
			if err := s.syncItems(existingSection.BreakdownSectionId, masterItems); err != nil {
				return err
			}
		} else {
			// Insert new section and its items
			newSection := models.BreakdownSection{
				BreakdownId:    breakdownId,
				SectionTitle:   masterSection.SectionTitle,
				Sort:           masterSection.Sort,
				CreatedBy:      masterSection.CreatedBy,
				OrganizationId: masterSection.OrganizationId,
			}
			if err := s.sectionRepo.Create(&newSection); err != nil {
				return fmt.Errorf("failed to create new section: %w", err)
			}

			// Add items to the new section
			masterItems, err := s.mstItemRepo.GetMstBreakdownItemsBySectionId(masterSection.MstBreakdownSectionId)
			if err != nil {
				return fmt.Errorf("failed to fetch master items for section %d: %w", masterSection.MstBreakdownSectionId, err)
			}
			for _, masterItem := range masterItems {
				newItem := models.BreakdownItem{
					SectionId:      newSection.BreakdownSectionId,
					Description:    masterItem.Description,
					Sort:           masterItem.Sort,
					UnitPrice:      0,
					CreatedBy:      masterItem.CreatedBy,
					OrganizationId: masterItem.OrganizationId,
				}
				if err := s.itemRepo.Create(&newItem); err != nil {
					return fmt.Errorf("failed to create new item: %w", err)
				}
			}
		}
	}

	// Delete sections not in the master
	for _, existingSection := range existingSections {
		if _, exists := masterSectionMap[existingSection.BreakdownSectionId]; !exists {
			if err := s.sectionRepo.Delete(existingSection.BreakdownSectionId); err != nil {
				return fmt.Errorf("failed to delete section: %w", err)
			}
		}
	}

	return s.PublishFullReindexEvent(breakdownId)

}

// syncItems synchronizes BreakdownItems within a section.
func (s *BreakdownService) syncItems(sectionId int, masterItems []models.MstBreakdownItem) error {
	// Fetch existing items for the section
	existingItems, err := s.itemRepo.GetBreakdownItemsBySectionId(sectionId)
	if err != nil {
		return fmt.Errorf("failed to fetch existing items: %w", err)
	}

	// Create maps for easier lookup
	existingItemMap := make(map[int]models.BreakdownItem)
	for _, item := range existingItems {
		existingItemMap[item.BreakdownItemId] = item
	}

	masterItemMap := make(map[int]models.MstBreakdownItem)
	for _, masterItem := range masterItems {
		masterItemMap[masterItem.MstBreakdownItemId] = masterItem
	}

	// Sync items
	for _, masterItem := range masterItems {
		if existingItem, exists := existingItemMap[masterItem.MstBreakdownItemId]; exists {
			// Update item if necessary
			if existingItem.Description != masterItem.Description || existingItem.Sort != masterItem.Sort {
				existingItem.Description = masterItem.Description
				existingItem.Sort = masterItem.Sort
				if err := s.itemRepo.Update(&existingItem); err != nil {
					return fmt.Errorf("failed to update item: %w", err)
				}
			}
		} else {
			// Insert new item
			newItem := models.BreakdownItem{
				SectionId:      sectionId,
				Description:    masterItem.Description,
				Sort:           masterItem.Sort,
				UnitPrice:      0,
				CreatedBy:      masterItem.CreatedBy,
				OrganizationId: masterItem.OrganizationId,
			}
			if err := s.itemRepo.Create(&newItem); err != nil {
				return fmt.Errorf("failed to create new item: %w", err)
			}
		}
	}

	// Delete items not in the master
	for _, existingItem := range existingItems {
		if _, exists := masterItemMap[existingItem.BreakdownItemId]; !exists {
			if err := s.itemRepo.Delete(existingItem.BreakdownItemId); err != nil {
				return fmt.Errorf("failed to delete item: %w", err)
			}
		}
	}

	return nil
}

func (s *BreakdownService) ImportFromExcel(metadata dto.BreakdownImportDTO, fileBytes []byte, organizationID int, userID int) (models.Breakdown, error) {
	// Load Excel from bytes
	f, err := excelize.OpenReader(bytes.NewReader(fileBytes))
	if err != nil {
		return models.Breakdown{}, fmt.Errorf("failed to open Excel file: %v", err)
	}

	// Get all rows from the first sheet
	rows, err := f.GetRows(f.GetSheetName(0))
	if err != nil {
		return models.Breakdown{}, fmt.Errorf("failed to read Excel rows: %v", err)
	}

	currentDate := time.Now().Format("2006-01-02")
	parsedDate, err := time.Parse("2006-01-02", currentDate)
	if err != nil {
		return models.Breakdown{}, fmt.Errorf("failed to parse date: %v", err)
	}

	// Create Breakdown record
	breakdown := models.Breakdown{
		ProjectId:      *metadata.ProjectId,
		Subject:        *metadata.Subject,
		Date:           &models.CustomDate{Time: parsedDate},
		Location:       metadata.Location,
		OrganizationId: &organizationID,
		CreatedBy:      &userID,
	}

	if err := s.breakdownRepo.Create(&breakdown); err != nil {
		return breakdown, fmt.Errorf("failed to create SPK: %v", err)
	}

	var currentSectionID int
	var sectionSortOrder = 0 // Start section sort order from 1
	var itemSortOrder = 1    // Start item sort order from 1 within each section

	headerFound := false
	var grandTotal float64

	for rowIndex, row := range rows {
		if len(row) == 0 {
			continue
		}

		// Detect header row
		if !headerFound && len(row) >= 6 &&
			strings.EqualFold(strings.TrimSpace(row[1]), "NO") &&
			strings.EqualFold(strings.TrimSpace(row[2]), "DESCRIPTION") &&
			strings.EqualFold(strings.TrimSpace(row[8]), "HARGA SATUAN") {
			log.Printf("Header row found at line %d", rowIndex+1)
			headerFound = true
			continue
		}

		if !headerFound {
			log.Printf("Skipping row %d: header not found yet", rowIndex+1)
			continue
		}

		// Handle sections
		if isAlphabet(row[1]) {
			if len(row) < 2 {
				log.Printf("Skipping invalid section row at line %d: insufficient columns", rowIndex+1)
				continue
			}
			sectionSortOrder++
			section := models.BreakdownSection{
				BreakdownId:    breakdown.BreakdownId,
				SectionTitle:   row[2],
				Sort:           sectionSortOrder, // Set incremental sort order
				OrganizationId: &organizationID,
				CreatedBy:      &userID,
			}
			if err := s.sectionRepo.Create(&section); err != nil {
				return breakdown, fmt.Errorf("failed to create Breakdown section: %v", err)
			}
			currentSectionID = section.BreakdownSectionId
			itemSortOrder = 0 // Reset item sort order for new section
			continue
		}

		// Handle items
		if isNumeric(row[1]) {
			if len(row) < 7 {
				log.Printf("Skipping invalid detail row at line %d: insufficient columns", rowIndex+1)
				continue
			}

			totalPrice := parseFloatNonPointer(row[8])
			if totalPrice != 0 {
				grandTotal += totalPrice
			}

			itemSortOrder++
			detail := models.BreakdownItem{
				SectionId:      currentSectionID,
				Description:    row[2],
				UnitPrice:      totalPrice,
				Sort:           itemSortOrder, // Set incremental sort order
				OrganizationId: &organizationID,
				CreatedBy:      &userID,
			}
			if err := s.itemRepo.Create(&detail); err != nil {
				return breakdown, fmt.Errorf("failed to create Breakdown detail: %v", err)
			}
		}
	}

	breakdown.Total = &grandTotal

	if err := s.breakdownRepo.Update(&breakdown); err != nil {
		log.Printf("Error updating Breakdown: %v", err)
	}

	s.PublishFullReindexEvent(breakdown.BreakdownId)

	return breakdown, nil
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
	// Remove "Rp" and trim spaces
	cleaned := strings.ReplaceAll(input, "Rp", "")
	// Remove commas and trim spaces
	cleaned = strings.ReplaceAll(cleaned, ",", "")
	cleaned = strings.TrimSpace(cleaned)

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
	cleaned := strings.ReplaceAll(input, "Rp", "")
	// Remove commas and trim spaces
	cleaned = strings.ReplaceAll(cleaned, ",", "")
	cleaned = strings.TrimSpace(cleaned)

	// Attempt to parse the cleaned string
	value, err := strconv.ParseFloat(cleaned, 64)
	if err != nil {
		return 0 // Return 0 if parsing fails (or choose another default value)
	}
	return value
}
