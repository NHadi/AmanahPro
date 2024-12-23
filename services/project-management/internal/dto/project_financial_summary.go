package dto

type ProjectFinancialSummaryDTO struct {
	ProjectID   int     `json:"ProjectID"`
	ProjectName string  `json:"ProjectName"`
	Tanggal     string  `json:"Tanggal"`
	PO_SPH      float64 `json:"PO_SPH"`
	Termin      float64 `json:"Termin"`
	Operational float64 `json:"Operational"`
	Deviden     float64 `json:"Deviden"`
	SPKMandor   float64 `json:"SPKMandor"`
	BayarMandor float64 `json:"BayarMandor"`
	SisaBayar   float64 `json:"SisaBayar"`
	BB          float64 `json:"BB"`
	OPR         float64 `json:"OPR"`
	Saldo       float64 `json:"Saldo"`
	FEE         float64 `json:"FEE"`
}
