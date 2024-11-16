package services

import (
	"context"
	"fmt"

	"github.com/AkifhanIlgaz/dictionary-api/utils/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserService struct {
	ctx        context.Context
	collection *mongo.Collection
}

func NewUserService(ctx context.Context, mongoDatabase *mongo.Database) (UserService, error) {
	collection := mongoDatabase.Collection(db.ApiKeysCollection)

	_, err := collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "uid", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return UserService{}, fmt.Errorf("initialize user service: %w", err)
	}

	return UserService{
		ctx:        ctx,
		collection: collection,
	}, nil
}
