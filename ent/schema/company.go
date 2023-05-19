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
		field.String("id").StorageKey("inn").Unique().Match(bind.InnRegexp).
			StructTag(`json:"inn,omitempty" example:"7707083893"`),

		field.String("name").Optional().Unique().MinLen(2).MaxLen(150).Nillable().
			StructTag(`json:"company_name,omitempty" example:"ООО \"Парк\""`),

		field.String("website").Optional().Match(bind.LinkRegexp).Nillable().
			StructTag(`json:"website,omitempty" example:"https://www.rusprofile.ru"`),

		field.String("economic_activity_branch").Optional().MinLen(5).MaxLen(50).
			StructTag(`json:"economicActivityBranch,omitempty" example:"Радиоэлектроника и приборостроение - Приборостроение"`),
	}
}

// Edges of the Company.
func (Company) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type).Unique(),
	}
}
