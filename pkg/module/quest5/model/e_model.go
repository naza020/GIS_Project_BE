package model

type Quest5EResponse struct {
	Data []Quest5EData `json:"data"`
}

type Quest5EData struct {
	ID         int     `json:"id"`
	Country    string  `json:"county"`
	City       string  `json:"city"`
	Latitude   float64 `json:"latitude"`
	Longtitude float64 `json:"longtitude"`
}
