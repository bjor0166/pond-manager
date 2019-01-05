package main

import (
	"log"
	"net/http"
	"fmt"
	"context"

	"pond-manager/rest-api/constants"
	"pond-manager/rest-api/router"
	"pond-manager/rest-api/controller"
	"pond-manager/rest-api/types"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	// "github.com/mongodb/mongo-go-driver/mongo/options"
)
	
func main () {
	fmt.Println("I'm alive..")

	// build a sample pond
	golden := controller.BuildPonds()

// ---MONGODB---
	// connect to mongodb
	client, err := mongo.Connect(context.TODO(), "mongodb+srv://pond_admin:pondpw@pondcluster-lhxrp.mongodb.net/test?retryWrites=true")

	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}


	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	} else {
		fmt.Println("Connected to MongoDB!")
	}

	collection := client.Database(constants.DbName).Collection(constants.CollectionName)
// ---MONGODB---

	// ADD A RECORD
	// add sample data to the collection
	insertResult, err := collection.InsertOne(context.TODO(), golden)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	// FIND A VALUE
	filter := bson.D{{"name", "Golden Pond"}}

	update := bson.D{
		{"$inc", bson.D{
			{"age", 9},
		}},
	}
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	// create a value into which the result can be decoded
	var result types.Pond

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", result)

	appRouter := router.CreateRouter()
	fmt.Println("running on port "+constants.Port)
	log.Fatal(http.ListenAndServe(":"+constants.Port, appRouter))
}