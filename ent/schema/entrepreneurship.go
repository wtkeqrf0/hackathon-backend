package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Entrepreneurship holds the schema definition for the Entrepreneurship entity.
type Entrepreneurship struct {
	ent.Schema
}

// Fields of the Entrepreneurship.
func (Entrepreneurship) Fields() []ent.Field {
	return []ent.Field{
		field.String("type").Unique(),
	}
}
