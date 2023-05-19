package postgres

import (
	"context"
	"github.com/wtkeqrf0/while.act/ent"
	"github.com/wtkeqrf0/while.act/ent/user"
	"github.com/wtkeqrf0/while.act/internal/controller/dao"
	"github.com/wtkeqrf0/while.act/internal/controller/dto"
	"github.com/wtkeqrf0/while.act/pkg/middleware/errs"
)

type UserStorage struct {
	userClient *ent.UserClient
}

func NewUserStorage(userClient *ent.UserClient) *UserStorage {
	return &UserStorage{userClient: userClient}
}

func (r *UserStorage) FindUserByID(ctx context.Context, id int) (*dao.Me, error) {
	var me []*dao.Me
	err := r.userClient.Query().Where(user.ID(id)).Select(user.FieldCity, user.FieldBiography,
		user.FieldCountry, user.FieldName, user.FieldLastName,
		user.FieldFirstName, user.FieldFatherName, user.FieldEmail,
		user.FieldPosition, "company_inn").Scan(ctx, &me)
	if me != nil {
		return me[0], err
	}
	return nil, errs.NoSuchUser.AddErr(err)
}

func (r *UserStorage) UpdateUser(ctx context.Context, updateUser dto.UpdateUser, id int) error {
	return r.userClient.UpdateOneID(id).
		SetNillableFatherName(updateUser.FatherName).
		SetNillableCountry(updateUser.Country).
		SetNillableCity(updateUser.City).
		SetNillablePosition(updateUser.Position).
		SetNillableBiography(updateUser.Biography).Exec(ctx)
}
