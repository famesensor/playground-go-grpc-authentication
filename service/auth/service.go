package auth

import (
	"context"

	"github.com/famesensor/playground-go-grpc-authentication/helper"
	"github.com/famesensor/playground-go-grpc-authentication/model"
	"github.com/famesensor/playground-go-grpc-authentication/port"
)

type Port interface {
	SignUp(context.Context, *model.SignUpReq) error
	SignIn(context.Context, *model.SignInReq) (*model.SignInRes, error)
}

type service struct {
	database   port.DatabasePort
	uuid       helper.UUID
	jwtManager helper.JWTManager
}

func NewService(database port.DatabasePort, uuid helper.UUID, jwtManager helper.JWTManager) Port {
	return &service{
		database,
		uuid,
		jwtManager,
	}
}
