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

type WordService struct {
	ctx        context.Context
	collection *mongo.Collection
}

func NewWordService(ctx context.Context, mongoDatabase *mongo.Database) WordService {
	return WordService{
		ctx:        ctx,
		collection: mongoDatabase.Collection(db.WordsCollection),
	}
}

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

func (service WordService) GetByIndex(wordIndex int) (models.WordInfo, error) {

	filter := bson.M{
		"index": wordIndex,
	}

	var word models.WordInfo

	err := service.collection.FindOne(service.ctx, filter).Decode(&word)
	if err != nil {
		return models.WordInfo{}, fmt.Errorf("get word by index: %w", err)
	}

	return word, nil
}
