package dto

type ProjectRecapDTO struct {
	TotalOpname      float64 `json:"TotalOpname"`
	TotalPengeluaran float64 `json:"TotalPengeluaran"`
	Margin           float64 `json:"Margin"`
	MarginPercentage float64 `json:"MarginPercentage"`
}
