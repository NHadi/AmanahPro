package dto

// ProjectFinancialSummaryDTO represents the header data (summary) for projects.
type ProjectFinancialSPVSummaryDTO struct {
	ProjectID       int                            `json:"ProjectID"`
	ProjectName     string                         `json:"ProjectName"`
	ProjectUserId   int                            `json:"ProjectUserId"`
	Category        string                         `json:"Category"`
	TotalUangMasuk  float64                        `json:"TotalUangMasuk"`
	TotalUangKeluar float64                        `json:"TotalUangKeluar"`
	Sisa            float64                        `json:"Sisa"`
	Details         []ProjectFinancialSPVDetailDTO `json:"Details" gorm:"-"` // Tell GORM to ignore this field
}

// ProjectFinancialDetailDTO represents the detail data for a specific project.
type ProjectFinancialSPVDetailDTO struct {
	ID              int     `json:"ID"`
	TransactionDate string  `json:"TransactionDate"`
	Description     string  `json:"Description"`
	Amount          float64 `json:"Amount"`
	TransactionType string  `json:"TransactionType"`
	ProjectUserId   int     `json:"ProjectUserId"`
	UserName        string  `json:"UserName"`
}
