package model

type SignUpReq struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
}
