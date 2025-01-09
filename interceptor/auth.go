package interceptor

import (
	"context"
	"strings"

	"github.com/famesensor/ghelper"
	"github.com/famesensor/playground-go-grpc-authentication/constant"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type AuthInterceptor interface {
	Unary() grpc.UnaryServerInterceptor
}

type authInterceptor struct {
	jwtManager ghelper.JWTManager
	skipPath   map[string]struct{}
}

func NewAuthInterceptor(jwtManager ghelper.JWTManager, skipPaths map[string]struct{}) AuthInterceptor {
	return &authInterceptor{jwtManager, skipPaths}
}

func (interceptor *authInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		userClaim, err := interceptor.authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}

		if userClaim == nil {
			return handler(ctx, req)
		}

		c := context.WithValue(ctx, constant.UserCtxKey, userClaim)
		return handler(c, req)
	}
}

func (interceptor *authInterceptor) authorize(ctx context.Context, method string) (*ghelper.UserClaims, error) {
	if _, ok := interceptor.skipPath[method]; ok {
		return nil, nil
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) <= 0 {
		return nil, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	token := strings.TrimPrefix(values[0], "Bearer ")
	user, err := interceptor.jwtManager.Verification(token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	return user, nil
}
