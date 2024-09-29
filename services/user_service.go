package services

import (
	"context"
	"fmt"

	"firebase.google.com/go/v4/auth"
)

// ? Firebase'deki user Mongo'ya eklenebilir
type UserService struct {
	ctx  context.Context
	auth *auth.Client
}

func NewUserService(ctx context.Context, auth *auth.Client) UserService {
	return UserService{
		ctx:  ctx,
		auth: auth,
	}
}

func (service UserService) GetUserFromIdToken(idToken string) (*auth.UserRecord, error) {
	token, err := service.auth.VerifyIDToken(context.TODO(), idToken)
	if err != nil {
		return nil, fmt.Errorf("get user from id token: %w", err)
	}

	user, err := service.auth.GetUser(context.Background(), token.UID)
	if err != nil {
		return nil, fmt.Errorf("get user from id token: %w", err)
	}

	return user, nil
}
