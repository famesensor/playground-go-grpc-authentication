package service

import (
	"github.com/famesensor/playground-go-grpc-authentication/helper"
	"github.com/famesensor/playground-go-grpc-authentication/port"
	"github.com/famesensor/playground-go-grpc-authentication/service/auth"
	"github.com/famesensor/playground-go-grpc-authentication/service/user"
)

type service struct {
	AuthService auth.Port
	UserService user.Port
}

func NewService(database port.DatabasePort, uuid helper.UUID, jwtManager helper.JWTManager) *service {
	return &service{
		AuthService: auth.NewService(database, uuid, jwtManager),
		UserService: user.NewService(database),
	}
}
