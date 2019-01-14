package controller

import (
	"net/http"
	"encoding/json"
	"fmt"

	"pond-manager/rest-api/types"
	"pond-manager/rest-api/database"
)

var ponds []types.Pond

// BuildPonds creates some ponds and returns a pond
func BuildPonds() types.Pond {
	ponds = append(ponds, types.Pond{ID: "1", Name: "Lilly Pond"})
	ponds = append(ponds, types.Pond{ID: "2", Name: "Golden Pond"})
	return ponds[1]
}

// PondLookup returns json data for a pond that was searched
func PondLookup(w http.ResponseWriter, r *http.Request) {
	var result = database.FindPond()
	json.NewEncoder(w).Encode(result)
}

// AllPonds retrieves an array of ponds from the mongodb and writes the results to /allponds
func AllPonds(w http.ResponseWriter, r *http.Request) {
	var results = database.FindAll()
	json.NewEncoder(w).Encode(results)
}

func DeleteAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting all records")
	database.DeleteAllRecords()
	json.NewEncoder(w).Encode("deleted")
}