package routes

import (
	"github.com/gorilla/mux"
	"github.com/sderohan/assessment.git/controller"
)

func InitRoutes() *mux.Router {
	route := mux.NewRouter()
	route.HandleFunc("/get-words", controller.GetTopTenWords).Methods("POST")
	return route
}
