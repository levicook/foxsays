package route

import "github.com/gorilla/mux"

type Routes []Route

func (routes Routes) Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			Handler(route)
	}

	return router
}
