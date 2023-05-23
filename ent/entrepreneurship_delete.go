// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wtkeqrf0/while.act/ent/entrepreneurship"
	"github.com/wtkeqrf0/while.act/ent/predicate"
)

// EntrepreneurshipDelete is the builder for deleting a Entrepreneurship entity.
type EntrepreneurshipDelete struct {
	config
	hooks    []Hook
	mutation *EntrepreneurshipMutation
}

// Where appends a list predicates to the EntrepreneurshipDelete builder.
func (ed *EntrepreneurshipDelete) Where(ps ...predicate.Entrepreneurship) *EntrepreneurshipDelete {
	ed.mutation.Where(ps...)
	return ed
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ed *EntrepreneurshipDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ed.sqlExec, ed.mutation, ed.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ed *EntrepreneurshipDelete) ExecX(ctx context.Context) int {
	n, err := ed.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ed *EntrepreneurshipDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(entrepreneurship.Table, sqlgraph.NewFieldSpec(entrepreneurship.FieldID, field.TypeInt))
	if ps := ed.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ed.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ed.mutation.done = true
	return affected, err
}

// EntrepreneurshipDeleteOne is the builder for deleting a single Entrepreneurship entity.
type EntrepreneurshipDeleteOne struct {
	ed *EntrepreneurshipDelete
}

// Where appends a list predicates to the EntrepreneurshipDelete builder.
func (edo *EntrepreneurshipDeleteOne) Where(ps ...predicate.Entrepreneurship) *EntrepreneurshipDeleteOne {
	edo.ed.mutation.Where(ps...)
	return edo
}

// Exec executes the deletion query.
func (edo *EntrepreneurshipDeleteOne) Exec(ctx context.Context) error {
	n, err := edo.ed.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{entrepreneurship.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (edo *EntrepreneurshipDeleteOne) ExecX(ctx context.Context) {
	if err := edo.Exec(ctx); err != nil {
		panic(err)
	}
}
