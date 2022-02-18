package models

import _ "gorm.io/gorm"

type Order struct {
	Id       int    `json:"id"`
	Customer int    `json:"customer"`
	Product  string `json:"product"`
	Quantity int    `json:"quantity"`
}

var Orders []Order
