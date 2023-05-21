package postgres

import (
	"context"
	"github.com/wtkeqrf0/while.act/ent"
	"github.com/wtkeqrf0/while.act/ent/company"
	"github.com/wtkeqrf0/while.act/internal/controller/dao"
	"github.com/wtkeqrf0/while.act/internal/controller/dto"
	"github.com/wtkeqrf0/while.act/pkg/middleware/errs"
)

type CompanyStorage struct {
	companyClient *ent.CompanyClient
}

func NewCompanyStorage(companyClient *ent.CompanyClient) *CompanyStorage {
	return &CompanyStorage{companyClient: companyClient}
}

func (r *CompanyStorage) CreateCompany(ctx context.Context, inn string, name, website *string) (*ent.Company, error) {
	return r.companyClient.Create().SetID(inn).
		SetNillableName(name).
		SetNillableWebsite(website).Save(ctx)
}

func (r *CompanyStorage) GetCompanyDTO(ctx context.Context, inn string) (*dao.Company, error) {
	var comp []*dao.Company
	err := r.companyClient.Query().Where(company.ID(inn)).Select(
		company.FieldID, company.FieldWebsite, company.FieldName).Scan(ctx, &comp)

	if comp != nil {
		return comp[0], err
	}
	return nil, errs.NoSuchCompany.AddErr(err)
}

func (r *CompanyStorage) GetCompany(ctx context.Context, inn string) (*ent.Company, error) {
	return r.companyClient.Get(ctx, inn)
}

func (r *CompanyStorage) UpdateCompany(ctx context.Context, updateCompany dto.UpdateCompany, inn string) error {
	return r.companyClient.UpdateOneID(inn).
		SetNillableName(updateCompany.Name).
		SetNillableWebsite(updateCompany.Website).Exec(ctx)
}
