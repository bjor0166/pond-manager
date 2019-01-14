package router

import (
	"github.com/gorilla/mux"

	"pond-manager/rest-api/controller"
	"pond-manager/rest-api/types"
)

// Routes defines the routes of the API
var Routes = []types.Route{
	{"GET","/allponds",controller.AllPonds,"allponds"},
	{"GET","/findpond",controller.PondLookup,"findpond"},
	{"DELETE","/deleteall", controller.DeleteAll,"deleteall"},
}

//CreateRouter creates a new router with all of the routes in Routes
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
