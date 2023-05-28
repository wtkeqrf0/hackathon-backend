package postgres

import (
	"context"
	"github.com/while-act/hackathon-backend/ent"
	"github.com/while-act/hackathon-backend/ent/businessactivity"
	"github.com/while-act/hackathon-backend/internal/controller/dto"
)

type BusinessStorage struct {
	businessClient *ent.BusinessActivityClient
}

func NewBusinessStorage(businessClient *ent.BusinessActivityClient) *BusinessStorage {
	return &BusinessStorage{businessClient: businessClient}
}

func (b *BusinessStorage) GetBusiness(ctx context.Context, bus *dto.BusinessActivity) (*int, error) {
	ids, err := b.businessClient.Query().Where(
		businessactivity.Type(bus.Type),
		businessactivity.SubType(bus.SubType),
	).Unique(true).IDs(ctx)
	if ids != nil {
		return &ids[0], err
	}
	return nil, err
}
