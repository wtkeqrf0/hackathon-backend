package service

import (
	"context"
)

type TaxPostgres interface {
	GetTax(ctx context.Context, num *int, tax *string) (float64, error)
}

type TaxService struct {
	postgres TaxPostgres
}

func NewTaxService(postgres TaxPostgres) *TaxService {
	return &TaxService{postgres: postgres}
}

func (t *TaxService) GetTax(num *int, tax *string) (float64, error) {
	return t.postgres.GetTax(context.Background(), num, tax)
}
