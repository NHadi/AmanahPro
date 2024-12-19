package repositories

import (
	"AmanahPro/services/project-management/internal/domain/models"
	"AmanahPro/services/project-management/internal/domain/repositories"
	"AmanahPro/services/project-management/internal/dto"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type projectFinancialRepositoryImpl struct {
	db *gorm.DB
}

// NewProjectFinancialRepository creates a new instance of ProjectFinancialRepository
func NewProjectFinancialRepository(db *gorm.DB) repositories.ProjectFinancialRepository {
	return &projectFinancialRepositoryImpl{
		db: db,
	}
}

// Create inserts a new ProjectFinancial record into the database
func (r *projectFinancialRepositoryImpl) Create(financial *models.ProjectFinancial) error {
	log.Printf("Creating ProjectFinancial: %+v", financial)

	if err := r.db.Create(financial).Error; err != nil {
		log.Printf("Failed to create ProjectFinancial: %v", err)
		return fmt.Errorf("failed to create ProjectFinancial: %w", err)
	}

	log.Printf("Successfully created ProjectFinancial: %+v", financial)
	return nil
}

// Update modifies an existing ProjectFinancial record in the database
func (r *projectFinancialRepositoryImpl) Update(financial *models.ProjectFinancial) error {
	log.Printf("Updating ProjectFinancial ID: %d", financial.ID)

	if err := r.db.Save(financial).Error; err != nil {
		log.Printf("Failed to update ProjectFinancial ID %d: %v", financial.ID, err)
		return fmt.Errorf("failed to update ProjectFinancial: %w", err)
	}

	log.Printf("Successfully updated ProjectFinancial ID: %d", financial.ID)
	return nil
}

// Delete removes a ProjectFinancial record from the database
func (r *projectFinancialRepositoryImpl) Delete(financialID int) error {
	log.Printf("Deleting ProjectFinancial ID: %d", financialID)

	if err := r.db.Delete(&models.ProjectFinancial{}, financialID).Error; err != nil {
		log.Printf("Failed to delete ProjectFinancial ID %d: %v", financialID, err)
		return fmt.Errorf("failed to delete ProjectFinancial: %w", err)
	}

	log.Printf("Successfully deleted ProjectFinancial ID: %d", financialID)
	return nil
}

// GetByID retrieves a ProjectFinancial record by its ID
func (r *projectFinancialRepositoryImpl) GetByID(financialID int) (*models.ProjectFinancial, error) {
	log.Printf("Retrieving ProjectFinancial by ID: %d", financialID)

	var financial models.ProjectFinancial
	if err := r.db.First(&financial, financialID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("ProjectFinancial ID %d not found", financialID)
			return nil, nil
		}
		log.Printf("Failed to retrieve ProjectFinancial ID %d: %v", financialID, err)
		return nil, fmt.Errorf("failed to retrieve ProjectFinancial: %w", err)
	}

	log.Printf("Successfully retrieved ProjectFinancial: %+v", financial)
	return &financial, nil
}

// GetAllByProjectID retrieves all financial records for a specific project
func (r *projectFinancialRepositoryImpl) GetAllByProjectID(projectID int) ([]models.ProjectFinancial, error) {
	log.Printf("Retrieving ProjectFinancial records for ProjectID: %d", projectID)

	var financials []models.ProjectFinancial
	if err := r.db.Where("ProjectID = ?", projectID).Find(&financials).Error; err != nil {
		log.Printf("Failed to retrieve records for ProjectID %d: %v", projectID, err)
		return nil, fmt.Errorf("failed to retrieve ProjectFinancial records: %w", err)
	}

	log.Printf("Successfully retrieved ProjectFinancial records for ProjectID: %d", projectID)
	return financials, nil
}

// GetProjectFinancialSummary retrieves financial summary data for all projects by OrganizationID
func (s *projectFinancialRepositoryImpl) GetProjectFinancialSummary(organizationID int) ([]dto.ProjectFinancialSummaryDTO, error) {
	var summaries []dto.ProjectFinancialSummaryDTO

	query := `
        SELECT 
            p.ProjectID,
            p.ProjectName,
            p.PO_SPHDate AS Tanggal,
            p.SPH AS PO_SPH,
            p.Termin AS Termin,
            ISNULL(SUM(CASE WHEN pf.TransactionType = 'Income' THEN pf.Amount ELSE 0 END), 0) AS Operational,
            (p.Termin - ISNULL(SUM(CASE WHEN pf.TransactionType = 'Income' THEN pf.Amount ELSE 0 END), 0)) AS Deviden,
            p.TotalSPK AS SPKMandor,
            ISNULL(SUM(CASE WHEN pf.ProjectUserId IS NOT NULL THEN pf.Amount ELSE 0 END), 0) AS BayarMandor,
            (p.TotalSPK - ISNULL(SUM(CASE WHEN pf.ProjectUserId IS NOT NULL THEN pf.Amount ELSE 0 END), 0)) AS SisaBayar,
            ISNULL(SUM(CASE WHEN pf.Category = 'BB' THEN pf.Amount ELSE 0 END), 0) AS BB,
            ISNULL(SUM(CASE WHEN pf.Category = 'Operational' THEN pf.Amount ELSE 0 END), 0) AS OPR,
            (
                ISNULL(SUM(CASE WHEN pf.TransactionType = 'Income' THEN pf.Amount ELSE 0 END), 0) -- Operational (Total Income)
                - ISNULL(SUM(CASE WHEN pf.ProjectUserId IS NOT NULL THEN pf.Amount ELSE 0 END), 0) -- Bayar Mandor
                - ISNULL(SUM(CASE WHEN pf.Category = 'BB' THEN pf.Amount ELSE 0 END), 0) -- BB
                - ISNULL(SUM(CASE WHEN pf.Category = 'Operational' THEN pf.Amount ELSE 0 END), 0) -- OPR
            ) AS Saldo
        FROM Projects p
        LEFT JOIN ProjectFinancial pf ON p.ProjectID = pf.ProjectID
        WHERE p.OrganizationID = ?
        GROUP BY 
            p.ProjectID, 
            p.ProjectName, 
            p.PO_SPHDate, 
            p.SPH, 
            p.Termin, 
            p.TotalSPK
    `

	// Execute the query with the OrganizationID as a filter
	if err := s.db.Raw(query, organizationID).Scan(&summaries).Error; err != nil {
		log.Printf("Error fetching project financial summary: %v", err)
		return nil, err
	}

	log.Println("Successfully retrieved project financial summary")
	return summaries, nil
}
