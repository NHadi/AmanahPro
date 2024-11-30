package services

import (
	"AmanahPro/services/breakdown-services/common/messagebroker"
	"AmanahPro/services/breakdown-services/internal/domain/models"
	"AmanahPro/services/breakdown-services/internal/domain/repositories"
	"fmt"
	"log"
)

type BreakdownService struct {
	breakdownRepo      repositories.BreakdownRepository
	sectionRepo        repositories.BreakdownSectionRepository
	itemRepo           repositories.BreakdownItemRepository
	rabbitPublisher    *messagebroker.RabbitMQPublisher
	breakdownQueueName string
}

// NewBreakdownService initializes the BreakdownService
func NewBreakdownService(
	breakdownRepo repositories.BreakdownRepository,
	sectionRepo repositories.BreakdownSectionRepository,
	itemRepo repositories.BreakdownItemRepository,
	rabbitPublisher *messagebroker.RabbitMQPublisher,
	breakdownQueueName string,
) *BreakdownService {
	return &BreakdownService{
		breakdownRepo:      breakdownRepo,
		sectionRepo:        sectionRepo,
		itemRepo:           itemRepo,
		rabbitPublisher:    rabbitPublisher,
		breakdownQueueName: breakdownQueueName,
	}
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
