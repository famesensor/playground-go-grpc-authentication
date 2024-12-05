package model

type SignInReq struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
}

type SignInRes struct {
	AccessToken string
}
