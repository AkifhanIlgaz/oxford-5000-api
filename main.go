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
	userService := services.NewUserService(ctx)

	wordController := controllers.NewWordController(wordService)
	userController := controllers.NewUserController(userService)
	userMiddleware := middlewares.NewUserMiddleware(userService)

	_ = userMiddleware

	server := gin.Default()

	router := server.Group("/api")

	wordController.SetupRoutes(router)

	err = server.Run(":" + config.Port)
	if err != nil {
		log.Fatal(err)
	}
}
