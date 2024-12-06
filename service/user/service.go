package user

import (
	"context"

	"github.com/famesensor/playground-go-grpc-authentication/model"
	"github.com/famesensor/playground-go-grpc-authentication/port"
)

type Port interface {
	GetUserByID(context.Context, string) (*model.User, error)
}

type service struct {
	database port.DatabasePort
}

func NewService(database port.DatabasePort) Port {
	return &service{database}
}
