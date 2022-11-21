package model

type Quest5CResponse struct {
	Data []Quest5CData `json:"data"`
}

type Quest5CData struct {
	Country    string  `json:"county"`
	City       string  `json:"city"`
	Latitude   float64 `json:"latitude"`
	Longtitude float64 `json:"longtitude"`
}
