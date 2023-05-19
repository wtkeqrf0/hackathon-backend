package schema

import (
	"context"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"fmt"
	entc "github.com/wtkeqrf0/while.act/ent"
	"github.com/wtkeqrf0/while.act/ent/hook"
	"github.com/wtkeqrf0/while.act/pkg/bind"
	"golang.org/x/crypto/bcrypt"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").StructTag(`json:"-"`),

		field.String("role").Default("USER").StructTag(`json:"role,omitempty" example:"USER"`),

		field.String("name").Unique().Match(bind.NameRegexp).Annotations(
			entsql.DefaultExpr("'user' || setval(pg_get_serial_sequence('users','id')," +
				"nextval(pg_get_serial_sequence('users','id'))-1)")).
			StructTag(`json:"username,omitempty" example:"user94"`),

		field.Bytes("password_hash").Sensitive(),

		field.String("email").Unique().Match(bind.EmailRegexp).
			StructTag(`json:"email,omitempty" example:"myemail@gmail.com"`),

		field.String("first_name").MinLen(2).MaxLen(30).
			StructTag(`json:"firstName,omitempty" example:"Ivan"`),

		field.String("last_name").MinLen(2).MaxLen(30).
			StructTag(`json:"lastName,omitempty" example:"Ivanov"`),

		field.String("company_inn").Unique().Match(bind.InnRegexp).
			StructTag(`json:"inn,omitempty" example:"7707083893"`),

		field.String("father_name").Optional().MinLen(2).MaxLen(30).
			StructTag(`json:"fatherName,omitempty" example:"Ivanovich"`),

		field.String("position").Optional().MinLen(2).MaxLen(50).
			StructTag(`json:"positionName,omitempty" example:"Director"`),

		field.String("country").Optional().Match(bind.TitleRegexp).
			StructTag(`json:"country,omitempty" example:"Россия"`),

		field.String("city").Optional().Match(bind.TitleRegexp).
			StructTag(`json:"city,omitempty" example:"Москва"`),

		field.Text("biography").Optional().MaxLen(1024).Nillable().
			StructTag(`json:"biography,omitempty" example:"I'd like to relax"`),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", Company.Type).
			Ref("users").
			Unique().Required().
			Field("company_inn"),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (User) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.If(bcryptUserPassword,
			hook.And(
				hook.HasFields("password_hash"),
				hook.HasOp(ent.OpUpdate|ent.OpUpdateOne|ent.OpCreate),
			),
		),
	}
}

func bcryptUserPassword(next ent.Mutator) ent.Mutator {
	return hook.UserFunc(func(ctx context.Context, m *entc.UserMutation) (ent.Value, error) {
		password, ok := m.PasswordHash()
		if !ok {
			return nil, fmt.Errorf("password_hash is not set")
		}

		hashedPassword, err := bcrypt.GenerateFromPassword(password, 12)
		if err != nil {
			return nil, err
		}

		m.SetPasswordHash(hashedPassword)

		return next.Mutate(ctx, m)
	})
}
