package mongo

import "go.mongodb.org/mongo-driver/mongo"

type Words struct {
	Database   *mongo.Database
	Collection *mongo.Collection
}

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
