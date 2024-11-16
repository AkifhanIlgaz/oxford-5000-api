package main

import (
	"context"
	"log"

	"github.com/AkifhanIlgaz/dictionary-api/config"
	"github.com/AkifhanIlgaz/dictionary-api/controllers"
	"github.com/AkifhanIlgaz/dictionary-api/middlewares"
	"github.com/AkifhanIlgaz/dictionary-api/services"
	"github.com/AkifhanIlgaz/dictionary-api/utils/db"
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

	mongoDatabase := mongoClient.Database(db.DatabaseName)

	wordService := services.NewWordService(ctx, mongoDatabase)
	tokenService := services.NewTokenService(config)

	userService, err := services.NewUserService(ctx, mongoDatabase)
	if err != nil {
		log.Fatal(err)
	}

	authService, err := services.NewAuthService(ctx, mongoDatabase)
	if err != nil {
		log.Fatal(err)
	}

	wordController := controllers.NewWordController(wordService)
	authController := controllers.NewAuthController(authService, tokenService)
	userController := controllers.NewUserController(userService)
	userMiddleware := middlewares.NewUserMiddleware(userService)

	_ = userMiddleware

	server := gin.Default()

	router := server.Group("/api")

	wordController.SetupRoutes(router)
	authController.SetupRoutes(router)
	userController.SetupRoutes(router)

	err = server.Run(":" + config.Port)
	if err != nil {
		log.Fatal(err)
	}
}
