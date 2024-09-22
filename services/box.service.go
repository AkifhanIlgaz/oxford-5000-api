package services

import (
	"context"

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
