package services

import (
	"AmanahPro/services/sph-services/internal/domain/models"
	"AmanahPro/services/sph-services/internal/domain/repositories"
	"fmt"
	"log"

	"github.com/NHadi/AmanahPro-common/messagebroker"
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
