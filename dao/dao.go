package dao

import (
	"context"
	"log"

	models "github.com/bauidch/hyrt-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB connection string
// for localhost mongoDB
const connectionString = "mongodb://localhost:27017"

// Database Name
const dbName = "hyres"

// collection object/instance
var collSeedJournal *mongo.Collection
var collectionSeries *mongo.Collection
var collectionSeed *mongo.Collection

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

	log.Println("Connected to MongoDB!")

	collectionSeries = client.Database(dbName).Collection("series")
	collectionSeed = client.Database(dbName).Collection("seed")
	collSeedJournal = client.Database(dbName).Collection("seed_journal")

	log.Println("Collection instance created!")
}

// Series
// get all task from the DB and return it
func GetAllSeries() []primitive.M {
	cur, err := collectionSeries.Find(context.Background(), bson.D{{}})
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


func InsertOneSerie(serie models.Series) {
	insertResult, err := collectionSeries.InsertOne(context.Background(), serie)

	if err != nil {

	}

	log.Println("Inserted a Single Record ", insertResult.InsertedID)
}

func GetOneSeries(input string) (*models.Series, error){
	id, e := primitive.ObjectIDFromHex(input)
	if e != nil {
		log.Fatal(e)
	}

	var result models.Series
	err := collectionSeries.FindOne(context.Background(), bson.M{"_id": id}).Decode(&result)
	if err != nil {
		log.Println("Error: getOneSeries", err)
	}

	return &result, nil
}

// Seed
func GetAllSeed() []primitive.M {
	cur, err := collectionSeed.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Println("Error: getAllSeed", err)
	}

	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			log.Println("Error: getAllSeed", e)
		}
		//fmt.Println("cur..>", cur, "result", reflect.TypeOf(result), reflect.TypeOf(result["_id"]))
		results = append(results, result)

	}

	if err := cur.Err(); err != nil {
		log.Println("Error: getAllSeed", err)
	}

	cur.Close(context.Background())
	return results

}

func InsertOneSeed(serie models.Seed) {
	insertResult, err := collectionSeed.InsertOne(context.Background(), serie)

	if err != nil {
		log.Println("Error: InsertOneSeed", err)
	}

	log.Println("Inserted a Single Record ", insertResult.InsertedID)
}

// SeedJournal
func GetAllSeedJournal() []primitive.M {
	cur, err := collSeedJournal.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Println("Error: getAllSeed", err)
	}

	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			log.Println("Error: GetAllSeedJournal", e)
		}
		//fmt.Println("cur..>", cur, "result", reflect.TypeOf(result), reflect.TypeOf(result["_id"]))
		results = append(results, result)

	}

	if err := cur.Err(); err != nil {
		log.Println("Error: GetAllSeedJournal", err)
	}

	cur.Close(context.Background())
	return results

}

func InsertSeedJournalEntry(serie models.Seed) {
	insertResult, err := collSeedJournal.InsertOne(context.Background(), serie)

	if err != nil {
		log.Println("Error: InsertSeedJournalEntry", err)
	}

	log.Println("Inserted a Single Record ", insertResult.InsertedID)
}
