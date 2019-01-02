package router

import (
	"github.com/gorilla/mux"

	"github.com/bjor0166/controller"
	"github.com/bjor0166/types"
)

var Routes = []types.Route{
	router.HandleFunc("/ponds", GetPonds).Methods("GET")
	{"GET","/ponds",controller.Ponds,"ponds"}
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
