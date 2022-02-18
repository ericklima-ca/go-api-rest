package routes

import (
	"log"
	"net/http"

	"github.com/ericklima-ca/go-api-rest/controllers"
	"github.com/ericklima-ca/go-api-rest/middlewares"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middlewares.ContentType)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/orders", controllers.ListAllOrders).Methods("Get")
	r.HandleFunc("/api/orders/{id}", controllers.GetOrderById).Methods("Get")
	r.HandleFunc("/api/orders", controllers.CreateOrder).Methods("Post")
	r.HandleFunc("/api/orders/{id}", controllers.DeleteOrderById).Methods("Delete")
	r.HandleFunc("/api/orders/{id}", controllers.EditOrderById).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000",
		handlers.CORS(
			handlers.AllowedOrigins(
				[]string{"*"},
			),
		)(r),
	),
	)
}
