package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Industry holds the schema definition for the Industry entity.
type Industry struct {
	ent.Schema
}

// Fields of the Industry.
func (Industry) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").StorageKey("branch").Unique(),

		field.Float("avg_workers_num").Positive(),

		field.Float("avg_workers_num_cad").Positive(),

		field.Float("avg_salary").Positive(),

		field.Float("avg_salary_cad").Positive(),
	}
}

// Edges of the Industry.
func (Industry) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("histories", History.Type),
	}
}
