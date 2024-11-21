package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email        string             `json:"email" bson:"email"`
	Plan         string             `json:"plan" bson:"plan"`
	PasswordHash string             `json:"-" bson:"passwordHash"`
}
