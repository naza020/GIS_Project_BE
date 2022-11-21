package model

type InsertDataRequest struct {
	Data []RequestData `json:"data"`
}
type RequestData struct {
	Country    string  `json:"country"`
	City       string  `json:"city"`
	Year       int     `json:"year"`
	Pm25       float64 `json:"pm25"`
	Latitude   float64 `json:"latitude"`
	Longtitude float64 `json:"longtitude"`
	Population float64 `json:"population"`
	Wbinc16    string  `json:"wbinc16"`
	Region     string  `json:"region"`
	Conc       string  `json:"conc"`
	Color      string  `json:"color"`
}

type InsertDataResponse struct {
	Message string `json:"message"`
	Record  int    `json:"record"`
}
