package postgres

import (
	"context"
	"github.com/while-act/hackathon-backend/ent"
	"github.com/while-act/hackathon-backend/internal/controller/dto"
)

type HistoryStorage struct {
	historyClient *ent.HistoryClient
}

func NewHistoryStorage(historyClient *ent.HistoryClient) *HistoryStorage {
	return &HistoryStorage{historyClient: historyClient}
}

func (h *HistoryStorage) GetHistory(ctx context.Context, companyName string) (*ent.History, error) {
	return h.historyClient.Get(ctx, companyName)
}

func (h *HistoryStorage) CreateHistory(ctx context.Context, data *dto.History, userId int) error {
	return h.historyClient.Create().
		SetID(data.CompanyName).
		SetDistrictTitle(data.DistrictTitle).
		SetAccountingServices(data.AccountingServices).
		SetEquipmentType(data.EquipmentType).
		SetFacilityType(data.FacilityType).
		SetConstructionFacilitiesArea(data.ConstructionFacilitiesArea).
		SetLandArea(data.LandArea).
		SetFullTimeEmployees(data.FullTimeEmployees).
		SetIndustryBranch(data.IndustryBranch).
		SetOther(data.Other).
		SetPatent(data.Patent).
		SetUserID(userId).Exec(ctx)
}
