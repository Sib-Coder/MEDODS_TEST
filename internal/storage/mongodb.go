package storage

import (
	"awesomeProject/internal/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var (
	mongoURI = "mongodb://root:rootpassword@localhost:27017"
)

type Database struct {
	Db *mongo.Client
}

func New() *Database {
	clientOptions := options.Client().ApplyURI(mongoURI) // Connect to //MongoDB
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
	return &Database{
		Db: client,
	}
}

func (db *Database) SelectInfoUser(idx string) (model.Traning, error) {
	var result model.Traning
	collection := db.Db.Database("modods").Collection("modods")

	// create filter for document to fund
	id, err := primitive.ObjectIDFromHex(idx)
	if err != nil {
		fmt.Println(err)
		return model.Traning{}, err
	}
	filter := bson.M{"_id": id}

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
		return model.Traning{}, err
	}

	return result, err
}

func (db *Database) UpdateRefresh(idx string, reftoken string) (bool, error) {
	collection := db.Db.Database("modods").Collection("modods")

	fmt.Println(idx)
	id, err := primitive.ObjectIDFromHex(idx)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	filter := bson.M{"_id": id}

	// create update to add funds to document
	update := bson.M{"$set": bson.M{"Refresh": reftoken}}

	// update document in collection
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	// print number of documents updated
	fmt.Printf("Updated %v document\n", result.ModifiedCount)
	return true, err

}
