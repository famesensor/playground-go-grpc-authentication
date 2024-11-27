package helper

import "github.com/google/uuid"

type UUID interface {
	New() string
}

type uuids struct{}

func NewUUID() UUID {
	return &uuids{}
}

func (u *uuids) New() string {
	return uuid.NewString()
}
