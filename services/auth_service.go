package services

import (
	"context"
	"fmt"

	"github.com/AkifhanIlgaz/dictionary-api/models"
	"github.com/AkifhanIlgaz/dictionary-api/utils/crypto"
	"github.com/AkifhanIlgaz/dictionary-api/utils/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AuthService struct {
	ctx context.Context
	db  *mongo.Database
}

func NewAuthService(ctx context.Context, mongodb *mongo.Database) (AuthService, error) {
	collection := mongodb.Collection(db.UsersCollection)

	_, err := collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return AuthService{}, fmt.Errorf("initialize auth service: %w", err)
	}

	return AuthService{
		ctx: ctx,
		db:  mongodb,
	}, nil
}

func (service AuthService) Create(req models.SignupRequest) (primitive.ObjectID, error) {
	passwordHash, err := crypto.HashPassword(req.Password)
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("create user: %w", err)
	}

	userToCreate := models.User{
		Email:        req.Email,
		PasswordHash: passwordHash,
	}

	collection := service.db.Collection(db.UsersCollection)

	result, err := collection.InsertOne(service.ctx, userToCreate)
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("create user: %w", err)
	}

	return result.InsertedID.(primitive.ObjectID), nil
}