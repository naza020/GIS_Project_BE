package model

type Quest5AResponse struct {
	Data []Quest5AData `json:"data"`
}

type Quest5AData struct {
	ID        int     `json:"id"`
	Country   string  `json:"county"`
	City      string  `json:"city"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	// Population float64 `json:"population"`
}

// type Quest5AData struct {
// 	ID         int     `json:"id"`
// 	Country    string  `json:"county"`
// 	City       string  `json:"city"`
// 	Year       int     `json:"year"`
// 	Pm25       float64 `json:"pm25"`
// 	Latitude   float64 `json:"latitude"`
// 	Longitude float64 `json:"longitude"`
// 	Population float64 `json:"population"`
// 	Wbinc16    string  `json:"wbinc16"`
// 	Region     string  `json:"region"`
// 	ConcPm25   string  `json:"concPm25"`
// 	ColorPm25  string  `json:"colorPm25"`
// }
