package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/AkifhanIlgaz/dictionary-api/models"
	"github.com/AkifhanIlgaz/dictionary-api/utils/crypto"
	"github.com/AkifhanIlgaz/dictionary-api/utils/db"
	"github.com/AkifhanIlgaz/dictionary-api/utils/message"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// AuthService handles user authentication operations using MongoDB
type AuthService struct {
	ctx        context.Context
	collection *mongo.Collection
}

// NewAuthService creates a new AuthService instance and initializes a unique index on the email field.
// It returns an error if the index creation fails.
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
		ctx:        ctx,
		collection: collection,
	}, nil
}

// Create registers a new user in the system.
// It takes an AuthRequest containing email and password, hashes the password,
// and stores the user in MongoDB.
// Returns the created user's ObjectID or an error if the operation fails.
func (service AuthService) Create(req models.AuthRequest) (primitive.ObjectID, error) {
	passwordHash, err := crypto.HashPassword(req.Password)
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("create user: %w", err)
	}

	userToCreate := models.User{
		Email:        req.Email,
		Plan:         "free",
		PasswordHash: passwordHash,
	}

	result, err := service.collection.InsertOne(service.ctx, userToCreate)
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("create user: %w", err)
	}

	return result.InsertedID.(primitive.ObjectID), nil
}

// AuthenticateUser verifies user credentials against stored data.
// It takes an AuthRequest containing email and password, finds the user by email,
// and verifies the password hash.
// Returns the user's ObjectID if authentication succeeds, or an error if credentials are invalid.
func (service AuthService) AuthenticateUser(req models.AuthRequest) (primitive.ObjectID, error) {

	filter := bson.M{
		"email": req.Email,
	}

	var user models.User

	err := service.collection.FindOne(service.ctx, filter).Decode(&user)
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("authenticate user: %w", err)
	}

	if err := crypto.VerifyPassword(user.PasswordHash, req.Password); err != nil {
		return primitive.NilObjectID, errors.New(message.WrongPassword)
	}

	return user.Id, nil
}

// GetUserPlan retrieves the subscription plan for a given user ID.
// Returns the plan type (e.g. "free", "pro") or an error if the user is not found.
func (service AuthService) GetUserPlan(uid string) (string, error) {
	objectId, err := primitive.ObjectIDFromHex(uid)
	if err != nil {
		return "", fmt.Errorf("invalid user id: %w", err)
	}

	filter := bson.M{
		"_id": objectId,
	}

	var user models.User
	err = service.collection.FindOne(service.ctx, filter).Decode(&user)
	if err != nil {
		return "", fmt.Errorf("get user plan: %w", err)
	}

	return user.Plan, nil
}
