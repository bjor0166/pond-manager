	package database

	import(
		"log"
		"context"
		"fmt"

		"pond-manager/constants"
		"pond-manager/types"
		"github.com/mongodb/mongo-go-driver/mongo"
		"github.com/mongodb/mongo-go-driver/bson"
		"github.com/mongodb/mongo-go-driver/mongo/options"
	)
	
	// PondsDAO is used to access the Pond database
	type PondsDAO struct {
		Database string
	}

	var Db *mongo.Collection

	// Connect starts a connection with the mongo db and returns a pointer to it
	func Connect() *mongo.Collection {
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

		// var collection = client.Database(constants.DbName).Collection(constants.CollectionName)
		Db = client.Database(constants.DbName).Collection(constants.CollectionName)

		return Db
	}

	// FindPond returns all ponds from mongodb database
	func FindPond() types.Pond {
		// FIND A VALUE
		filter := bson.D{{"name", "Golden Pond"}}

		update := bson.D{
			{"$inc", bson.D{
				{"age", 9},
			}},
		}
		updateResult, err := Db.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

		// create a value into which the result can be decoded
		var result types.Pond

		err = Db.FindOne(context.TODO(), filter).Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		return result
	}

	// FindAll queries the mongodb 
	func FindAll() []*types.Pond {
		var results []*types.Pond

		// Add extra options to queries using the options package
		options := options.Find()
		options.SetLimit(200)

		// Passing nil as the filter matches all documents in the collection
		cur, err := Db.Find(context.TODO(), nil, options)
		if err != nil {
			log.Fatal(err)
		}

		// Finding multiple documents returns a cursor
		// Iterating through the cursor allows us to decode documents one at a time
		for cur.Next(context.TODO()) {

			// create a value into which the single document can be decoded
			var elem types.Pond
			err := cur.Decode(&elem)
			if err != nil {
				log.Fatal(err)
			}

			results = append(results, &elem)
		}

		if err := cur.Err(); err != nil {
			log.Fatal(err)
		}

		// Close the cursor once finished
		cur.Close(context.TODO())

		fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
		return results
}

	func DeleteAllRecords() {
		deleteResult, err := Db.DeleteMany(context.TODO(), nil)
		fmt.Println(err)
		// if err := cur.Err(); err != nil {
		// 	log.Fatal(err)
		// }
		fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
	}