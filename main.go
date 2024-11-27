package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/famesensor/playground-go-grpc-authentication/handler"
	"github.com/famesensor/playground-go-grpc-authentication/helper"
	"github.com/famesensor/playground-go-grpc-authentication/proto/auth"
	"github.com/famesensor/playground-go-grpc-authentication/proto/user"
	"github.com/famesensor/playground-go-grpc-authentication/repository/database"
	"github.com/famesensor/playground-go-grpc-authentication/service"
	"github.com/go-playground/validator/v10"
	"github.com/patrickmn/go-cache"
	"google.golang.org/grpc"
)

func main() {

	db := cache.New(5*time.Minute, 10*time.Minute)
	validate := validator.New()

	databaseAdapter := database.NewDatabase(db)

	services := service.NewService(databaseAdapter, helper.NewUUID())

	s := grpc.NewServer()
	auth.RegisterAuthServiceServer(s, handler.NewAuthHandler(services.AuthService, validate))
	user.RegisterUserServiceServer(s, handler.NewUserHandler())

	port := 9000
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
