package service

import (
	"context"
	"github.com/while-act/hackathon-backend/internal/controller/dao"
)

type IndustryPostgres interface {
	GetIndustry(ctx context.Context, title string) (*dao.Industry, error)
}

type IndustryService struct {
	postgres IndustryPostgres
}

func NewIndustryService(postgres IndustryPostgres) *IndustryService {
	return &IndustryService{postgres: postgres}
}

func (i *IndustryService) GetIndustry(title string) (*dao.Industry, error) {
	return i.postgres.GetIndustry(context.Background(), title)
}
