package router

import (
	"github.com/gorilla/mux"

	"pond-manager/controller"
	"pond-manager/types"
)

// Routes defines the routes of the API
var Routes = []types.Route {
	{"GET","/welcome",controller.WelcomeFunc,"welcome"},
	{"GET","/public/main.js",controller.FileServeFunc,"fileserve"},
	{"GET","/allponds",controller.AllPonds,"allponds"},
	{"GET","/findpond",controller.PondLookup,"findpond"},
	{"DELETE","/deleteall", controller.DeleteAll,"deleteall"},
	{"GET","/water-favicon.png",controller.FaviconHandler,"favicon"},
	{"GET","/css/welcome-template.css",controller.CSS,"css"},
	{"GET","/pond-background.jpg",controller.Background,"background"},
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
