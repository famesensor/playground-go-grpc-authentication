package auth

import (
	"context"

	"github.com/famesensor/playground-go-grpc-authentication/helper"
	"github.com/famesensor/playground-go-grpc-authentication/model"
	"github.com/famesensor/playground-go-grpc-authentication/port"
)

type Port interface {
	SignUp(context.Context, *model.SignUpReq) error
}

type service struct {
	database port.DatabasePort
	uuid     helper.UUID
}

func NewService(database port.DatabasePort, uuid helper.UUID) Port {
	return &service{
		database,
		uuid,
	}
}
