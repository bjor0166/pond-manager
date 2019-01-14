package types

//import "github.com/mongodb/mongo-go-driver/bson"

// Pond defines a Pond type
type Pond struct {
	// ID	bson.ObjectId `bson:"_id" json:"id"`
	ID	string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}