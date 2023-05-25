package service

import (
	"context"
	"github.com/while-act/hackathon-backend/ent"
	"github.com/while-act/hackathon-backend/internal/controller/dto"
)

type HistoryPostgres interface {
	GetHistory(ctx context.Context, companyName string) (*ent.History, error)
	CreateHistory(ctx context.Context, h *dto.History, userId int) error
}

type HistoryService struct {
	postgres HistoryPostgres
}

func NewHistoryService(postgres HistoryPostgres) *HistoryService {
	return &HistoryService{postgres: postgres}
}

func (i *HistoryService) GetHistory(companyName string) (*ent.History, error) {
	return i.postgres.GetHistory(context.Background(), companyName)
}

func (i *HistoryService) CreateHistory(h *dto.History, id int) error {
	return i.postgres.CreateHistory(context.Background(), h, id)
}
