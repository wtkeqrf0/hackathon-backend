package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/while-act/hackathon-backend/internal/controller/dto"
)

// History holds the schema definition for the History entity.
type History struct {
	ent.Schema
}

// Fields of the History.
func (History) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MinLen(2).MaxLen(200),

		field.String("organizational_legal").Immutable(),

		field.String("industry_branch").Immutable(),

		field.Int("full_time_employees").Positive().Immutable(),

		field.Float("avg_salary").Immutable(),

		field.String("district_title").Immutable(),

		field.Float("land_area").Positive().Immutable(),

		field.Bool("is_buy").Immutable(),

		field.Float("construction_facilities_area").Positive().Immutable(),

		field.String("building_type").Immutable(),

		field.JSON("equipment", []dto.Equipment{}).Immutable(),

		field.Bool("accounting_support").Immutable(), //true == next two

		field.Int("taxation_system_operations").Optional().Immutable(),

		field.String("operation_type").Optional().Immutable(),

		field.Bool("patent_calc").Immutable(), //true == next one

		field.Int("business_activity_id").Optional().Immutable(),

		field.Text("other").Optional().Immutable(),

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

		edge.From("taxation_systems", TaxationSystem.Type).
			Ref("histories").Immutable().
			Unique().Field("taxation_system_operations"),

		edge.From("business_activity", BusinessActivity.Type).
			Ref("histories").Immutable().
			Unique().Field("business_activity_id"),

		edge.From("district", District.Type).
			Ref("histories").
			Required().Immutable().
			Unique().Field("district_title"),

		edge.From("users", User.Type).
			Ref("histories").
			Required().Immutable().
			Unique().Field("user_id"),
	}
}
