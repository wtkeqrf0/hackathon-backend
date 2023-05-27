package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// History holds the schema definition for the History entity.
type History struct {
	ent.Schema
}

// Fields of the History.
func (History) Fields() []ent.Field {
	return []ent.Field{
		field.String("company_name").MinLen(2).MaxLen(150),

		field.String("industry_branch").Immutable(),

		field.Int("full_time_employees").Positive().Immutable(),

		field.String("district_title").Immutable(),

		field.Float("land_area").Positive().Immutable(),

		field.Float("construction_facilities_area").Positive().Immutable(),

		field.String("equipment_type").Immutable(),

		field.String("organization_type").Immutable(),

		field.String("facility_type").Immutable(),

		field.Bool("accounting_services").Immutable(),

		field.Bool("patent").Immutable(),

		field.Text("other").Immutable(),

		field.Int("user_id").Immutable(),
	}
}

// Edges of the History.
func (History) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("industry", Industry.Type).
			Ref("histories").
			Required().Immutable().
			Unique().Field("industry_branch"),

		edge.From("district", District.Type).
			Ref("histories").
			Required().Immutable().
			Unique().Field("district_title"),

		edge.From("equipment", Equipment.Type).
			Ref("histories").
			Required().Immutable().
			Unique().Field("equipment_type"),

		edge.From("users", User.Type).
			Ref("histories").
			Required().Immutable().
			Unique().Field("user_id"),
	}
}
