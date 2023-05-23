package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/wtkeqrf0/while.act/pkg/bind"
)

// Company holds the schema definition for the Company entity.
type Company struct {
	ent.Schema
}

// Fields of the Company.
func (Company) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").StructTag(`json:"-"`),

		field.String("inn").Match(bind.InnRegexp).
			StructTag(`json:"inn,omitempty" example:"7707083893"`),

		field.String("name").Optional().Unique().MinLen(2).MaxLen(150).Nillable().
			StructTag(`json:"company_name,omitempty" example:"ООО 'Парк'"`),

		field.String("website").Optional().Match(bind.LinkRegexp).Nillable().
			StructTag(`json:"website,omitempty" example:"https://www.rusprofile.ru"`),
	}
}

// Edges of the Company.
func (Company) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type).Unique(),
	}
}
