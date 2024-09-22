package services

import (
	"context"

	"github.com/AkifhanIlgaz/dictionary-api/models"
	"github.com/AkifhanIlgaz/dictionary-api/utils/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type BoxService struct {
	ctx        context.Context
	collection *mongo.Collection
}

func NewBoxService(ctx context.Context, mongoDatabase *mongo.Database) BoxService {
	return BoxService{
		ctx:        ctx,
		collection: mongoDatabase.Collection(db.BoxLogsCollection),
	}
}

func (service BoxService) Action(req models.BoxActionRequest) error {
	// TODO: Normalize request
	return nil
}

func (service BoxService) levelUp(boxAction models.BoxAction) error {
	return nil
}

func (service BoxService) levelDown(boxAction models.BoxAction) error {
	return nil
}

func (service BoxService) levelReset(boxAction models.BoxAction) error {
	return nil
}
