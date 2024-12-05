package auth

import (
	"context"
	"errors"

	"github.com/famesensor/playground-go-grpc-authentication/model"
)

func (s *service) SignIn(ctx context.Context, req *model.SignInReq) (*model.SignInRes, error) {
	user, err := s.database.GetByUsername(ctx, req.Username)
	if err != nil {
		return nil, errors.New("username already exists")
	}

	if user.Password != req.Password {
		return nil, errors.New("invalid password")
	}

	tk, err := s.jwtManager.Generate(user.ID, user.Username)
	if err != nil {
		return nil, err
	}

	return &model.SignInRes{
		AccessToken: tk,
	}, nil
}
