package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Equipment holds the schema definition for the Equipment entity.
type Equipment struct {
	ent.Schema
}

// Fields of the Equipment.
func (Equipment) Fields() []ent.Field {
	return []ent.Field{
		field.String("type").Unique(),
		field.Int("avg_price_dol"),
		field.Int("avg_price_rub"),
	}
}
