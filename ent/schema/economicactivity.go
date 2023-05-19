package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// EconomicActivity holds the schema definition for the EconomicActivity entity.
type EconomicActivity struct {
	ent.Schema
}

// Fields of the EconomicActivity.
func (EconomicActivity) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").StorageKey("main").Unique().MinLen(5).MaxLen(100).
			StructTag(`json:"main,omitempty" example:"Автомобильная промышленность"`),

		field.String("subs").Optional().Unique().MinLen(5).MaxLen(100).
			StructTag(`json:"subs,omitempty" example:"Автомобильная промышленность"`),
	}
}

// Edges of the EconomicActivity.
func (EconomicActivity) Edges() []ent.Edge {
	return nil
}
