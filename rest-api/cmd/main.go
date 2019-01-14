package main

import (
	"log"
	"net/http"
	"fmt"
	"context"

	"pond-manager/rest-api/constants"
	"pond-manager/rest-api/router"
	"pond-manager/rest-api/controller"
	"pond-manager/rest-api/database"
)

func main () {
	// Start the database connection
	db := database.Connect()
	fmt.Println("connected to: ", db)
	// build a sample pond
	golden := controller.BuildPonds()
	fmt.Println("golden is: ", golden)

	// add sample data to the collection
	insertResult, err := db.InsertOne(context.TODO(), golden)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	appRouter := router.CreateRouter()
	fmt.Println("running on port "+constants.Port)
	log.Fatal(http.ListenAndServe(":"+constants.Port, appRouter))
}