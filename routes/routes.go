package routes

import (
	"github.com/gorilla/mux"
	"github.com/sderohan/assessment.git/controller"
)

func InitRoutes() *mux.Router {
	route := mux.NewRouter()
	v1 := route.PathPrefix("/v1").Subrouter()
	v1.HandleFunc("/get-words", controller.GetTopTenWords).Methods("POST")
	return route
}
