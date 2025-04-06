package main

import (
	"api/routes"
)

func main() {
	router := routes.SetupRouter()
	router.Run(":8080")
}
