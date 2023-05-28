// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/while-act/hackathon-backend/ent/businessactivity"
	"github.com/while-act/hackathon-backend/ent/history"
	"github.com/while-act/hackathon-backend/ent/predicate"
)

// BusinessActivityUpdate is the builder for updating BusinessActivity entities.
type BusinessActivityUpdate struct {
	config
	hooks    []Hook
	mutation *BusinessActivityMutation
}

// Where appends a list predicates to the BusinessActivityUpdate builder.
func (bau *BusinessActivityUpdate) Where(ps ...predicate.BusinessActivity) *BusinessActivityUpdate {
	bau.mutation.Where(ps...)
	return bau
}

// SetType sets the "type" field.
func (bau *BusinessActivityUpdate) SetType(s string) *BusinessActivityUpdate {
	bau.mutation.SetType(s)
	return bau
}

// SetSubType sets the "sub_type" field.
func (bau *BusinessActivityUpdate) SetSubType(s string) *BusinessActivityUpdate {
	bau.mutation.SetSubType(s)
	return bau
}

// SetTotal sets the "total" field.
func (bau *BusinessActivityUpdate) SetTotal(f float64) *BusinessActivityUpdate {
	bau.mutation.ResetTotal()
	bau.mutation.SetTotal(f)
	return bau
}

// AddTotal adds f to the "total" field.
func (bau *BusinessActivityUpdate) AddTotal(f float64) *BusinessActivityUpdate {
	bau.mutation.AddTotal(f)
	return bau
}

// AddHistoryIDs adds the "histories" edge to the History entity by IDs.
func (bau *BusinessActivityUpdate) AddHistoryIDs(ids ...int) *BusinessActivityUpdate {
	bau.mutation.AddHistoryIDs(ids...)
	return bau
}

// AddHistories adds the "histories" edges to the History entity.
func (bau *BusinessActivityUpdate) AddHistories(h ...*History) *BusinessActivityUpdate {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return bau.AddHistoryIDs(ids...)
}

// Mutation returns the BusinessActivityMutation object of the builder.
func (bau *BusinessActivityUpdate) Mutation() *BusinessActivityMutation {
	return bau.mutation
}

// ClearHistories clears all "histories" edges to the History entity.
func (bau *BusinessActivityUpdate) ClearHistories() *BusinessActivityUpdate {
	bau.mutation.ClearHistories()
	return bau
}

// RemoveHistoryIDs removes the "histories" edge to History entities by IDs.
func (bau *BusinessActivityUpdate) RemoveHistoryIDs(ids ...int) *BusinessActivityUpdate {
	bau.mutation.RemoveHistoryIDs(ids...)
	return bau
}

// RemoveHistories removes "histories" edges to History entities.
func (bau *BusinessActivityUpdate) RemoveHistories(h ...*History) *BusinessActivityUpdate {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return bau.RemoveHistoryIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bau *BusinessActivityUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, bau.sqlSave, bau.mutation, bau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bau *BusinessActivityUpdate) SaveX(ctx context.Context) int {
	affected, err := bau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bau *BusinessActivityUpdate) Exec(ctx context.Context) error {
	_, err := bau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bau *BusinessActivityUpdate) ExecX(ctx context.Context) {
	if err := bau.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bau *BusinessActivityUpdate) check() error {
	if v, ok := bau.mutation.Total(); ok {
		if err := businessactivity.TotalValidator(v); err != nil {
			return &ValidationError{Name: "total", err: fmt.Errorf(`ent: validator failed for field "BusinessActivityId.total": %w`, err)}
		}
	}
	return nil
}

func (bau *BusinessActivityUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := bau.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(businessactivity.Table, businessactivity.Columns, sqlgraph.NewFieldSpec(businessactivity.FieldID, field.TypeInt))
	if ps := bau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bau.mutation.GetType(); ok {
		_spec.SetField(businessactivity.FieldType, field.TypeString, value)
	}
	if value, ok := bau.mutation.SubType(); ok {
		_spec.SetField(businessactivity.FieldSubType, field.TypeString, value)
	}
	if value, ok := bau.mutation.Total(); ok {
		_spec.SetField(businessactivity.FieldTotal, field.TypeFloat64, value)
	}
	if value, ok := bau.mutation.AddedTotal(); ok {
		_spec.AddField(businessactivity.FieldTotal, field.TypeFloat64, value)
	}
	if bau.mutation.HistoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   businessactivity.HistoriesTable,
			Columns: []string{businessactivity.HistoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(history.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bau.mutation.RemovedHistoriesIDs(); len(nodes) > 0 && !bau.mutation.HistoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   businessactivity.HistoriesTable,
			Columns: []string{businessactivity.HistoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(history.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bau.mutation.HistoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   businessactivity.HistoriesTable,
			Columns: []string{businessactivity.HistoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(history.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{businessactivity.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bau.mutation.done = true
	return n, nil
}

// BusinessActivityUpdateOne is the builder for updating a single BusinessActivity entity.
type BusinessActivityUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BusinessActivityMutation
}

// SetType sets the "type" field.
func (bauo *BusinessActivityUpdateOne) SetType(s string) *BusinessActivityUpdateOne {
	bauo.mutation.SetType(s)
	return bauo
}

// SetSubType sets the "sub_type" field.
func (bauo *BusinessActivityUpdateOne) SetSubType(s string) *BusinessActivityUpdateOne {
	bauo.mutation.SetSubType(s)
	return bauo
}

// SetTotal sets the "total" field.
func (bauo *BusinessActivityUpdateOne) SetTotal(f float64) *BusinessActivityUpdateOne {
	bauo.mutation.ResetTotal()
	bauo.mutation.SetTotal(f)
	return bauo
}

// AddTotal adds f to the "total" field.
func (bauo *BusinessActivityUpdateOne) AddTotal(f float64) *BusinessActivityUpdateOne {
	bauo.mutation.AddTotal(f)
	return bauo
}

// AddHistoryIDs adds the "histories" edge to the History entity by IDs.
func (bauo *BusinessActivityUpdateOne) AddHistoryIDs(ids ...int) *BusinessActivityUpdateOne {
	bauo.mutation.AddHistoryIDs(ids...)
	return bauo
}

// AddHistories adds the "histories" edges to the History entity.
func (bauo *BusinessActivityUpdateOne) AddHistories(h ...*History) *BusinessActivityUpdateOne {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return bauo.AddHistoryIDs(ids...)
}

// Mutation returns the BusinessActivityMutation object of the builder.
func (bauo *BusinessActivityUpdateOne) Mutation() *BusinessActivityMutation {
	return bauo.mutation
}

// ClearHistories clears all "histories" edges to the History entity.
func (bauo *BusinessActivityUpdateOne) ClearHistories() *BusinessActivityUpdateOne {
	bauo.mutation.ClearHistories()
	return bauo
}

// RemoveHistoryIDs removes the "histories" edge to History entities by IDs.
func (bauo *BusinessActivityUpdateOne) RemoveHistoryIDs(ids ...int) *BusinessActivityUpdateOne {
	bauo.mutation.RemoveHistoryIDs(ids...)
	return bauo
}

// RemoveHistories removes "histories" edges to History entities.
func (bauo *BusinessActivityUpdateOne) RemoveHistories(h ...*History) *BusinessActivityUpdateOne {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return bauo.RemoveHistoryIDs(ids...)
}

// Where appends a list predicates to the BusinessActivityUpdate builder.
func (bauo *BusinessActivityUpdateOne) Where(ps ...predicate.BusinessActivity) *BusinessActivityUpdateOne {
	bauo.mutation.Where(ps...)
	return bauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (bauo *BusinessActivityUpdateOne) Select(field string, fields ...string) *BusinessActivityUpdateOne {
	bauo.fields = append([]string{field}, fields...)
	return bauo
}

// Save executes the query and returns the updated BusinessActivity entity.
func (bauo *BusinessActivityUpdateOne) Save(ctx context.Context) (*BusinessActivity, error) {
	return withHooks(ctx, bauo.sqlSave, bauo.mutation, bauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bauo *BusinessActivityUpdateOne) SaveX(ctx context.Context) *BusinessActivity {
	node, err := bauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (bauo *BusinessActivityUpdateOne) Exec(ctx context.Context) error {
	_, err := bauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bauo *BusinessActivityUpdateOne) ExecX(ctx context.Context) {
	if err := bauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bauo *BusinessActivityUpdateOne) check() error {
	if v, ok := bauo.mutation.Total(); ok {
		if err := businessactivity.TotalValidator(v); err != nil {
			return &ValidationError{Name: "total", err: fmt.Errorf(`ent: validator failed for field "BusinessActivityId.total": %w`, err)}
		}
	}
	return nil
}

func (bauo *BusinessActivityUpdateOne) sqlSave(ctx context.Context) (_node *BusinessActivity, err error) {
	if err := bauo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(businessactivity.Table, businessactivity.Columns, sqlgraph.NewFieldSpec(businessactivity.FieldID, field.TypeInt))
	id, ok := bauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "BusinessActivityId.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := bauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, businessactivity.FieldID)
		for _, f := range fields {
			if !businessactivity.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != businessactivity.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := bauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bauo.mutation.GetType(); ok {
		_spec.SetField(businessactivity.FieldType, field.TypeString, value)
	}
	if value, ok := bauo.mutation.SubType(); ok {
		_spec.SetField(businessactivity.FieldSubType, field.TypeString, value)
	}
	if value, ok := bauo.mutation.Total(); ok {
		_spec.SetField(businessactivity.FieldTotal, field.TypeFloat64, value)
	}
	if value, ok := bauo.mutation.AddedTotal(); ok {
		_spec.AddField(businessactivity.FieldTotal, field.TypeFloat64, value)
	}
	if bauo.mutation.HistoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   businessactivity.HistoriesTable,
			Columns: []string{businessactivity.HistoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(history.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bauo.mutation.RemovedHistoriesIDs(); len(nodes) > 0 && !bauo.mutation.HistoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   businessactivity.HistoriesTable,
			Columns: []string{businessactivity.HistoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(history.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bauo.mutation.HistoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   businessactivity.HistoriesTable,
			Columns: []string{businessactivity.HistoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(history.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &BusinessActivity{config: bauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, bauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{businessactivity.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	bauo.mutation.done = true
	return _node, nil
}
