package services

import (
	"AmanahPro/services/project-management/internal/domain/models"
	"AmanahPro/services/project-management/internal/domain/repositories"
	"AmanahPro/services/project-management/internal/dto"
	"fmt"
	"log"
)

type ProjectFinancialService struct {
	projectFinancialRepo repositories.ProjectFinancialRepository
}

func NewProjectFinancialService(
	projectFinancialRepo repositories.ProjectFinancialRepository,
) *ProjectFinancialService {
	return &ProjectFinancialService{
		projectFinancialRepo: projectFinancialRepo,
	}
}

// GetProjectFinancialSummary retrieves financial summary data for all projects by OrganizationID
func (s *ProjectFinancialService) GetProjectFinancialSummary(organizationID int) ([]dto.ProjectFinancialSummaryDTO, error) {
	var summaries []dto.ProjectFinancialSummaryDTO

	summaries, err := s.projectFinancialRepo.GetProjectFinancialSummary(organizationID)
	if err != nil {
		log.Printf("Error fetching project financial summary: %v", err)
		return nil, fmt.Errorf("failed to fetch records: %w", err)
	}

	log.Println("Successfully retrieved project financial summary")
	return summaries, nil
}

// CreateProjectFinancial creates a new financial record
func (s *ProjectFinancialService) CreateProjectFinancial(financial *models.ProjectFinancial, traceID string) error {
	log.Printf("TraceID: %s - Creating Project Financial record: %+v", traceID, financial)

	// Input validation
	if financial.ProjectID == 0 {
		return fmt.Errorf("ProjectID is required")
	}
	if financial.Amount <= 0 {
		return fmt.Errorf("Amount must be greater than zero")
	}

	// Save to the database
	if err := s.projectFinancialRepo.Create(financial); err != nil {
		log.Printf("TraceID: %s - Error creating financial record: %v", traceID, err)
		return fmt.Errorf("error creating financial record: %w", err)
	}

	log.Printf("TraceID: %s - Successfully created Project Financial record", traceID)
	return nil
}

// UpdateProjectFinancial updates an existing financial record
func (s *ProjectFinancialService) UpdateProjectFinancial(financial *models.ProjectFinancial, traceID string) error {
	log.Printf("TraceID: %s - Updating Project Financial ID: %d", traceID, financial.ID)

	if financial.ID == 0 {
		return fmt.Errorf("ID is required for updating")
	}

	// Update the record
	if err := s.projectFinancialRepo.Update(financial); err != nil {
		log.Printf("TraceID: %s - Error updating financial record: %v", traceID, err)
		return fmt.Errorf("error updating financial record: %w", err)
	}

	log.Printf("TraceID: %s - Successfully updated Project Financial record ID: %d", traceID, financial.ID)
	return nil
}

// DeleteProjectFinancial deletes a financial record by ID
func (s *ProjectFinancialService) DeleteProjectFinancial(id int, traceID string, userID int) error {
	log.Printf("TraceID: %s - Deleting Project Financial ID: %d", traceID, id)

	if id <= 0 {
		return fmt.Errorf("Invalid ID for deletion")
	}

	// Perform deletion
	if err := s.projectFinancialRepo.Delete(id); err != nil {
		log.Printf("TraceID: %s - Error deleting financial record: %v", traceID, err)
		return fmt.Errorf("error deleting financial record: %w", err)
	}

	log.Printf("TraceID: %s - Successfully deleted Project Financial record ID: %d", traceID, id)
	return nil
}

// GetProjectFinancialByID retrieves a financial record by ID
func (s *ProjectFinancialService) GetProjectFinancialByID(id int, traceID string) (*models.ProjectFinancial, error) {
	log.Printf("TraceID: %s - Fetching Project Financial record ID: %d", traceID, id)

	if id <= 0 {
		return nil, fmt.Errorf("Invalid ID")
	}

	record, err := s.projectFinancialRepo.GetByID(id)
	if err != nil {
		log.Printf("TraceID: %s - Error fetching financial record ID: %d, %v", traceID, id, err)
		return nil, fmt.Errorf("failed to fetch financial record: %w", err)
	}

	if record == nil {
		log.Printf("TraceID: %s - Project Financial record ID: %d not found", traceID, id)
		return nil, nil
	}

	log.Printf("TraceID: %s - Successfully fetched Project Financial record ID: %d", traceID, id)
	return record, nil
}

// GetAllFinancialByProjectID retrieves all financial records for a given ProjectID
func (s *ProjectFinancialService) GetAllFinancialByProjectID(projectID int, traceID string) ([]models.ProjectFinancial, error) {
	log.Printf("TraceID: %s - Fetching all financial records for ProjectID: %d", traceID, projectID)

	if projectID <= 0 {
		return nil, fmt.Errorf("Invalid ProjectID")
	}

	records, err := s.projectFinancialRepo.GetAllByProjectID(projectID)
	if err != nil {
		log.Printf("TraceID: %s - Error fetching records for ProjectID: %d, %v", traceID, projectID, err)
		return nil, fmt.Errorf("failed to fetch records: %w", err)
	}

	log.Printf("TraceID: %s - Successfully fetched %d financial records for ProjectID: %d", traceID, len(records), projectID)
	return records, nil
}
