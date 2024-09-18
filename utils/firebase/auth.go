package firebase

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/AkifhanIlgaz/dictionary-api/config"
	"google.golang.org/api/option"
)

func Auth(ctx context.Context, config config.Config) (*auth.Client, error) {
	opt := option.WithCredentialsFile(config.GoogleApplicationCredentials)

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, fmt.Errorf("firebase app: %v\n", err)
	}

	auth, err := app.Auth(ctx)
	if err != nil {
		return nil, fmt.Errorf("firebase auth: %v\n", err)
	}

	return auth, nil
}
