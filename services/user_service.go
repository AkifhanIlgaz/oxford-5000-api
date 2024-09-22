package services

import (
	"firebase.google.com/go/v4/auth"
)

// ? Firebase'deki user Mongo'ya eklenebilir
type UserService struct {
	auth *auth.Client
}

func NewUserService(auth *auth.Client) UserService {
	return UserService{
		auth: auth,
	}
}
