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
	ctx              context.Context
	userCollection   *mongo.Collection
	apiKeyCollection *mongo.Collection
	usageCollection  *mongo.Collection
}

func NewUserService(ctx context.Context, mongoDatabase *mongo.Database) (UserService, error) {
	userCollection := mongoDatabase.Collection(db.UsersCollection)
	apiKeyCollection := mongoDatabase.Collection(db.ApiKeysCollection)
	usageCollection := mongoDatabase.Collection(db.DailyUsageCollection)

	_, err := userCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return UserService{}, fmt.Errorf("initialize user service: %w", err)
	}

	// Create index for api keys collection
	_, err = apiKeyCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
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
		ctx:              ctx,
		userCollection:   userCollection,
		apiKeyCollection: apiKeyCollection,
		usageCollection:  usageCollection,
	}, nil
}

// CreateApiKey generates and stores a new API key for the given user ID
func (s *UserService) CreateApiKey(uid string) (*models.APIKey, error) {
	apiKey, err := apikey.GenerateAPIKey()
	if err != nil {
		return nil, fmt.Errorf("generate api key: %w", err)
	}

	apiKeyDoc := models.APIKey{
		Uid:        uid,
		Key:        apiKey,
		TotalUsage: 0,
		CreatedAt:  time.Now(),
	}

	_, err = s.apiKeyCollection.InsertOne(s.ctx, apiKeyDoc)
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
	result, err := s.apiKeyCollection.DeleteOne(s.ctx, filter)
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

	err := s.apiKeyCollection.FindOne(s.ctx, filter).Decode(&apiKey)
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
	err := s.apiKeyCollection.FindOne(s.ctx, filter).Decode(&apiKey)
	if err != nil {
		return nil, fmt.Errorf("get by api key: %w", err)
	}

	return &apiKey, nil
}

func (s *UserService) IncrementUsage(uid, key string) (int, error) {
	dailyUsage, err := s.incrementDailyUsage(uid, key)
	if err != nil {
		return 0, fmt.Errorf("increment daily usage: %w", err)
	}

	err = s.incrementTotalUsage(key)
	if err != nil {
		return 0, fmt.Errorf("increment total usage: %w", err)
	}

	return dailyUsage, nil
}

func (s *UserService) incrementDailyUsage(uid, 	key string) (int, error) {
	// Get today's date at midnight (00:00:00)
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	filter := bson.M{
		"key":  key,
		"uid":  uid,
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

	_, err := s.apiKeyCollection.UpdateOne(s.ctx, filter, update)
	if err != nil {
		return fmt.Errorf("increment total usage: %w", err)
	}

	return nil
}

func (s *UserService) GetTodayUsage(uid string) (int, error) {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	filter := bson.M{
		"uid":  uid,
		"date": today,
	}

	var result models.DailyUsageEntry
	err := s.usageCollection.FindOne(s.ctx, filter).Decode(&result)
	if err != nil {
		return 0, fmt.Errorf("get today usage: %w", err)
	}

	return result.Usage, nil
}

func (s *UserService) GetTotalUsage(uid string) (int, error) {
	filter := bson.M{
		"uid": uid,
	}

	var result models.APIKey
	err := s.apiKeyCollection.FindOne(s.ctx, filter).Decode(&result)
	if err != nil {
		return 0, fmt.Errorf("get total usage: %w", err)
	}

	return result.TotalUsage, nil
}
