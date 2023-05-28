package postgres

import (
	"context"
	"github.com/while-act/hackathon-backend/ent"
)

type DistrictStorage struct {
	districtClient *ent.DistrictClient
}

func NewDistrictStorage(districtClient *ent.DistrictClient) *DistrictStorage {
	return &DistrictStorage{districtClient: districtClient}
}

func (d *DistrictStorage) GetDistrict(ctx context.Context, title string) (*ent.District, error) {
	return d.districtClient.Get(ctx, title)
}
