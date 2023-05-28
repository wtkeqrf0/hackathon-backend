package service

import (
	"context"
	"github.com/while-act/hackathon-backend/internal/controller/dto"
)

type BusinessPostgres interface {
	GetBusiness(ctx context.Context, bus *dto.BusinessActivity) (*int, error)
}

type BusinessService struct {
	postgres BusinessPostgres
}

func NewBusinessService(postgres BusinessPostgres) *BusinessService {
	return &BusinessService{postgres: postgres}
}

func (b *BusinessService) GetBusiness(bus *dto.BusinessActivity) (*int, error) {
	return b.postgres.GetBusiness(context.Background(), bus)
}
