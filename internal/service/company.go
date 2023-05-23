package service

import (
	"context"
	"github.com/wtkeqrf0/while.act/ent"
	"github.com/wtkeqrf0/while.act/internal/controller/dao"
	"github.com/wtkeqrf0/while.act/internal/controller/dto"
)

type CompanyPostgres interface {
	CreateCompany(ctx context.Context, inn string, name, website *string) (*ent.Company, error)
	GetCompany(ctx context.Context, id int) (*ent.Company, error)
	GetCompanyDTO(ctx context.Context, id int) (*dao.Company, error)
	UpdateCompany(ctx context.Context, updateCompany dto.UpdateCompany, id int) error
}

type CompanyService struct {
	postgres CompanyPostgres
}

func NewCompanyService(postgres CompanyPostgres) *CompanyService {
	return &CompanyService{postgres: postgres}
}
func (c *CompanyService) CreateCompany(inn string, name, website *string) (*ent.Company, error) {
	return c.postgres.CreateCompany(context.Background(), inn, name, website)
}

func (c *CompanyService) GetCompany(id int) (*ent.Company, error) {
	return c.postgres.GetCompany(context.Background(), id)
}

func (c *CompanyService) GetCompanyDTO(id int) (*dao.Company, error) {
	return c.postgres.GetCompanyDTO(context.Background(), id)
}

func (c *CompanyService) UpdateCompany(updateCompany dto.UpdateCompany, id int) error {
	return c.postgres.UpdateCompany(context.Background(), updateCompany, id)
}
