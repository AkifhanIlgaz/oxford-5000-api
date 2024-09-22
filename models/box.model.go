package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BoxLog struct {
	Id         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	WordId     primitive.ObjectID `json:"wordId" bson:"wordId"`
	Level      int                `json:"level" bson:"level"`
	LastUpdate time.Time          `json:"lastUpdate" bson:"lastUpdate"`
	NextRepeat time.Time          `json:"nextRepeat" bson:"nextRepeat"`
}
