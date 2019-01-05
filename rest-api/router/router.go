package router

import (
	"github.com/gorilla/mux"

	"pond-manager/rest-api/controller"
	"pond-manager/rest-api/types"
)

var Routes = []types.Route{
	{"GET","/ponds",controller.GetPonds,"ponds"},
}

//CreateRouter creates a new router
func CreateRouter () *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range Routes {
		router.
			Methods(route.Method).
			Path(route.Path).
			Handler(route.HandlerFunc).
			Name(route.Name)
	}
	return router
}
