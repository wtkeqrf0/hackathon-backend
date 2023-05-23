// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wtkeqrf0/while.act/ent/equipment"
)

// EquipmentCreate is the builder for creating a Equipment entity.
type EquipmentCreate struct {
	config
	mutation *EquipmentMutation
	hooks    []Hook
}

// SetType sets the "type" field.
func (ec *EquipmentCreate) SetType(s string) *EquipmentCreate {
	ec.mutation.SetType(s)
	return ec
}

// SetAvgPriceDol sets the "avg_price_dol" field.
func (ec *EquipmentCreate) SetAvgPriceDol(i int) *EquipmentCreate {
	ec.mutation.SetAvgPriceDol(i)
	return ec
}

// SetAvgPriceRub sets the "avg_price_rub" field.
func (ec *EquipmentCreate) SetAvgPriceRub(i int) *EquipmentCreate {
	ec.mutation.SetAvgPriceRub(i)
	return ec
}

// Mutation returns the EquipmentMutation object of the builder.
func (ec *EquipmentCreate) Mutation() *EquipmentMutation {
	return ec.mutation
}

// Save creates the Equipment in the database.
func (ec *EquipmentCreate) Save(ctx context.Context) (*Equipment, error) {
	return withHooks(ctx, ec.sqlSave, ec.mutation, ec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EquipmentCreate) SaveX(ctx context.Context) *Equipment {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ec *EquipmentCreate) Exec(ctx context.Context) error {
	_, err := ec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ec *EquipmentCreate) ExecX(ctx context.Context) {
	if err := ec.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ec *EquipmentCreate) check() error {
	if _, ok := ec.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Equipment.type"`)}
	}
	if _, ok := ec.mutation.AvgPriceDol(); !ok {
		return &ValidationError{Name: "avg_price_dol", err: errors.New(`ent: missing required field "Equipment.avg_price_dol"`)}
	}
	if _, ok := ec.mutation.AvgPriceRub(); !ok {
		return &ValidationError{Name: "avg_price_rub", err: errors.New(`ent: missing required field "Equipment.avg_price_rub"`)}
	}
	return nil
}

func (ec *EquipmentCreate) sqlSave(ctx context.Context) (*Equipment, error) {
	if err := ec.check(); err != nil {
		return nil, err
	}
	_node, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ec.mutation.id = &_node.ID
	ec.mutation.done = true
	return _node, nil
}

func (ec *EquipmentCreate) createSpec() (*Equipment, *sqlgraph.CreateSpec) {
	var (
		_node = &Equipment{config: ec.config}
		_spec = sqlgraph.NewCreateSpec(equipment.Table, sqlgraph.NewFieldSpec(equipment.FieldID, field.TypeInt))
	)
	if value, ok := ec.mutation.GetType(); ok {
		_spec.SetField(equipment.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if value, ok := ec.mutation.AvgPriceDol(); ok {
		_spec.SetField(equipment.FieldAvgPriceDol, field.TypeInt, value)
		_node.AvgPriceDol = value
	}
	if value, ok := ec.mutation.AvgPriceRub(); ok {
		_spec.SetField(equipment.FieldAvgPriceRub, field.TypeInt, value)
		_node.AvgPriceRub = value
	}
	return _node, _spec
}

// EquipmentCreateBulk is the builder for creating many Equipment entities in bulk.
type EquipmentCreateBulk struct {
	config
	builders []*EquipmentCreate
}

// Save creates the Equipment entities in the database.
func (ecb *EquipmentCreateBulk) Save(ctx context.Context) ([]*Equipment, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*Equipment, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EquipmentMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ecb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ecb *EquipmentCreateBulk) SaveX(ctx context.Context) []*Equipment {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecb *EquipmentCreateBulk) Exec(ctx context.Context) error {
	_, err := ecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecb *EquipmentCreateBulk) ExecX(ctx context.Context) {
	if err := ecb.Exec(ctx); err != nil {
		panic(err)
	}
}
