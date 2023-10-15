package contollers

//Imp imports
import (
	"context"
	"fmt"
	"log"

	"github.com/kaiwalyakoparkar/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Variable options to use later in the database connection
const connectionString = "mongodb+srv://kaiwalya:trial@cluster0.ku5a95e.mongodb.net/?retryWrites=true&w=majority"
const dbname = "netflix"
const colname = "watched"

var collection *mongo.Collection

//Method to connect and initiate the database
func init() {
	//Attaching the stream to the options
	clientOption := options.Client().ApplyURI(connectionString)
	//Connecting the database
	client, err := mongo.Connect(context.TODO(),clientOption)

	//Handling the connection error if any
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB Connection Successful")

	//Creating a collection on the database
	collection = client.Database(dbname).Collection(colname)

	fmt.Println("Collection instance ready")
}

//Mongo helper used to just push the data to mongodb database
func insertOneMovie(movie models.Netflix) {
	//Inserting the data into the database
	inserted, err := collection.InsertOne(context.Background(), movie)

	//Handling the error for the failed insertion
	if err != nil {
		log.Fatal(err)
	}

	//Printing out the id of the inserted data (id comes from mongodb)
	fmt.Println("Inserter one movie in DB with id", inserted.InsertedID)
}

//Mongo helper to update the data in database (change watched to true/false)
func updateOneMovie(movieId string) {

	//Extracting the id out of the string
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}

	//Constructing a filter/rule to find the record in the database
	filter := bson.M{"_id": id}
	//Create a new updated object using bson
	updated := bson.M{"$set": bson.M{"Watched": true}}
	//Execute the updation
	result, err := collection.UpdateOne(context.Background(), filter, updated)

	if err != nil {
		log.Fatal(err)
	}

	//Print success message along with the count of records modified
	fmt.Println("Movie watched ", result.ModifiedCount)
}

//Mongo helper to delete one movie data in database
func deleteOneMovie(movieId string) {

	//Extracting the id out of the string
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}
	
	//Constructing a filter/rule to find the record in database
	filter := bson.M{"_id": id}

	//Performing deletion
	deleted, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	//Print success message along with the count of records deleted
	fmt.Println("Movie record delete successfully", deleted.DeletedCount)
}

//Mongo helper to delete all movie data in database
func deleteAllMovie() int64 {
	//Create an empty filter to select every record in the database
	filter := bson.D{{}}
	//Execute the delete opteration
	deleted, err := collection.DeleteMany(context.Background(), filter, nil)

	//Log the errors if any
	if err != nil {
		log.Fatal(err)
	}

	//return the delete count
	return deleted.DeletedCount
}