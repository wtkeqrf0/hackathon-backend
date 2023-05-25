package postgres

import (
	"context"
	"github.com/while-act/hackathon-backend/ent"
	"github.com/while-act/hackathon-backend/ent/company"
	"github.com/while-act/hackathon-backend/internal/controller/dao"
	"github.com/while-act/hackathon-backend/internal/controller/dto"
)

type CompanyStorage struct {
	companyClient *ent.CompanyClient
}

func NewCompanyStorage(companyClient *ent.CompanyClient) *CompanyStorage {
	return &CompanyStorage{companyClient: companyClient}
}

func (r *CompanyStorage) CreateCompany(ctx context.Context, inn string, name, website *string) (*ent.Company, error) {
	return r.companyClient.Create().
		SetInn(inn).SetNillableName(name).
		SetNillableWebsite(website).Save(ctx)
}

func (r *CompanyStorage) GetCompanyDTO(ctx context.Context, id int) (*dao.Company, error) {
	var comp []*dao.Company
	err := r.companyClient.Query().Where(company.ID(id)).Select(
		company.FieldID, company.FieldWebsite, company.FieldName).Scan(ctx, &comp)

	if comp != nil {
		return comp[0], nil
	}
	return nil, err
}

func (r *CompanyStorage) GetCompany(ctx context.Context, id int) (*ent.Company, error) {
	return r.companyClient.Get(ctx, id)
}

func (r *CompanyStorage) UpdateCompany(ctx context.Context, updateCompany *dto.UpdateCompany, id int) error {
	return r.companyClient.UpdateOneID(id).
		SetNillableName(updateCompany.Name).
		SetNillableWebsite(updateCompany.Website).Exec(ctx)
}
