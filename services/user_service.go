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
	ctx             context.Context
	collection      *mongo.Collection
	usageCollection *mongo.Collection
	
}

func NewUserService(ctx context.Context, mongoDatabase *mongo.Database) (UserService, error) {
	collection := mongoDatabase.Collection(db.ApiKeysCollection)
	usageCollection := mongoDatabase.Collection(db.DailyUsageCollection)

	// Create index for api keys collection
	_, err := collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "uid", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return UserService{}, fmt.Errorf("initialize user service: %w", err)
	}

	// Create compound index for usage collection
	_, err = usageCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{
			{Key: "key", Value: 1},
			{Key: "date", Value: 1},
		},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return UserService{}, fmt.Errorf("initialize usage collection: %w", err)
	}

	return UserService{
		ctx:             ctx,
		collection:      collection,
		usageCollection: usageCollection,
	}, nil
}

// CreateApiKey generates and stores a new API key for the given user ID
func (s *UserService) CreateApiKey(uid, name string) (*models.APIKey, error) {
	apiKey, err := apikey.GenerateAPIKey()
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

func (s *UserService) GetByKey(key string) (*models.APIKey, error) {
	filter := bson.M{
		"key": key,
	}

	var apiKey models.APIKey
	err := s.collection.FindOne(s.ctx, filter).Decode(&apiKey)
	if err != nil {
		return nil, fmt.Errorf("get by api key: %w", err)
	}

	return &apiKey, nil
}

func (s *UserService) IncrementUsage(key string) (int, error) {
	dailyUsage, err := s.incrementDailyUsage(key)
	if err != nil {
		return 0, fmt.Errorf("increment daily usage: %w", err)
	}

	err = s.incrementTotalUsage(key)
	if err != nil {
		return 0, fmt.Errorf("increment total usage: %w", err)
	}

	return dailyUsage, nil
}

func (s *UserService) incrementDailyUsage(key string) (int, error) {
	// Get today's date at midnight (00:00:00)
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	filter := bson.M{
		"key":  key,
		"date": today,
	}

	update := bson.M{
		"$inc": bson.M{"count": 1},
	}

	opts := options.FindOneAndUpdate().
		SetUpsert(true).
		SetReturnDocument(options.After)

	var result models.DailyUsageEntry
	err := s.usageCollection.FindOneAndUpdate(s.ctx, filter, update, opts).Decode(&result)
	if err != nil {
		return 0, fmt.Errorf("increment daily usage: %w", err)
	}

	return result.Usage, nil
}

func (s *UserService) incrementTotalUsage(key string) error {
	filter := bson.M{
		"key": key,
	}

	update := bson.M{
		"$inc": bson.M{"totalUsage": 1},
	}

	_, err := s.collection.UpdateOne(s.ctx, filter, update)
	if err != nil {
		return fmt.Errorf("increment total usage: %w", err)
	}

	return nil
}
