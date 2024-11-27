package model

import "time"

type CreateUserReq struct {
	ID        string
	Username  string
	Password  string
	CreatedAt time.Time
}
