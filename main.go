package main

import (
	"context"
	"log"

	"github.com/AkifhanIlgaz/dictionary-api/config"
	"github.com/AkifhanIlgaz/dictionary-api/controllers"
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

	// TODO: Initialize services with indexes ?

	wordController := controllers.NewWordController(wordService)
	boxController := controllers.NewBoxController(boxService)

	server := gin.Default()

	router := server.Group("/api")

	wordController.SetupRoutes(router)
	boxController.SetupRoutes(router)

	// TODO: Get port from Config file Dev - Prod

	err = server.Run(":" + config.Port)
	if err != nil {
		log.Fatal(err)
	}

}
