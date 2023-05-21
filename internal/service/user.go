package service

import (
	"context"
	"github.com/wtkeqrf0/while.act/internal/controller/dao"
	"github.com/wtkeqrf0/while.act/internal/controller/dto"
)

type UserPostgres interface {
	FindUserByID(ctx context.Context, id int) (*dao.Me, error)
	UpdateUser(ctx context.Context, updateUser dto.UpdateUser, id int) error
	UpdatePassword(ctx context.Context, updPassword dto.UpdatePassword, id int) error
}

type UserService struct {
	postgres UserPostgres
}

func NewUserService(postgres UserPostgres) *UserService {
	return &UserService{postgres: postgres}
}

func (u *UserService) FindUserByID(id int) (*dao.Me, error) {
	return u.postgres.FindUserByID(context.Background(), id)
}

func (u *UserService) UpdateUser(updateUser dto.UpdateUser, id int) error {
	return u.postgres.UpdateUser(context.Background(), updateUser, id)
}

func (u *UserService) UpdatePassword(updPassword dto.UpdatePassword, id int) error {
	return u.postgres.UpdatePassword(context.Background(), updPassword, id)
}
