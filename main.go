package main

import (
	"context"
	"fmt"

	"github.com/AkifhanIlgaz/oxford-5000-api/mongo"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	godotenv.Load()

	mongoClient := mongo.ConnectToMongo()

	words := mongoClient.GetCollection("Oxford5000", "Words")

	// Find a single document from the Words collection whose word is "abandon"
	// var result Word
	// err := words.FindOne(context.Background(), bson.M{"word": "abandon"}).Decode(&result)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	var abandon mongo.Word
	words.FindOne(context.Background(), bson.M{"word": "abandon"}).Decode(&abandon)

	fmt.Println(abandon)
}
