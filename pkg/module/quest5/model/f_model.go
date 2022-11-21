package model

type Quest5FResponse struct {
	Data []Quest5FData `json:"data"`
}

type Quest5FData struct {
	ID        int     `json:"id"`
	Country   string  `json:"county"`
	City      string  `json:"city"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	// Population float64 `json:"population"`
}
