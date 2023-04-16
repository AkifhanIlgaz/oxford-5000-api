package mongo

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	Client *mongo.Client
}

func ConnectToMongo() MongoClient {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("MONGO")).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		fmt.Println("Error connecting to MongoDB")
		panic(err)
	}

	fmt.Println("Successfully connected to MongoDB!")

	return MongoClient{Client: client}
}

func (m MongoClient) GetDatabase(databaseName string) *mongo.Database {
	return m.Client.Database(databaseName)
}

func (m MongoClient) GetCollection(databaseName, collectionName string) *mongo.Collection {
	return m.Client.Database(databaseName).Collection(collectionName)
}
