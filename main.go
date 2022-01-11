package main

import (
	"log"
	"net/http"

	"github.com/sderohan/assessment.git/routes"
)

func main() {
	route := routes.InitRoutes()
	log.Println("Server listening on port 8000")
	http.ListenAndServe(":8000", route)
}
