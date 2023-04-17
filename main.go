package main

import (
	"fmt"
	"time"

	"github.com/AkifhanIlgaz/oxford-5000-api/mongo"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	mongoClient := mongo.ConnectToMongo()

	words := mongoClient.GetWordsCollection()

	noConnectionstart := time.Now()

	_, err := mongo.GetWord(words, "abandon", "verb")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(time.Since(noConnectionstart))

	withConnectionstart := time.Now()

	_, err = mongo.GetWord(words, "abandon", "verb")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(time.Since(withConnectionstart))

}
