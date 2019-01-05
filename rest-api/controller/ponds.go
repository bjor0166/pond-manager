package controller

import (
	"net/http"
	"encoding/json"
	"fmt"

	"pond-manager/rest-api/types"
)

var ponds []types.Pond

// BuildPonds creates some ponds and returns a pond
func BuildPonds() types.Pond {
	ponds = append(ponds, types.Pond{ID: "1", Name: "Lilly Pond"})
	ponds = append(ponds, types.Pond{ID: "2", Name: "Golden Pond"})
	return ponds[1]
}
// GetPonds provides list of ponds
func GetPonds(w http.ResponseWriter, r *http.Request){
	// return ponds array
	json.NewEncoder(w).Encode(ponds)
	fmt.Println("printing ponds")
}