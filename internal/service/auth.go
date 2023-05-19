package service

import (
	"context"
	"github.com/wtkeqrf0/while.act/ent"
	"github.com/wtkeqrf0/while.act/internal/controller/dto"
)

type AuthPostgres interface {
	IDExist(ctx context.Context, id int) (bool, error)
	CreateUserWithPassword(ctx context.Context, auth dto.SignUp, company *ent.Company) (*ent.User, error)
	AuthUserByEmail(ctx context.Context, email string) (*ent.User, error)
}

type AuthService struct {
	postgres AuthPostgres
}

func NewAuthService(postgres AuthPostgres) *AuthService {
	return &AuthService{postgres: postgres}
}

// IDExist returns true if user Exists
func (a *AuthService) IDExist(id int) (bool, error) {
	return a.postgres.IDExist(context.Background(), id)
}

// CreateUserWithPassword without verified email and returns it (only on registration)
func (a *AuthService) CreateUserWithPassword(auth dto.SignUp, company *ent.Company) (*ent.User, error) {
	return a.postgres.CreateUserWithPassword(context.Background(), auth, company)
}

// AuthUserByEmail returns the user's password hash and username with given email (only on jwts)
func (a *AuthService) AuthUserByEmail(email string) (*ent.User, error) {
	return a.postgres.AuthUserByEmail(context.Background(), email)
}
