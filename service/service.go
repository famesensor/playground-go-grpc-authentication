package service

import (
	"github.com/famesensor/ghelper"
	"github.com/famesensor/playground-go-grpc-authentication/port"
	"github.com/famesensor/playground-go-grpc-authentication/service/auth"
	"github.com/famesensor/playground-go-grpc-authentication/service/user"
)

type service struct {
	AuthService auth.Port
	UserService user.Port
}

func NewService(database port.DatabasePort, uuid ghelper.UUID, jwtManager ghelper.JWTManager) *service {
	return &service{
		AuthService: auth.NewService(database, uuid, jwtManager),
		UserService: user.NewService(database),
	}
}
