package service

import (
	"context"
	"github.com/wtkeqrf0/while.act/ent"
)

type UserPostgres interface {
	FindUserByID(ctx context.Context, id int) (*ent.User, error)
}

type UserService struct {
	postgres UserPostgres
}

func NewUserService(postgres UserPostgres) *UserService {
	return &UserService{postgres: postgres}
}

func (u UserService) FindUserByID(id int) (*ent.User, error) {
	return u.postgres.FindUserByID(context.Background(), id)
}
