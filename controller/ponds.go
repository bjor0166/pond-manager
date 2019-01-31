package controller

import (
	"net/http"
	"html/template"
	"encoding/json"
	"fmt"
	"time"

	"pond-manager/types"
	"pond-manager/database"
)

var ponds []types.Pond

//Instantiate a Welcome struct object and pass in some random information. 
//We shall get the name of the user as a query parameter from the URL
var welcome = types.Welcome{"Anonymous", time.Now().Format(time.Stamp)}

//We tell Go exactly where we can find our html file. We ask Go to parse the html file (Notice
// the relative path). We wrap it in a call to template.Must() which handles any errors and halts if there are fatal errors
var templates = template.Must(template.ParseFiles("templates/welcome-template.html"))

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

//This method takes in the URL path "/" and a function that takes in a response writer, and a http request.
func WelcomeFunc(w http.ResponseWriter, r *http.Request) {
	//Takes the name from the URL query e.g ?name=Martin, will set welcome.Name = Martin.
	if name := r.FormValue("name"); name != "" {
		welcome.Name = name;
	}
	//If errors show an internal server error message
	//I also pass the welcome struct to the welcome-template.html file.
	if err := templates.ExecuteTemplate(w, "welcome-template.html", welcome); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func FileServeFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("loading js...")
	http.ServeFile(w, r, "public/main.js")
}

func FaviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/water-favicon.png")
}

// CSS serves the welcome-template.css file 
func CSS(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/css/welcome-template.css")
}

// when requests come for /pond-background.jpg, serve the file 
func Background(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/pond-background.jpg")
}