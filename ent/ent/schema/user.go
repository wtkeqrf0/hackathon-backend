package schema

import (
	"context"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"fmt"
	gen "github.com/wtkeqrf0/while.act/ent"
	"github.com/wtkeqrf0/while.act/ent/hook"
	"github.com/wtkeqrf0/while.act/pkg/bind"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

const emailRegexp string = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").StructTag(`json:"-"`),

		field.String("name").Unique().Match(bind.NameRegexp).Annotations(
			entsql.DefaultExpr("'user' || setval(pg_get_serial_sequence('users','id')," +
				"nextval(pg_get_serial_sequence('users','id'))-1)")).
			StructTag(`json:"name,omitempty" example:"user94"`),

		field.String("email").Unique().Match(regexp.MustCompile(emailRegexp)).
			StructTag(`json:"email,omitempty" example:"myemail@gmail.com"`),

		field.Bytes("password_hash").Optional().Sensitive().Nillable(),

		field.Text("biography").Optional().MaxLen(512).Nillable().
			StructTag(`json:"biography,omitempty" example:"I'd like to relax"`),

		field.String("role").Default("USER").StructTag(`json:"role,omitempty" example:"USER"`),

		field.String("first_name").Optional().MinLen(2).MaxLen(30).Nillable().
			StructTag(`json:"firstName,omitempty" example:"Tele"`),

		field.String("last_name").Optional().MinLen(2).MaxLen(30).Nillable().
			StructTag(`json:"lastName,omitempty" example:"phone"`),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
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
	return hook.UserFunc(func(ctx context.Context, m *gen.UserMutation) (ent.Value, error) {
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
