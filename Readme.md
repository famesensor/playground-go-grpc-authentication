# Go gRPC Authentication

This project is an example of how to implement authentication in a Go gRPC service using JWT.

## Features

* User registration and login
* JWT token generation and verification
* gRPC interceptor for authentication

## Technologies

* Go
* gRPC
* Protocol Buffers
* JWT

## Requirements

* Go 1.16+
* gRPC 1.40+
* Protocol Buffers 3.15+

## Installation

1. Clone the repository
2. Run `go build` to build the binary
3. Run `go run main.go` to start the service

## Usage

1. Register a new user using the `proto.AuthService.SignUp` method
2. Login using the `proto.AuthService.SignIn` method
3. Use the `proto.AuthService.SignIn` method to authenticate requests
