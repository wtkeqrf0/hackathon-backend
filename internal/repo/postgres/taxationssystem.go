package postgres

import (
	"context"
	"github.com/while-act/hackathon-backend/ent"
	"github.com/while-act/hackathon-backend/ent/taxationsystem"
)

type TaxStorage struct {
	taxClient *ent.TaxationSystemClient
}

func NewTaxStorage(taxClient *ent.TaxationSystemClient) *TaxStorage {
	return &TaxStorage{taxClient: taxClient}
}

func (t *TaxStorage) GetTax(ctx context.Context, num *int, tax *string) (float64, error) {
	id := *num
	switch {
	case id < 10:
		id = 10
	case id < 20:
		id = 20
	case id < 40:
		id = 40
	case id < 70:
		id = 70
	case id < 100:
		id = 100
	case id < 150:
		id = 150
	case id < 200:
		id = 200
	case id < 300:
		id = 300
	}

	only, err := t.taxClient.Query().Where(taxationsystem.ID(id)).Only(ctx)
	if only == nil {
		return 0, err
	}
	switch *tax {
	case "usn6":
		return only.Usn6, nil
	case "usn15":
		return only.Usn15, nil
	case "osn":
		return only.Osn, nil
	}
	return 0, err
}
