package dao

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"log"

	models "github.com/bauidch/hyrt-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB connection string
// for localhost mongoDB
// const connectionString = "mongodb://localhost:27017"
const connectionString = "mongodb://localhost:27017"

// Database Name
const dbName = "hyres"
// Collection name
const collName = "series"

// collection object/instance
var collection *mongo.Collection

// create connection with mongo db
func init() {

	// Set client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection = client.Database(dbName).Collection(collName)

	fmt.Println("Collection instance created!")
}

func GetAllSeries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := getAllSeries()
	json.NewEncoder(w).Encode(payload)
}

func CreateSerie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var task models.Series
	_ = json.NewDecoder(r.Body).Decode(&task)
	fmt.Println(task, r.Body)
	insertOneSerie(task)
	json.NewEncoder(w).Encode(task)
}

func GetOneSeries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	payload, e := getOneSeries(params["id"])
	if e != nil {
		log.Fatal(e)
	}
	json.NewEncoder(w).Encode(payload)
}
// get all task from the DB and return it
func getAllSeries() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Println("Error: getAllSeries", err)
	}

	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			log.Println("Error: getAllSeries", e)
		}
		//fmt.Println("cur..>", cur, "result", reflect.TypeOf(result), reflect.TypeOf(result["_id"]))
		results = append(results, result)

	}

	if err := cur.Err(); err != nil {
		log.Println("Error: getAllSeries", err)
	}

	cur.Close(context.Background())
	return results

}


func insertOneSerie(serie models.Series) {
	insertResult, err := collection.InsertOne(context.Background(), serie)

	if err != nil {

	}

	fmt.Println("Inserted a Single Record ", insertResult.InsertedID)
}

func getOneSeries(input string) (*models.Series, error){
	id, e := primitive.ObjectIDFromHex(input)
	if e != nil {
		log.Fatal(e)
	}

	var result models.Series
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&result)
	if err != nil {
		log.Println("Error: getOneSeries", err)
	}

	return &result, nil
}
