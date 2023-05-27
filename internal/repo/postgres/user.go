package postgres

import (
	"context"
	"github.com/while-act/hackathon-backend/ent"
	"github.com/while-act/hackathon-backend/ent/history"
	"github.com/while-act/hackathon-backend/ent/user"
	"github.com/while-act/hackathon-backend/internal/controller/dao"
	"github.com/while-act/hackathon-backend/internal/controller/dto"
	"golang.org/x/crypto/bcrypt"
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
		user.FieldPosition, user.FieldRole, user.FieldCompanyID).Scan(ctx, &me)
	if me != nil {
		return me[0], nil
	}
	return nil, err
}

func (r *UserStorage) UpdateUser(ctx context.Context, updateUser *dto.UpdateUser, id int) error {
	return r.userClient.UpdateOneID(id).
		SetNillableFatherName(updateUser.FatherName).
		SetNillableCountry(updateUser.Country).
		SetNillableCity(updateUser.City).
		SetNillablePosition(updateUser.Position).
		SetNillableBiography(updateUser.Biography).Exec(ctx)
}

func (r *UserStorage) UpdatePassword(ctx context.Context, newPassword []byte, email string) error {
	return r.userClient.Update().SetPasswordHash(newPassword).Where(user.Email(email)).Exec(ctx)
}

func (r *UserStorage) UpdateEmail(ctx context.Context, password []byte, newEmail string, id int) error {
	customer, err := r.userClient.Get(ctx, id)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword(customer.PasswordHash, password)
	if err != nil {
		return err
	}

	return r.userClient.UpdateOneID(id).SetEmail(newEmail).Exec(ctx)
}

func (r *UserStorage) GetAllHistory(ctx context.Context, userId int) ([]*dao.Histories, error) {
	customer, err := r.userClient.Get(ctx, userId)
	if err != nil {
		return nil, err
	}

	var histories []*dao.Histories
	err = customer.QueryHistories().Select(history.FieldID, history.FieldCompanyName).Scan(ctx, &histories)

	if histories != nil {
		return histories, nil
	}
	return nil, err
}

func (r *UserStorage) GetOneHistory(ctx context.Context, historyId int, userId int) (*ent.History, error) {
	customer, err := r.userClient.Get(ctx, userId)
	if err != nil {
		return nil, err
	}

	return customer.QueryHistories().Where(history.ID(historyId)).Only(ctx)
}
