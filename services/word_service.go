package services

import (
	"context"
	"fmt"

	"github.com/AkifhanIlgaz/dictionary-api/models"
	"github.com/AkifhanIlgaz/dictionary-api/utils/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// WordService handles operations related to word data in MongoDB
type WordService struct {
	ctx        context.Context
	collection *mongo.Collection
}

// NewWordService creates and returns a new WordService instance
// with the provided context and MongoDB database connection
func NewWordService(ctx context.Context, mongoDatabase *mongo.Database) WordService {
	return WordService{
		ctx:        ctx,
		collection: mongoDatabase.Collection(db.WordsCollection),
	}
}

// GetById retrieves a word from the database by its ID
// Returns the word information and any error encountered
// If the ID is invalid or the word is not found, returns an error
func (service WordService) GetById(wordId string) (models.WordInfo, error) {
	id, err := primitive.ObjectIDFromHex(wordId)
	if err != nil {
		return models.WordInfo{}, fmt.Errorf("get ad by id: %w", err)
	}

	filter := bson.M{
		"_id": id,
	}

	var word models.WordInfo

	err = service.collection.FindOne(service.ctx, filter).Decode(&word)
	if err != nil {
		return models.WordInfo{}, fmt.Errorf("get word by id: %w", err)
	}

	return word, nil
}

// GetByName retrieves a word from the database by its name and optionally its part of speech
// If partOfSpeech is empty, it will search only by word name
// Returns the word information and any error encountered
// If the word is not found, returns an error
func (service WordService) GetByName(wordName string, partOfSpeech string) (models.WordInfo, error) {
	filter := bson.M{
		"word": wordName,
	}

	if partOfSpeech != "" {
		filter["header.partOfSpeech"] = partOfSpeech
	}

	var word models.WordInfo

	err := service.collection.FindOne(service.ctx, filter).Decode(&word)
	if err != nil {
		return models.WordInfo{}, fmt.Errorf("get word by name: %w", err)
	}

	return word, nil
}
