package model

type Quest4BResponse struct {
	Data []Quest4BData `json:"data"`
}

type Quest4BData struct {
	Country string  `json:"county"`
	AvgPm25 float64 `json:"avgPm25"`
}
