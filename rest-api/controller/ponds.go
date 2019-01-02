package controller

import (
	"net/http"
	"os"
)

type Pond struct {
	ID	string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

var ponds []Pond

func GetPonds(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(ponds)
	fmt.Println("ponds!!!")
}