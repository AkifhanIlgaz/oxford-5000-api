package main

import (
	"context"
	"log"

	"github.com/AkifhanIlgaz/dictionary-api/config"
	"github.com/AkifhanIlgaz/dictionary-api/controllers"
	_ "github.com/AkifhanIlgaz/dictionary-api/docs"
	"github.com/AkifhanIlgaz/dictionary-api/middlewares"
	"github.com/AkifhanIlgaz/dictionary-api/services"
	"github.com/AkifhanIlgaz/dictionary-api/utils/db"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	userMiddleware := middlewares.NewUserMiddleware(tokenService)

	wordController := controllers.NewWordController(wordService)
	authController := controllers.NewAuthController(authService, tokenService)
	userController := controllers.NewUserController(userService, userMiddleware)

	server := gin.Default()

	router := server.Group("/api")

	wordController.SetupRoutes(router)
	authController.SetupRoutes(router)
	userController.SetupRoutes(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err = server.Run(":" + config.Port)
	if err != nil {
		log.Fatal(err)
	}
}
