package service

import (
	"context"
	"github.com/while-act/hackathon-backend/ent"
	"github.com/while-act/hackathon-backend/internal/controller/dto"
)

type HistoryPostgres interface {
	GetHistory(ctx context.Context, historyId int) (*ent.History, error)
	CreateHistory(ctx context.Context, h *dto.History, busactId *int, userId int) error
}

type HistoryService struct {
	postgres HistoryPostgres
}

func NewHistoryService(postgres HistoryPostgres) *HistoryService {
	return &HistoryService{postgres: postgres}
}

func (i *HistoryService) GetHistory(historyId int) (*ent.History, error) {
	return i.postgres.GetHistory(context.Background(), historyId)
}

func (i *HistoryService) CreateHistory(h *dto.History, busactId *int, id int) error {
	return i.postgres.CreateHistory(context.Background(), h, busactId, id)
}
