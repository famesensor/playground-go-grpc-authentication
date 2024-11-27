package handler

import (
	"context"

	"github.com/famesensor/playground-go-grpc-authentication/model"
	proto "github.com/famesensor/playground-go-grpc-authentication/proto/auth"
	"github.com/famesensor/playground-go-grpc-authentication/service/auth"
	"github.com/go-playground/validator/v10"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type authHandler struct {
	proto.UnimplementedAuthServiceServer
	authService auth.Port
	validate    *validator.Validate
}

func NewAuthHandler(authService auth.Port, validate *validator.Validate) proto.AuthServiceServer {
	return &authHandler{
		authService: authService,
		validate:    validate,
	}
}

func (h *authHandler) SignIn(ctx context.Context, req *proto.SignInReq) (*proto.SignInRes, error) {
	body := model.SignUpReq{
		Username: req.Username,
		Password: req.Password,
	}

	if err := h.validate.Struct(&body); err != nil {
		return nil, err
	}

	if err := h.authService.SignUp(ctx, &body); err != nil {
		return nil, err
	}

	return &proto.SignInRes{
		Timestamp: timestamppb.Now(),
	}, nil
}
