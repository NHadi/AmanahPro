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

// GetAllByProjectID retrieves all financial records for a specific project, including related ProjectUser data
func (r *projectFinancialRepositoryImpl) GetAllByProjectID(projectID int) ([]models.ProjectFinancial, error) {
	log.Printf("Retrieving ProjectFinancial records for ProjectID: %d", projectID)

	var financials []models.ProjectFinancial
	if err := r.db.
		Where("ProjectID = ?", projectID).
		Preload("ProjectUser").Preload("Project"). // Preload the ProjectUser relationship
		Find(&financials).Error; err != nil {
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
		-- Total amount of income
		ISNULL(SUM(CASE WHEN pf.TransactionType = 'Income' THEN pf.Amount ELSE 0 END), 0) AS Operational,
		-- Total AmountDeviden of income
		ISNULL(SUM(CASE WHEN pf.TransactionType = 'Income' THEN pf.AmountDeviden ELSE 0 END), 0) AS Deviden,
		-- Sum of Operational and Deviden
		ISNULL(SUM(CASE WHEN pf.TransactionType = 'Income' THEN pf.Amount ELSE 0 END), 0) +
		ISNULL(SUM(CASE WHEN pf.TransactionType = 'Income' THEN pf.AmountDeviden ELSE 0 END), 0) AS Termin,
        p.TotalSPK AS SPKMandor,
        ISNULL(SUM(CASE WHEN pf.ProjectUserId IS NOT NULL THEN pf.Amount ELSE 0 END), 0) AS BayarMandor,
        (p.TotalSPK - ISNULL(SUM(CASE WHEN pf.ProjectUserId IS NOT NULL THEN pf.Amount ELSE 0 END), 0)) AS SisaBayar,
        ISNULL(SUM(CASE WHEN pf.Category = 'BB' THEN pf.Amount ELSE 0 END), 0) AS BB,
        ISNULL(SUM(CASE WHEN pf.Category = 'Operational' THEN pf.Amount ELSE 0 END), 0) AS OPR,
		  ISNULL(SUM(CASE WHEN pf.Category = 'FEE' THEN pf.Amount ELSE 0 END), 0) AS FEE,
        (
            ISNULL(SUM(CASE WHEN pf.TransactionType = 'Income' THEN pf.Amount ELSE 0 END), 0) -- Operational (Total Income)
            - ISNULL(SUM(CASE WHEN pf.ProjectUserId IS NOT NULL THEN pf.Amount ELSE 0 END), 0) -- Bayar Mandor
            - ISNULL(SUM(CASE WHEN pf.Category = 'BB' THEN pf.Amount ELSE 0 END), 0) -- BB
            - ISNULL(SUM(CASE WHEN pf.Category = 'Operational' THEN pf.Amount ELSE 0 END), 0) -- OPR
			- ISNULL(SUM(CASE WHEN pf.Category = 'FEE' THEN pf.Amount ELSE 0 END), 0) -- FEE
        ) AS Saldo
    FROM Projects p
    LEFT JOIN ProjectFinancial pf ON p.ProjectID = pf.ProjectID
    WHERE p.OrganizationID = ?
    GROUP BY 
        p.ProjectID, 
        p.ProjectName, 
        p.PO_SPHDate, 
        p.SPH, 
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

// GetProjectFinancialSPVSummary retrieves financial summary data for all projects for a specific user, including details.
func (s *projectFinancialRepositoryImpl) GetProjectFinancialSPVSummary(userID int) ([]dto.ProjectFinancialSPVSummaryDTO, error) {
	var summaries []dto.ProjectFinancialSPVSummaryDTO

	// Query for summaries
	summaryQuery := `
    SELECT 
        a.ProjectID,
        b.ProjectName,
		a.ProjectUserId,
		a.Category,
        SUM(CASE WHEN a.TransactionType = 'Expense' THEN a.Amount ELSE 0 END) AS TotalUangMasuk,
        SUM(CASE WHEN a.TransactionType = 'Expense-SPV' THEN a.Amount ELSE 0 END) AS TotalUangKeluar,
        SUM(CASE WHEN a.TransactionType = 'Expense' THEN a.Amount ELSE 0 END) -
        SUM(CASE WHEN a.TransactionType = 'Expense-SPV' THEN a.Amount ELSE 0 END) AS Sisa
    FROM [AmanahProjectDb].[dbo].[ProjectFinancial] a
    LEFT JOIN Projects b ON a.ProjectID = b.ProjectID
    LEFT JOIN ProjectUser c ON a.ProjectUserId = c.Id
    WHERE c.Role = 'SPV' AND c.UserID = ?
    GROUP BY a.ProjectID, b.ProjectName,	a.ProjectUserId, a.Category;
    `

	// Execute the query for summaries
	if err := s.db.Debug().Raw(summaryQuery, userID).Scan(&summaries).Error; err != nil {
		log.Printf("Error fetching project financial summary: %v", err)
		return nil, err
	}

	// Populate details for each summary
	for i, summary := range summaries {
		var details []dto.ProjectFinancialSPVDetailDTO

		detailQuery := `
        SELECT 
			a.ID,
            a.TransactionDate,
            a.Descrtiption AS Description,
            a.Amount,
            a.TransactionType,
            c.UserrName as UserName
        FROM [AmanahProjectDb].[dbo].[ProjectFinancial] a
        LEFT JOIN ProjectUser c ON a.ProjectUserId = c.Id
        WHERE a.ProjectID = ? AND c.Role = 'SPV' AND c.UserID = ?
		order by a.TransactionDate desc;
        `

		// Execute the query for details
		if err := s.db.Raw(detailQuery, summary.ProjectID, userID).Scan(&details).Error; err != nil {
			log.Printf("Error fetching details for project ID %d: %v", summary.ProjectID, err)
			return nil, err
		}

		// Assign details to the summary
		summaries[i].Details = details
	}

	log.Println("Successfully retrieved project financial summary with details")
	return summaries, nil
}

// GetProjectFinancialDetails retrieves financial detail data for a specific project.
func (s *projectFinancialRepositoryImpl) GetProjectFinancialSPVDetails(userID int, projectID int) ([]dto.ProjectFinancialSPVDetailDTO, error) {
	var details []dto.ProjectFinancialSPVDetailDTO

	query := `
    SELECT 
        a.TransactionDate,
        a.Descrtiption,
        a.Amount,
        a.TransactionType,
        a.ProjectUserId,
        c.UserName
    FROM [AmanahProjectDb].[dbo].[ProjectFinancial] a
    LEFT JOIN Projects b ON a.ProjectID = b.ProjectID
    LEFT JOIN ProjectUser c ON a.ProjectUserId = c.Id
    WHERE c.Role = 'SPV' AND c.UserID = ? AND a.ProjectID = ?;
    `

	// Execute the query
	if err := s.db.Raw(query, userID, projectID).Scan(&details).Error; err != nil {
		log.Printf("Error fetching project financial details: %v", err)
		return nil, err
	}

	log.Println("Successfully retrieved project financial details")
	return details, nil
}
