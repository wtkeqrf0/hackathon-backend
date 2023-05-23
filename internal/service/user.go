package service

import (
	"context"
	"github.com/wtkeqrf0/while.act/internal/controller/dao"
	"github.com/wtkeqrf0/while.act/internal/controller/dto"
	"time"
)

type UserPostgres interface {
	FindUserByID(ctx context.Context, id int) (*dao.Me, error)
	UpdateUser(ctx context.Context, updateUser dto.UpdateUser, id int) error
	UpdatePassword(ctx context.Context, newPassword []byte, email string) error
	UpdateEmail(ctx context.Context, password []byte, newEmail string, id int) error
}

type UserService struct {
	postgres UserPostgres
	redis    UserRedis
}

func NewUserService(postgres UserPostgres, redis UserRedis) *UserService {
	return &UserService{postgres: postgres, redis: redis}
}

func (u *UserService) FindUserByID(id int) (*dao.Me, error) {
	return u.postgres.FindUserByID(context.Background(), id)
}

func (u *UserService) UpdateUser(updateUser dto.UpdateUser, id int) error {
	return u.postgres.UpdateUser(context.Background(), updateUser, id)
}

func (u *UserService) UpdatePassword(newPassword []byte, email string) error {
	return u.postgres.UpdatePassword(context.Background(), newPassword, email)
}

func (u *UserService) UpdateEmail(password []byte, newEmail string, id int) error {
	return u.postgres.UpdateEmail(context.Background(), password, newEmail, id)
}

type UserRedis interface {
	ContainsKeys(ctx context.Context, keys ...string) (int64, error)
	SetVariable(ctx context.Context, key string, value any, exp time.Duration) error
}

// ContainsKeys of redis by key
func (u *UserService) ContainsKeys(keys ...string) (int64, error) {
	return u.redis.ContainsKeys(context.Background(), keys...)
}

// SetVariable of redis by key, his value and exploration time
func (u *UserService) SetVariable(key string, value any, exp time.Duration) error {
	return u.redis.SetVariable(context.Background(), key, value, exp)
}
