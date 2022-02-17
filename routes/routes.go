package routes

import (
	"log"
	"net/http"

	"github.com/ericklima-ca/go-api-rest/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/orders", controllers.ListAllOrders).Methods("Get")
	r.HandleFunc("/api/orders/{id}", controllers.GetOrderById).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", r))
}
