package service

import (
	"context"
	"github.com/wtkeqrf0/while.act/ent"
	"github.com/wtkeqrf0/while.act/internal/controller/dao"
	"github.com/wtkeqrf0/while.act/internal/controller/dto"
)

type CompanyPostgres interface {
	CreateCompany(ctx context.Context, inn string, name, website *string) (*ent.Company, error)
	GetCompany(ctx context.Context, inn string) (*ent.Company, error)
	GetCompanyDTO(ctx context.Context, inn string) (*dao.Company, error)
	UpdateCompany(ctx context.Context, updateCompany dto.UpdateCompany, inn string) error
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

func (c *CompanyService) GetCompany(inn string) (*ent.Company, error) {
	return c.postgres.GetCompany(context.Background(), inn)
}

func (c *CompanyService) GetCompanyDTO(inn string) (*dao.Company, error) {
	return c.postgres.GetCompanyDTO(context.Background(), inn)
}

func (c *CompanyService) UpdateCompany(updateCompany dto.UpdateCompany, inn string) error {
	return c.postgres.UpdateCompany(context.Background(), updateCompany, inn)
}
