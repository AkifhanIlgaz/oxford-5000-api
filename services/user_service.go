package services

import (
	"context"
	"fmt"
	"time"

	"github.com/AkifhanIlgaz/dictionary-api/models"
	"github.com/AkifhanIlgaz/dictionary-api/utils/apikey"
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

// CreateApiKey generates and stores a new API key for the given user ID
func (s *UserService) CreateApiKey(uid, name string) (*models.APIKey, error) {
	apiKey, err := apikey.GenerateApiKey()
	if err != nil {
		return nil, fmt.Errorf("generate api key: %w", err)
	}

	apiKeyDoc := models.APIKey{
		Uid:        uid,
		Name:       name,
		Key:        apiKey,
		TotalUsage: 0,
		CreatedAt:  time.Now(),
	}

	_, err = s.collection.InsertOne(s.ctx, apiKeyDoc)
	if err != nil {
		return nil, fmt.Errorf("create api key: %w", err)
	}

	return &apiKeyDoc, nil
}

// DeleteApiKey removes the API key associated with the given user ID
func (s *UserService) DeleteApiKey(uid string) error {
	filter := bson.M{
		"uid": uid,
	}
	result, err := s.collection.DeleteOne(s.ctx, filter)
	if err != nil {
		return fmt.Errorf("delete api key: %w", err)
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("no api key found for user: %s", uid)
	}

	return nil
}

// GetApiKey retrieves the API key associated with the given user ID
func (s *UserService) GetApiKey(uid string) (*models.APIKey, error) {
	var apiKey models.APIKey

	filter := bson.M{
		"uid": uid,
	}

	err := s.collection.FindOne(s.ctx, filter).Decode(&apiKey)
	if err != nil {
		return nil, fmt.Errorf("get api key: %w", err)
	}

	return &apiKey, nil
}
