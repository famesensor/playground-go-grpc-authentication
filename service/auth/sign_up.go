package auth

import (
	"context"
	"time"

	"github.com/famesensor/playground-go-grpc-authentication/model"
)

func (s *service) SignUp(ctx context.Context, req *model.SignUpReq) error {
	if err := s.database.Create(ctx, &model.CreateUserReq{
		Username:  req.Username,
		Password:  req.Password,
		ID:        s.uuid.New(),
		CreatedAt: time.Now(),
	}); err != nil {
		return err
	}
	return nil
}
