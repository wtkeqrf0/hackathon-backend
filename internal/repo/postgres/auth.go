package postgres

import (
	"context"
	"github.com/wtkeqrf0/while.act/ent"
	"github.com/wtkeqrf0/while.act/ent/user"
	"github.com/wtkeqrf0/while.act/internal/controller/dto"
)

func (r *UserStorage) IDExist(ctx context.Context, id int) (bool, error) {
	return r.userClient.Query().Where(user.ID(id)).Exist(ctx)
}

func (r *UserStorage) CreateUserWithPassword(ctx context.Context, auth dto.SignUp, company *ent.Company) (*ent.User, error) {
	return r.userClient.Create().
		SetNillableBiography(auth.Biography).
		SetNillableCountry(auth.Country).
		SetNillableCity(auth.City).
		SetNillablePosition(auth.Position).
		SetNillableFatherName(auth.FatherName).
		SetCompany(company).
		SetFirstName(auth.FirstName).
		SetLastName(auth.LastName).
		SetPasswordHash([]byte(auth.Password)).
		SetEmail(auth.Email).Save(ctx)
}

func (r *UserStorage) AuthUserByEmail(ctx context.Context, email string) (*ent.User, error) {
	return r.userClient.Query().Where(
		user.Email(email),
	).Only(ctx)
}
