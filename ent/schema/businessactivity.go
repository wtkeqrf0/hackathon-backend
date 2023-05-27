package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// BusinessActivity holds the schema definition for the BusinessActivity entity.
type BusinessActivity struct {
	ent.Schema
}

// Fields of the BusinessActivity.
func (BusinessActivity) Fields() []ent.Field {
	return []ent.Field{
		field.String("type"),
		field.String("sub_type"),
		field.Float("total").Positive(),
	}
}

// Edges of the BusinessActivity.
func (BusinessActivity) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("histories", History.Type),
	}
}
