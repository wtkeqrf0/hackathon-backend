package service

import (
	"context"
	"github.com/while-act/hackathon-backend/ent"
)

type DistrictPostgres interface {
	GetDistrict(ctx context.Context, title string) (*ent.District, error)
}

type DistrictService struct {
	postgres DistrictPostgres
}

func NewDistrictService(postgres DistrictPostgres) *DistrictService {
	return &DistrictService{postgres: postgres}
}

func (d *DistrictService) GetDistrict(title string) (*ent.District, error) {
	return d.postgres.GetDistrict(context.Background(), title)
}
