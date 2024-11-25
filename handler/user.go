package handler

import (
	"context"

	"github.com/famesensor/playground-go-grpc-authentication/proto/user"
)

type userHandler struct {
	user.UnimplementedUserServiceServer
}

func NewUserHandler() user.UserServiceServer {
	return &userHandler{}
}

func (h *userHandler) CreateUser(ctx context.Context, req *user.CreateUserReq) (*user.CreateUserResp, error) {
	return nil, nil
}
