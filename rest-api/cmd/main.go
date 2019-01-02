package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"

	"pond-manager/constants"
)

// type Pond struct {
// 	ID	string `json:"id,omitempty"`
// 	Name string `json:"name,omitempty"`
// }

// var ponds []Pond

// func GetPonds(w http.ResponseWriter, r *http.Request) {
// 	json.NewEncoder(w).Encode(ponds)
// 	fmt.Println("ponds!!!")
// }

func main () {
	//ponds = append(ponds, Pond{ID: "1", Name: "Lilly Pond"})
	
	router := mux.NewRouter()
	//router.HandleFunc("/ponds", GetPonds).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+constants.Port, router))
}
