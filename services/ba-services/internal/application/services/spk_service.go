package services

import (
	"AmanahPro/services/ba-services/internal/domain/models"
	"AmanahPro/services/ba-services/internal/domain/repositories"
	"context"
	"fmt"
	"log"

	"github.com/NHadi/AmanahPro-common/messagebroker"
	pb "github.com/NHadi/AmanahPro-common/protos"
)

type SpkService struct {
	baRepo          repositories.SPKRepository
	baSectionRepo   repositories.SPKSectionRepository
	detailRepo      repositories.SPKDetailRepository
	rabbitPublisher *messagebroker.RabbitMQPublisher
	baQueueName     string
	sphGrpcClient   pb.SphServiceClient // gRPC client for SPH
}

// NewSpkService initializes the SpkService
func NewSpkService(
	baRepo repositories.SPKRepository,
	baSectionRepo repositories.SPKSectionRepository,
	detailRepo repositories.SPKDetailRepository,
	rabbitPublisher *messagebroker.RabbitMQPublisher,
	baQueueName string,
	sphGrpcClient pb.SphServiceClient, // Inject SPH gRPC client
) *SpkService {
	return &SpkService{
		baRepo:          baRepo,
		baSectionRepo:   baSectionRepo,
		detailRepo:      detailRepo,
		rabbitPublisher: rabbitPublisher,
		baQueueName:     baQueueName,
		sphGrpcClient:   sphGrpcClient,
	}
}

// FilterSPKs filters SPKs by organization ID and optional SPK ID or project ID
func (s *SpkService) Filter(organizationID int, baID *int, projectID *int) ([]models.SPK, error) {
	log.Printf("Filtering SPKs for OrganizationID: %d, SpkID: %v, ProjectID: %v", organizationID, baID, projectID)

	bas, err := s.baRepo.Filter(organizationID, baID, projectID)
	if err != nil {
		log.Printf("Error filtering SPKs: %v", err)
		return nil, fmt.Errorf("failed to filter SPKs: %w", err)
	}

	log.Printf("Found %d SPKs for OrganizationID: %d", len(bas), organizationID)
	return bas, nil
}

// Helper: PublishFullReindexEvent sends a RabbitMQ event to re-index the full SPK
func (s *SpkService) PublishFullReindexEvent(baID int) error {
	log.Printf("Triggering re-index for SpkID: %d", baID)

	// Retrieve the full SPK structure for re-indexing
	ba, err := s.baRepo.GetByID(baID)
	if err != nil {
		log.Printf("Error retrieving SPK for reindexing: %v", err)
		return fmt.Errorf("error retrieving SPK for reindexing: %w", err)
	}

	if ba == nil {
		log.Printf("Reindex failed because ba ID: %d Not Found", baID)
		return nil
	}

	event := map[string]interface{}{
		"event":   "Reindexed",
		"payload": ba,
		"meta": map[string]interface{}{
			"idField": "SpkId",
		},
	}
	if err := s.rabbitPublisher.PublishEvent(s.baQueueName, event); err != nil {
		log.Printf("Error publishing re-index event for SpkID: %d, %v", baID, err)
		return fmt.Errorf("error publishing re-index event: %w", err)
	}

	log.Printf("Successfully triggered re-index for SpkID: %d", baID)
	return nil
}

// CRUD Operations for SPK
func (s *SpkService) GetSpkByID(baID int) (*models.SPK, error) {
	log.Printf("Fetching SPK with ID: %d", baID)

	ba, err := s.baRepo.GetByID(baID)
	if err != nil {
		log.Printf("Error fetching SPK: %v", err)
		return nil, fmt.Errorf("failed to fetch SPK: %w", err)
	}

	return ba, nil
}

// CreateSpk creates an SPK and populates sections and details from SPH
func (s *SpkService) CreateSpk(ba *models.SPK, sphId int32) error {
	log.Printf("Creating SPK: %+v with SPH ID: %d", ba, sphId)

	// Call SPH gRPC service to get sections and details
	log.Printf("Calling SPH gRPC service for SPH ID: %d", sphId)
	sphDetailsResponse, err := s.sphGrpcClient.GetSphDetails(context.Background(), &pb.GetSphDetailsRequest{SphId: sphId})
	if err != nil {
		log.Printf("Failed to fetch SPH details from gRPC: %v", err)
		return fmt.Errorf("failed to fetch SPH details from gRPC: %w", err)
	}

	// Populate SPK with sections and details from SPH
	for _, grpcSection := range sphDetailsResponse.Sections {
		// Create a section
		section := models.SPKSection{
			SphSectionId:   int(grpcSection.SphSectionId),
			SectionTitle:   &grpcSection.SectionTitle,
			CreatedBy:      ba.CreatedBy,
			OrganizationId: ba.OrganizationId,
		}

		// Populate details for this section
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
				CreatedBy:      ba.CreatedBy,
				OrganizationId: ba.OrganizationId,
			}
			// Add detail to section
			section.Details = append(section.Details, detail)
		}

		// Add section to SPK
		ba.Sections = append(ba.Sections, section)
	}

	// Save the SPK
	if err := s.baRepo.Create(ba); err != nil {
		log.Printf("Error creating SPK: %v", err)
		return fmt.Errorf("error creating SPK: %w", err)
	}

	s.PublishFullReindexEvent(ba.SpkId)

	log.Printf("Successfully created SPK: %+v", ba)
	return nil
}

func (s *SpkService) UpdateSpk(ba *models.SPK) error {
	log.Printf("Updating SPK: %+v", ba)
	if err := s.baRepo.Update(ba); err != nil {
		log.Printf("Error updating SPK: %v", err)
		return fmt.Errorf("error updating SPK: %w", err)
	}

	return s.PublishFullReindexEvent(ba.SpkId)
}

func (s *SpkService) DeleteSpk(baID int) error {
	log.Printf("Deleting SPK ID: %d", baID)
	if err := s.baRepo.Delete(baID); err != nil {
		log.Printf("Error deleting SPK: %v", err)
		return fmt.Errorf("error deleting SPK: %w", err)
	}

	// Send a delete event to RabbitMQ
	event := map[string]interface{}{
		"event":   "Deleted",
		"payload": map[string]int{"SpkId": baID},
		"meta": map[string]interface{}{
			"idField": "SpkId",
		},
	}
	if err := s.rabbitPublisher.PublishEvent(s.baQueueName, event); err != nil {
		log.Printf("Error publishing delete event for SpkID: %d, %v", baID, err)
		return fmt.Errorf("error publishing delete event: %w", err)
	}

	log.Printf("Successfully deleted SPK ID: %d", baID)
	return nil
}

// CRUD Operations for SPK Sections

func (s *SpkService) CreateSpkSection(section *models.SPKSection, baID int) error {
	log.Printf("Creating SPK Section for SPK ID: %d", baID)

	// Assign the SPK ID
	section.SpkId = baID

	// Save the section to the repository
	if err := s.baSectionRepo.Create(section); err != nil {
		log.Printf("Error creating SPK Section: %v", err)
		return fmt.Errorf("failed to create SPK Section: %w", err)
	}

	s.PublishFullReindexEvent(baID)

	log.Printf("Successfully created SPK Section with ID: %d", section.SectionId)
	return nil
}

func (s *SpkService) UpdateSpkSection(updatedSection *models.SPKSection) error {
	log.Printf("Updating SPK Section ID: %d", updatedSection.SectionId)

	// Update the section in the repository
	if err := s.baSectionRepo.Update(updatedSection); err != nil {
		log.Printf("Error updating SPK Section: %v", err)
		return fmt.Errorf("failed to update SPK Section: %w", err)
	}

	s.PublishFullReindexEvent(updatedSection.SpkId)

	log.Printf("Successfully updated SPK Section with ID: %d", updatedSection.SectionId)
	return nil
}

func (s *SpkService) DeleteSpkSection(sectionID, SPKId int) error {
	log.Printf("Deleting SPK Section with ID: %d", sectionID)

	// Delete the section from the repository
	if err := s.baSectionRepo.Delete(sectionID); err != nil {
		log.Printf("Error deleting SPK Section: %v", err)
		return fmt.Errorf("failed to delete SPK Section: %w", err)
	}

	s.PublishFullReindexEvent(SPKId)

	log.Printf("Successfully deleted SPK Section with ID: %d", sectionID)
	return nil
}

func (s *SpkService) GetSpkSectionByID(sectionID int) (*models.SPKSection, error) {
	log.Printf("Fetching SPK Section with ID: %d", sectionID)

	// Retrieve the section from the repository
	section, err := s.baSectionRepo.GetByID(sectionID)
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

	s.PublishFullReindexEvent(SPKId)

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

	s.PublishFullReindexEvent(SPKId)

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

	s.PublishFullReindexEvent(SPKId)

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
