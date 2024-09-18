package main

import (
	"context"
	"fmt"
	"log"

	"github.com/AkifhanIlgaz/dictionary-api/config"
	"github.com/AkifhanIlgaz/dictionary-api/utils/db"
	"github.com/AkifhanIlgaz/dictionary-api/utils/firebase"
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

	user, err := auth.GetUser(ctx, "bmK5ommUPRRyxVsjKEpFD0KQlKp2")
	if err != nil {
		log.Fatal("user", err)
	}

	fmt.Println(user)

}
