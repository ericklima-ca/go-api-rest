package models

type Order struct {
	Id       int      `json:"id"`
	Customer Customer `json:"customer"`
	Product  Product  `json:"product"`
	Quantity int      `json:"quantity"`
}

var Orders []Order