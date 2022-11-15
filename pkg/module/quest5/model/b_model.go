package model

type Quest5BResponse struct {
	Data []Quest5BData `json:"data"`
}

type Quest5BData struct {
	ID         int     `json:"id"`
	Country    string  `json:"county"`
	City       string  `json:"city"`
	Year       int     `json:"year"`
	Pm25       float64 `json:"pm25"`
	Latitude   float64 `json:"latitude"`
	Longtitude float64 `json:"longtitude"`
	Population float64 `json:"population"`
	Wbinc16    string  `json:"wbinc16"`
	Region     string  `json:"region"`
	ConcPm25   string  `json:"concPm25"`
	ColorPm25  string  `json:"colorPm25"`
	Distance   float64 `json:"distance"`
}
