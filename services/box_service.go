package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/AkifhanIlgaz/dictionary-api/models"
	"github.com/AkifhanIlgaz/dictionary-api/utils/db"
	"github.com/AkifhanIlgaz/dictionary-api/utils/message"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MaxLevel int = 7
	MinLevel int = 1
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
	action, err := service.GetWordAction(req.Uid, req.WordId)
	if err != nil {
		return fmt.Errorf("action: %w", err)
	}

	switch req.ActionName {
	case "level-up":
		return service.levelUp(action)
	case "level-down":
		return service.levelDown(action)
	case "reset":
		return service.levelReset(action)
	default:
		return errors.New(message.UnsupportedActionType)
	}
}

// TODO: Aggregate
func (service BoxService) GetWordAction(uid, wordId string) (models.BoxAction, error) {
	var action models.BoxAction

	filter := bson.M{
		"uid":    uid,
		"wordId": wordId,
	}

	err := service.collection.FindOne(service.ctx, filter).Decode(&action)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			act := models.BoxAction{
				Uid:        uid,
				WordId:     wordId,
				Level:      0,
				LastAction: "init",
				LastUpdate: time.Now(),
				NextRepeat: time.Now(),
			}

			err := service.initAction(act)
			if err != nil {
				return models.BoxAction{}, fmt.Errorf("get word action: %w", err)
			}

			return act, nil
		}
		return models.BoxAction{}, fmt.Errorf("get word action: %w", err)
	}

	return action, nil
}

func (service BoxService) initAction(boxAction models.BoxAction) error {
	_, err := service.collection.InsertOne(service.ctx, boxAction)
	if err != nil {
		return fmt.Errorf("init action: %w", err)
	}

	return nil
}

func (service BoxService) levelUp(boxAction models.BoxAction) error {
	boxAction.LastAction = "level-up"
	boxAction.LastUpdate = time.Now()
	boxAction.Level = incrementLevel(boxAction.Level)
	boxAction.NextRepeat = getNextRepeat(boxAction.Level)

	filter := bson.M{
		"_id": boxAction.Id,
	}

	res, err := service.collection.ReplaceOne(service.ctx, filter, boxAction)
	if err != nil {
		return fmt.Errorf("level up: %w", err)
	}

	if res.ModifiedCount == 0 {
		return fmt.Errorf("level up: %w", err)
	}

	return nil

}

func (service BoxService) levelDown(boxAction models.BoxAction) error {
	boxAction.LastAction = "level-down"
	boxAction.LastUpdate = time.Now()
	boxAction.Level = decrementLevel(boxAction.Level)
	boxAction.NextRepeat = getNextRepeat(boxAction.Level)

	filter := bson.M{
		"_id": boxAction.Id,
	}

	res, err := service.collection.ReplaceOne(service.ctx, filter, boxAction)
	if err != nil {
		return fmt.Errorf("level down: %w", err)
	}

	if res.ModifiedCount == 0 {
		return fmt.Errorf("level down: %w", err)
	}

	return nil
}

func (service BoxService) levelReset(boxAction models.BoxAction) error {
	boxAction.LastAction = "reset"
	boxAction.LastUpdate = time.Now()
	boxAction.Level = 0
	boxAction.NextRepeat = time.Now().Add(24 * time.Hour)

	filter := bson.M{
		"_id": boxAction.Id,
	}

	res, err := service.collection.ReplaceOne(service.ctx, filter, boxAction)
	if err != nil {
		return fmt.Errorf("reset: %w", err)
	}

	if res.ModifiedCount == 0 {
		return fmt.Errorf("reset: %w", err)
	}

	return nil
}

func incrementLevel(level int) int {
	if level+1 > MaxLevel {
		return level
	}

	return level + 1
}

func decrementLevel(level int) int {
	if level-1 < MinLevel {
		return level
	}

	return level - 1
}

const (
	Day   time.Duration = 24 * time.Hour
	Week  time.Duration = 7 * Day
	Month time.Duration = 4 * Week
	Year  time.Duration = 12 * Month
)

const (
	IntervalLevel1 time.Duration = 1 * Day
	IntervalLevel2 time.Duration = 3 * Day
	IntervalLevel3 time.Duration = 1 * Week
	IntervalLevel4 time.Duration = 2 * Week
	IntervalLevel5 time.Duration = 1 * Month
	IntervalLevel6 time.Duration = 3 * Month
	IntervalLevel7 time.Duration = 1 * Year
)

func getNextRepeat(level int) time.Time {
	// TODO: Return based on oldNextRepeat. Eğer belli bir süre geçmişse time.Now'a göre yap

	switch level {
	case 1:
		return time.Now().Add(IntervalLevel1)
	case 2:
		return time.Now().Add(IntervalLevel2)
	case 3:
		return time.Now().Add(IntervalLevel3)
	case 4:
		return time.Now().Add(IntervalLevel4)
	case 5:
		return time.Now().Add(IntervalLevel5)
	case 6:
		return time.Now().Add(IntervalLevel6)
	case 7:
		return time.Now().Add(IntervalLevel7)
	default:
		return time.Now().Add(IntervalLevel7)
	}
}
