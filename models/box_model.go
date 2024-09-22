package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TODO: LastAction için enum kullan
type BoxAction struct {
	Id         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Uid        primitive.ObjectID `json:"uid" bson:"uid"`
	WordId     primitive.ObjectID `json:"wordId" bson:"wordId"`
	Level      int                `json:"level" bson:"level"`
	LastAction string             `json:"lastAction" bson:"lastAction"`
	LastUpdate time.Time          `json:"lastUpdate" bson:"lastUpdate"`
	NextRepeat time.Time          `json:"nextRepeat" bson:"nextRepeat"`
}

type BoxActionRequest struct {
	Uid        string `json:"uid"`
	WordId     string `json:"wordId" binding:"required"`
	ActionName string `json:"actionName" binding:"required"` // Up Down Reset
}
