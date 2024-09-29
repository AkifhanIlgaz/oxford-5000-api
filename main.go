package main

import (
	"context"
	"log"

	"github.com/AkifhanIlgaz/dictionary-api/config"
	"github.com/AkifhanIlgaz/dictionary-api/controllers"
	"github.com/AkifhanIlgaz/dictionary-api/middlewares"
	"github.com/AkifhanIlgaz/dictionary-api/services"
	"github.com/AkifhanIlgaz/dictionary-api/utils/db"
	"github.com/AkifhanIlgaz/dictionary-api/utils/firebase"
	"github.com/gin-gonic/gin"
)

func main() {
	config, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.TODO()

	mongoClient, err := db.Connect(ctx, config)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = mongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	auth, err := firebase.Auth(ctx, config)
	if err != nil {
		log.Fatal("auth", err)
	}

	_ = auth

	mongoDatabase := mongoClient.Database(db.DatabaseName)

	wordService := services.NewWordService(ctx, mongoDatabase)
	boxService := services.NewBoxService(ctx, mongoDatabase)
	userService := services.NewUserService(ctx, auth)

	wordController := controllers.NewWordController(wordService)
	boxController := controllers.NewBoxController(boxService)

	userMiddleware := middlewares.NewUserMiddleware(userService)

	server := gin.Default()

	router := server.Group("/api")

	// TODO: Add option functions to restrict access to the endpoints ?
	router.Use(userMiddleware.GetUserFromIdToken())

	wordController.SetupRoutes(router)
	boxController.SetupRoutes(router)

	err = server.Run(":" + config.Port)
	if err != nil {
		log.Fatal(err)
	}
}
