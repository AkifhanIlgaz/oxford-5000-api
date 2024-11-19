package models

import "time"

type APIKey struct {
	Key        string    `json:"key" bson:"key"`
	Uid        string    `json:"uid" bson:"uid"`
	Name       string    `json:"name" bson:"name"`
	TotalUsage int       `json:"totalUsage" bson:"totalUsage"`
	CreatedAt  time.Time `json:"createdAt" bson:"createdAt"`
}

type APIKeyResponse struct {
	APIKey *APIKey `json:"apiKey"`
}
