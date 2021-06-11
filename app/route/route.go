package route

import (
	"awesomeProject/app/controllers"
	"github.com/gorilla/mux"
)

func GetRoutes() *mux.Router {
	routes := mux.NewRouter().StrictSlash(false)
	routes = routes.PathPrefix("/api/v1/").Subrouter()
	routes.HandleFunc("/hello", controllers.Hello).Methods("GET")
	return routes
}
