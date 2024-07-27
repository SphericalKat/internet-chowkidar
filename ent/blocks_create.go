// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gnulinuxindia/internet-chowkidar/ent/blocks"
)

// BlocksCreate is the builder for creating a Blocks entity.
type BlocksCreate struct {
	config
	mutation *BlocksMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (bc *BlocksCreate) SetCreatedAt(t time.Time) *BlocksCreate {
	bc.mutation.SetCreatedAt(t)
	return bc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bc *BlocksCreate) SetNillableCreatedAt(t *time.Time) *BlocksCreate {
	if t != nil {
		bc.SetCreatedAt(*t)
	}
	return bc
}

// SetUpdatedAt sets the "updated_at" field.
func (bc *BlocksCreate) SetUpdatedAt(t time.Time) *BlocksCreate {
	bc.mutation.SetUpdatedAt(t)
	return bc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (bc *BlocksCreate) SetNillableUpdatedAt(t *time.Time) *BlocksCreate {
	if t != nil {
		bc.SetUpdatedAt(*t)
	}
	return bc
}

// SetIP sets the "ip" field.
func (bc *BlocksCreate) SetIP(s string) *BlocksCreate {
	bc.mutation.SetIP(s)
	return bc
}

// SetDomain sets the "domain" field.
func (bc *BlocksCreate) SetDomain(s string) *BlocksCreate {
	bc.mutation.SetDomain(s)
	return bc
}

// Mutation returns the BlocksMutation object of the builder.
func (bc *BlocksCreate) Mutation() *BlocksMutation {
	return bc.mutation
}

// Save creates the Blocks in the database.
func (bc *BlocksCreate) Save(ctx context.Context) (*Blocks, error) {
	bc.defaults()
	return withHooks(ctx, bc.sqlSave, bc.mutation, bc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BlocksCreate) SaveX(ctx context.Context) *Blocks {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BlocksCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BlocksCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bc *BlocksCreate) defaults() {
	if _, ok := bc.mutation.CreatedAt(); !ok {
		v := blocks.DefaultCreatedAt()
		bc.mutation.SetCreatedAt(v)
	}
	if _, ok := bc.mutation.UpdatedAt(); !ok {
		v := blocks.DefaultUpdatedAt()
		bc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BlocksCreate) check() error {
	if _, ok := bc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Blocks.created_at"`)}
	}
	if _, ok := bc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Blocks.updated_at"`)}
	}
	if _, ok := bc.mutation.IP(); !ok {
		return &ValidationError{Name: "ip", err: errors.New(`ent: missing required field "Blocks.ip"`)}
	}
	if _, ok := bc.mutation.Domain(); !ok {
		return &ValidationError{Name: "domain", err: errors.New(`ent: missing required field "Blocks.domain"`)}
	}
	return nil
}

func (bc *BlocksCreate) sqlSave(ctx context.Context) (*Blocks, error) {
	if err := bc.check(); err != nil {
		return nil, err
	}
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	bc.mutation.id = &_node.ID
	bc.mutation.done = true
	return _node, nil
}

func (bc *BlocksCreate) createSpec() (*Blocks, *sqlgraph.CreateSpec) {
	var (
		_node = &Blocks{config: bc.config}
		_spec = sqlgraph.NewCreateSpec(blocks.Table, sqlgraph.NewFieldSpec(blocks.FieldID, field.TypeInt))
	)
	if value, ok := bc.mutation.CreatedAt(); ok {
		_spec.SetField(blocks.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := bc.mutation.UpdatedAt(); ok {
		_spec.SetField(blocks.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := bc.mutation.IP(); ok {
		_spec.SetField(blocks.FieldIP, field.TypeString, value)
		_node.IP = value
	}
	if value, ok := bc.mutation.Domain(); ok {
		_spec.SetField(blocks.FieldDomain, field.TypeString, value)
		_node.Domain = value
	}
	return _node, _spec
}

// BlocksCreateBulk is the builder for creating many Blocks entities in bulk.
type BlocksCreateBulk struct {
	config
	err      error
	builders []*BlocksCreate
}

// Save creates the Blocks entities in the database.
func (bcb *BlocksCreateBulk) Save(ctx context.Context) ([]*Blocks, error) {
	if bcb.err != nil {
		return nil, bcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Blocks, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BlocksMutation)
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
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BlocksCreateBulk) SaveX(ctx context.Context) []*Blocks {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BlocksCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BlocksCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}
