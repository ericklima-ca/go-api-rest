package models

type Product struct {
	Sku         string  `json:"sku"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
}
