package router

import "github.com/gorilla/mux"

func New() *mux.Router {
	routes := buildRoutes()
	return routes.Router()
}
