package main

import (
	"fmt"

	"github.com/ericklima-ca/go-api-rest/database"
	"github.com/ericklima-ca/go-api-rest/routes"
)

func main() {
	database.Conn()
	fmt.Println("Server running...")
	routes.HandleRequest()
}
