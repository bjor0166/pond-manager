package main

import (
	"log"
	"net/http"
	"fmt"

	"pond-manager/constants"
	"pond-manager/router"
	"pond-manager/database"
	// "pond-manager/controller"
)

func main () {
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Start the database connection
	db := database.Connect()		
	fmt.Println("connected to: ", db)
	// build a sample pond
	// golden := controller.BuildPonds()
	// fmt.Println("golden is: ", golden)

	// add sample data to the collection
	// insertResult, err := db.InsertOne(context.TODO(), golden)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	appRouter := router.CreateRouter()
	fmt.Println("running on port "+constants.Port)
	log.Fatal(http.ListenAndServe("localhost"+":"+constants.Port, appRouter))
}	