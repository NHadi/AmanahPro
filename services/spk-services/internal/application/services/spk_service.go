package services

import (
	"AmanahPro/services/spk-services/common/messagebroker"
	"AmanahPro/services/spk-services/internal/domain/models"
	"AmanahPro/services/spk-services/internal/domain/repositories"
	"fmt"
	"log"
)

type SpkService struct {
	spkRepo         repositories.SPKRepository
	spkSectionRepo  repositories.SPKSectionRepository
	detailRepo      repositories.SPKDetailRepository
	rabbitPublisher *messagebroker.RabbitMQPublisher
	spkQueueName    string
}

// NewSpkService initializes the SpkService
func NewSpkService(
	spkRepo repositories.SPKRepository,
	spkSectionRepo repositories.SPKSectionRepository,
	detailRepo repositories.SPKDetailRepository,
	rabbitPublisher *messagebroker.RabbitMQPublisher,
	spkQueueName string,
) *SpkService {
	return &SpkService{
		spkRepo:         spkRepo,
		spkSectionRepo:  spkSectionRepo,
		detailRepo:      detailRepo,
		rabbitPublisher: rabbitPublisher,
		spkQueueName:    spkQueueName,
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
func (s *SpkService) PublishFullReindexEvent(spkID int) error {
	log.Printf("Triggering re-index for SpkID: %d", spkID)

	// Retrieve the full SPK structure for re-indexing
	spk, err := s.spkRepo.GetByID(spkID)
	if err != nil {
		log.Printf("Error retrieving SPK for reindexing: %v", err)
		return fmt.Errorf("error retrieving SPK for reindexing: %w", err)
	}

	event := map[string]interface{}{
		"event":   "Reindexed",
		"payload": spk,
		"meta": map[string]interface{}{
			"idField": "SpkId",
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

	spk, err := s.spkRepo.GetByID(spkID)
	if err != nil {
		log.Printf("Error fetching SPK: %v", err)
		return nil, fmt.Errorf("failed to fetch SPK: %w", err)
	}

	return spk, nil
}

func (s *SpkService) CreateSpk(spk *models.SPK) error {
	log.Printf("Creating SPK: %+v", spk)
	if err := s.spkRepo.Create(spk); err != nil {
		log.Printf("Error creating SPK: %v", err)
		return fmt.Errorf("error creating SPK: %w", err)
	}

	event := map[string]interface{}{
		"event":   "Created",
		"payload": spk,
		"meta": map[string]interface{}{
			"idField": "SpkId",
		},
	}
	if err := s.rabbitPublisher.PublishEvent(s.spkQueueName, event); err != nil {
		log.Printf("Error publishing index event for SpkID: %d, %v", spk.SpkId, err)
		return fmt.Errorf("error publishing index event: %w", err)
	}

	log.Printf("Successfully triggered index for SpkID: %d", spk.SpkId)
	return nil
}

func (s *SpkService) UpdateSpk(spk *models.SPK) error {
	log.Printf("Updating SPK: %+v", spk)
	if err := s.spkRepo.Update(spk); err != nil {
		log.Printf("Error updating SPK: %v", err)
		return fmt.Errorf("error updating SPK: %w", err)
	}

	return s.PublishFullReindexEvent(spk.SpkId)
}

func (s *SpkService) DeleteSpk(spkID int) error {
	log.Printf("Deleting SPK ID: %d", spkID)
	if err := s.spkRepo.Delete(spkID); err != nil {
		log.Printf("Error deleting SPK: %v", err)
		return fmt.Errorf("error deleting SPK: %w", err)
	}

	// Send a delete event to RabbitMQ
	event := map[string]interface{}{
		"event":   "Deleted",
		"payload": map[string]int{"SpkId": spkID},
		"meta": map[string]interface{}{
			"idField": "SpkId",
		},
	}
	if err := s.rabbitPublisher.PublishEvent(s.spkQueueName, event); err != nil {
		log.Printf("Error publishing delete event for SpkID: %d, %v", spkID, err)
		return fmt.Errorf("error publishing delete event: %w", err)
	}

	log.Printf("Successfully deleted SPK ID: %d", spkID)
	return nil
}


