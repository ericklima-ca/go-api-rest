package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ericklima-ca/go-api-rest/database"
	"github.com/ericklima-ca/go-api-rest/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func ListAllOrders(w http.ResponseWriter, r *http.Request) {
	var orders []models.Order
	database.DB.Find(&orders)
	json.NewEncoder(w).Encode(orders)
}

func GetOrderById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var order models.Order

	database.DB.First(&order, id)
	json.NewEncoder(w).Encode(order)
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var newOrder models.Order
	json.NewDecoder(r.Body).Decode(&newOrder)
	log.Print(r.Body)
	database.DB.Create(&newOrder)
	json.NewEncoder(w).Encode(newOrder)
}

func DeleteOrderById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var order models.Order

	database.DB.Delete(&order, id)
	json.NewEncoder(w).Encode(order)
}

func EditOrderById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var order models.Order

	database.DB.First(&order, id)
	json.NewDecoder(r.Body).Decode(&order)
	database.DB.Save(&order)
	json.NewEncoder(w).Encode(order)
}
