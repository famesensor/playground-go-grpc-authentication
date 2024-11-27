package port

import (
	"context"

	"github.com/famesensor/playground-go-grpc-authentication/model"
)

type DatabasePort interface {
	Create(context.Context, *model.CreateUserReq) error
	GetByUsername(context.Context, string) (*model.User, error)
	GetByID(context.Context, string) (*model.User, error)
}
