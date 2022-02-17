package main

import (
	"fmt"

	"github.com/ericklima-ca/go-api-rest/routes"
)

func main() {

	fmt.Println("Server running...")
	routes.HandleRequest()
}
