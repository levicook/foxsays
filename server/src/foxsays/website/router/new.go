package router

import (
	"github.com/gorilla/mux"
)

func New() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.Path(route.path).
			Methods(route.method).
			Handler(route).
			Name(route.name)
	}

	return router
}
