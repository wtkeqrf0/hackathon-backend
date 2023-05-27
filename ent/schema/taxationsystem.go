package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// TaxationSystem holds the schema definition for the TaxationSystem entity.
type TaxationSystem struct {
	ent.Schema
}

// Fields of the TaxationSystem.
func (TaxationSystem) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").StorageKey("operations").Unique().Positive(),
		field.Float("usn6").Positive(),
		field.Float("usn15").Positive(),
		field.Float("osn").Positive(),
	}
}

// Edges of the TaxationSystem.
func (TaxationSystem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("histories", History.Type),
	}
}
