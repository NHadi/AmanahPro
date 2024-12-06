package services

import (
	"AmanahPro/services/spk-services/internal/domain/models"
	"AmanahPro/services/spk-services/internal/domain/repositories"
	"context"
	"fmt"
	"log"

	commonServices "github.com/NHadi/AmanahPro-common/services"

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

	return s.PublishFullReindexEvent(spk.SphId, traceID, *spk.UpdatedBy)
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
