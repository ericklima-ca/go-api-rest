package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ericklima-ca/go-api-rest/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func ListAllOrders(w http.ResponseWriter, r *http.Request) {
	models.Orders = []models.Order{{
		Id:       1,
		Customer: models.Customer{Name: "Erick"},
		Product: models.Product{
			Sku:         "123456",
			Description: "TV",
			Price:       5000.0,
			Image:       "https://url.com",
		},
		Quantity: 2,
	}}
	json.NewEncoder(w).Encode(models.Orders)
}

func GetOrderById(w http.ResponseWriter, r *http.Request) {
	models.Orders = []models.Order{{
		Id:       1,
		Customer: models.Customer{Name: "Erick"},
		Product: models.Product{
			Sku:         "123456",
			Description: "TV",
			Price:       5000.0,
			Image:       "https://url.com",
		},
		Quantity: 2,
	}}
	
	vars := mux.Vars(r)
	id := vars["id"]

	for _, o := range models.Orders {
		if strconv.Itoa(o.Id) == id {
			json.NewEncoder(w).Encode(o)
		}
	}
}
