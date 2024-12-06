package handler

import (
	"context"

	"github.com/famesensor/playground-go-grpc-authentication/constant"
	"github.com/famesensor/playground-go-grpc-authentication/helper"
	proto "github.com/famesensor/playground-go-grpc-authentication/proto/user"
	"github.com/famesensor/playground-go-grpc-authentication/service/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type userHandler struct {
	proto.UnimplementedUserServiceServer
	userService user.Port
}

func NewUserHandler(userService user.Port) proto.UserServiceServer {
	return &userHandler{
		userService: userService,
	}
}

func (u *userHandler) GetUserMe(ctx context.Context, req *proto.GetUserMeReq) (*proto.User, error) {
	user, ok := ctx.Value(constant.UserCtxKey).(*helper.UserClaims)
	if !ok {
		return nil, status.Errorf(codes.DataLoss, "user not found")
	}

	res, err := u.userService.GetUserByID(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return &proto.User{
		Id:       res.ID,
		Username: res.Username,
	}, nil
}
