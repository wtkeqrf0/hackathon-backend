package postgres

import (
	"context"
	"github.com/while-act/hackathon-backend/ent"
	"github.com/while-act/hackathon-backend/ent/industry"
	"github.com/while-act/hackathon-backend/internal/controller/dao"
)

type IndustryStorage struct {
	industryClient *ent.IndustryClient
}

func NewIndustryStorage(industryClient *ent.IndustryClient) *IndustryStorage {
	return &IndustryStorage{industryClient: industryClient}
}

func (i *IndustryStorage) GetIndustry(ctx context.Context, title string) (*dao.Industry, error) {
	var ind []*dao.Industry
	err := i.industryClient.Query().Where(industry.ID(title)).Select(
		industry.FieldAvgSalary, industry.FieldAvgSalaryCad,
		industry.FieldAvgWorkersNumCad, industry.FieldAvgWorkersNum,
	).Scan(ctx, &ind)
	if ind != nil {
		return ind[0], nil
	}
	return nil, err
}
