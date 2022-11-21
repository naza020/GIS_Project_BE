package model

type Quest5DResponse struct {
	Data []Quest5DMBR `json:"data"`
}

type Quest5DMBR struct {
	Latitude   float64 `json:"latitude"`
	Longtitude float64 `json:"longtitude"`
}
