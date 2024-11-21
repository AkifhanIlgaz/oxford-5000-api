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

// @title           Dictionary API
// @version         1.0
// @description     A Dictionary API with authentication and word management features
// @termsOfService  http://swagger.io/terms/

// @contact.name   Your Name
// @contact.url    http://your-url.com
// @contact.email  your-email@domain.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
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
	wordMiddleware := middlewares.NewWordMiddleware(wordService, userService, authService)

	wordController := controllers.NewWordController(wordService, wordMiddleware)
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
