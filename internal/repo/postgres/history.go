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

func (h *HistoryStorage) GetHistory(ctx context.Context, historyId int) (*ent.History, error) {
	return h.historyClient.Get(ctx, historyId)
}

func (h *HistoryStorage) CreateHistory(ctx context.Context, data *dto.History, busactId *int, userId int) error {
	return h.historyClient.Create().
		SetName(data.Name).
		SetOrganizationalLegal(data.OrganizationLegal).
		SetIndustryBranch(data.IndustryBranch).
		SetFullTimeEmployees(data.FullTimeEmployees).
		SetAvgSalary(data.AvgSalary).
		SetDistrictTitle(data.DistrictTitle).
		SetLandArea(data.LandArea).
		SetIsBuy(data.IsBuy).
		SetConstructionFacilitiesArea(data.ConstructionFacilitiesArea).
		SetBuildingType(data.BuildingType).
		SetEquipment(data.Equipment).
		SetAccountingSupport(data.AccountingSupport).
		SetNillableTaxationSystemOperations(data.TaxationSystemOperations).
		SetNillableOperationType(data.OperationsType).
		SetPatentCalc(data.PatentCalc).
		SetNillableBusinessActivityID(busactId).
		SetNillableOther(data.Other).
		SetUserID(userId).Exec(ctx)

}
