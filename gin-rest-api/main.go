package main

import (
	"go-gin-api/database"
	"go-gin-api/routes"
)

func main() {
	database.Connect()

	routes.HandleRequests()

}
