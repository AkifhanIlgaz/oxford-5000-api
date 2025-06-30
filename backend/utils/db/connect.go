package db

import (
	"context"
	"fmt"

	"github.com/AkifhanIlgaz/dictionary-api/config"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(ctx context.Context, config config.Config) (*mongo.Client, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(config.MongoURI).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	if err := client.Database(DatabaseName).RunCommand(ctx, bson.M{"ping": 1}).Err(); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return client, nil
}

func ConnectToRedis(config config.Config) (*redis.Client, error) {
	opt, err := redis.ParseURL(config.RedisConnectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Redis connection string: %w", err)
	}

	client := redis.NewClient(opt)

	// Test the connection
	if err := client.Ping(client.Context()).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	fmt.Println("Connected to Redis successfully!")
	return client, nil
}
