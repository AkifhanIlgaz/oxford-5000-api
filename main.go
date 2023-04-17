package main

import (
	"fmt"

	"github.com/AkifhanIlgaz/oxford-5000-api/mongo"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	mongoClient := mongo.ConnectToMongo()

	words := mongoClient.GetWordsCollection()

	a, err := mongo.GetWord(words, "about", "adverb")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(a)
}
