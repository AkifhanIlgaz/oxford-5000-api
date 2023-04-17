package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Word struct {
	Word         string    `json:"Word"`
	PartOfSpeech string    `json:"PartOfSpeech"`
	CEFRLevel    string    `json:"CEFRLevel"`
	Meanings     []Meaning `json:"Meanings"`
	Idioms       []Idiom   `json:"Idioms"`
	Url          string    `json:"URL"`
}

type Meaning struct {
	CF         string   `json:"cf"`
	Definition string   `json:"Definition"`
	Examples   []string `json:"Examples"`
}

type Idiom struct {
	Idm        string   `json:"idm"`
	Definition string   `json:"def"`
	Examples   []string `json:"examples"`
}

// Some words have same name but different meanings or they are different part of speech
// For example, there are two entry for "about" in our database. One for preposition and one for adverb
// So, this function returns all entries that matches the given word

// To differentiate between the entries, use parameters on api
// For example, /word/about?partOfSpeech=preposition

func GetWord(collection *mongo.Collection, word string, partOfSpeech ...string) (Word, error) {
	var result Word
	filter := bson.M{"word": word}

	if len(partOfSpeech) != 1 {
		filter = bson.M{"word": word, "partofspeech": partOfSpeech[0]}
	} else {
		if len(partOfSpeech) > 1 {
			return result, fmt.Errorf("please specify only one part of speech")
		} else {
			return result, fmt.Errorf("please specify the part of speech")
		}
	}

	r := collection.FindOne(context.Background(), filter)
	if r.Err() != nil {
		return result, r.Err()
	}

	r.Decode(&result)

	return result, nil
}

func AddWord(collection *mongo.Collection, word Word) {

	// Add error handling
	// If the word already exists mongo creates another one
	// Check word name

	result, err := collection.InsertOne(context.Background(), word)

	if err != nil {
		fmt.Println("Couldn't add the word to the database")
	}

	fmt.Println(result)
}

func DeleteWord(collection *mongo.Collection, word string) {
	result, err := collection.DeleteOne(context.Background(), bson.M{"word": word})

	if err != nil {
		fmt.Println("Couldn't delete the word from the database")
	}

	fmt.Println(result)
}

//	 TODO:
//		- Check if the word exists
//		- If it does, delete it
//		- Add the new word
//
// Functional option pattern
func UpdateWord(collection *mongo.Collection, word string, newWord Word) {
	DeleteWord(collection, word)
	AddWord(collection, newWord)
}
