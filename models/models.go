package models

type Property struct {
	ID           string  `json:"ID"`
	PlotNumber   int     `json:"BatchNumber"`
	Builder      string  `json:"Builder"`
	OwnedBy      string  `json:"OwnedBy"`
	Quantity     int     `json:"Quantity"`
	CurrentPrice float32 `json:"CurrentPrice"`
	SellingPrice float32 `json:"SellingPrice"`
}
