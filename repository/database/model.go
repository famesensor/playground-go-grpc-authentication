package database

import (
	"time"

	"github.com/famesensor/playground-go-grpc-authentication/model"
)

type User struct {
	ID        string
	Username  string
	Password  string
	CreatedAt time.Time
}

func (u *User) ToUserModel() *model.User {
	return &model.User{
		ID:       u.ID,
		Username: u.Username,
		Password: u.Password,
	}
}
