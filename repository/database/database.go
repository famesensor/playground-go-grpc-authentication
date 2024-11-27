package database

import (
	"context"
	"errors"

	"github.com/famesensor/playground-go-grpc-authentication/model"
	"github.com/famesensor/playground-go-grpc-authentication/port"
	"github.com/patrickmn/go-cache"
)

type databasePort struct {
	db *cache.Cache
}

func NewDatabase(db *cache.Cache) port.DatabasePort {
	return &databasePort{
		db,
	}
}

func (d *databasePort) Create(ctx context.Context, req *model.CreateUserReq) error {
	user := &User{
		ID:        req.ID,
		Username:  req.Username,
		Password:  req.Password,
		CreatedAt: req.CreatedAt,
	}
	d.db.Set(req.ID, user, -1)
	d.db.Set(req.Username, user, -1)
	return nil
}

func (d *databasePort) GetByID(ctx context.Context, id string) (*model.User, error) {
	res, ok := d.db.Get(id)
	if !ok {
		return nil, errors.New("user not found")
	}
	user := res.(*User)
	return user.ToUserModel(), nil
}

func (d *databasePort) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	res, ok := d.db.Get(username)
	if !ok {
		return nil, errors.New("user not found")
	}
	user := res.(*User)
	return user.ToUserModel(), nil
}
