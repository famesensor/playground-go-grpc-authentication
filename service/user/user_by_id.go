package user

import (
	"context"

	"github.com/famesensor/playground-go-grpc-authentication/model"
)

func (s *service) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	return s.database.GetByID(ctx, id)
}
