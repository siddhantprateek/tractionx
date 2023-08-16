package models

type Property struct {
	ID           string  `json:"ID"`
	PlotNumber   int     `json:"BatchNumber"`
	Builder      string  `json:"Builder"`
	OwnedBy      string  `json:"OwnedBy"`
	CurrentPrice float32 `json:"CurrentPrice"`
	SellingPrice float32 `json:"SellingPrice"`
}
