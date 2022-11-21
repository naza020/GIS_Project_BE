package model

type Quest5DResponse struct {
	Data []Quest5DMBR `json:"data"`
}

type Quest5DMBR struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
