// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gnulinuxindia/internet-chowkidar/ent/blocks"
	"github.com/gnulinuxindia/internet-chowkidar/ent/predicate"
)

// BlocksUpdate is the builder for updating Blocks entities.
type BlocksUpdate struct {
	config
	hooks    []Hook
	mutation *BlocksMutation
}

// Where appends a list predicates to the BlocksUpdate builder.
func (bu *BlocksUpdate) Where(ps ...predicate.Blocks) *BlocksUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetUpdatedAt sets the "updated_at" field.
func (bu *BlocksUpdate) SetUpdatedAt(t time.Time) *BlocksUpdate {
	bu.mutation.SetUpdatedAt(t)
	return bu
}

// SetIP sets the "ip" field.
func (bu *BlocksUpdate) SetIP(s string) *BlocksUpdate {
	bu.mutation.SetIP(s)
	return bu
}

// SetNillableIP sets the "ip" field if the given value is not nil.
func (bu *BlocksUpdate) SetNillableIP(s *string) *BlocksUpdate {
	if s != nil {
		bu.SetIP(*s)
	}
	return bu
}

// SetDomain sets the "domain" field.
func (bu *BlocksUpdate) SetDomain(s string) *BlocksUpdate {
	bu.mutation.SetDomain(s)
	return bu
}

// SetNillableDomain sets the "domain" field if the given value is not nil.
func (bu *BlocksUpdate) SetNillableDomain(s *string) *BlocksUpdate {
	if s != nil {
		bu.SetDomain(*s)
	}
	return bu
}

// Mutation returns the BlocksMutation object of the builder.
func (bu *BlocksUpdate) Mutation() *BlocksMutation {
	return bu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BlocksUpdate) Save(ctx context.Context) (int, error) {
	bu.defaults()
	return withHooks(ctx, bu.sqlSave, bu.mutation, bu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BlocksUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BlocksUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BlocksUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bu *BlocksUpdate) defaults() {
	if _, ok := bu.mutation.UpdatedAt(); !ok {
		v := blocks.UpdateDefaultUpdatedAt()
		bu.mutation.SetUpdatedAt(v)
	}
}

func (bu *BlocksUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(blocks.Table, blocks.Columns, sqlgraph.NewFieldSpec(blocks.FieldID, field.TypeInt))
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.UpdatedAt(); ok {
		_spec.SetField(blocks.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := bu.mutation.IP(); ok {
		_spec.SetField(blocks.FieldIP, field.TypeString, value)
	}
	if value, ok := bu.mutation.Domain(); ok {
		_spec.SetField(blocks.FieldDomain, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{blocks.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bu.mutation.done = true
	return n, nil
}

// BlocksUpdateOne is the builder for updating a single Blocks entity.
type BlocksUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BlocksMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (buo *BlocksUpdateOne) SetUpdatedAt(t time.Time) *BlocksUpdateOne {
	buo.mutation.SetUpdatedAt(t)
	return buo
}

// SetIP sets the "ip" field.
func (buo *BlocksUpdateOne) SetIP(s string) *BlocksUpdateOne {
	buo.mutation.SetIP(s)
	return buo
}

// SetNillableIP sets the "ip" field if the given value is not nil.
func (buo *BlocksUpdateOne) SetNillableIP(s *string) *BlocksUpdateOne {
	if s != nil {
		buo.SetIP(*s)
	}
	return buo
}

// SetDomain sets the "domain" field.
func (buo *BlocksUpdateOne) SetDomain(s string) *BlocksUpdateOne {
	buo.mutation.SetDomain(s)
	return buo
}

// SetNillableDomain sets the "domain" field if the given value is not nil.
func (buo *BlocksUpdateOne) SetNillableDomain(s *string) *BlocksUpdateOne {
	if s != nil {
		buo.SetDomain(*s)
	}
	return buo
}

// Mutation returns the BlocksMutation object of the builder.
func (buo *BlocksUpdateOne) Mutation() *BlocksMutation {
	return buo.mutation
}

// Where appends a list predicates to the BlocksUpdate builder.
func (buo *BlocksUpdateOne) Where(ps ...predicate.Blocks) *BlocksUpdateOne {
	buo.mutation.Where(ps...)
	return buo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BlocksUpdateOne) Select(field string, fields ...string) *BlocksUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Blocks entity.
func (buo *BlocksUpdateOne) Save(ctx context.Context) (*Blocks, error) {
	buo.defaults()
	return withHooks(ctx, buo.sqlSave, buo.mutation, buo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BlocksUpdateOne) SaveX(ctx context.Context) *Blocks {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BlocksUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BlocksUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (buo *BlocksUpdateOne) defaults() {
	if _, ok := buo.mutation.UpdatedAt(); !ok {
		v := blocks.UpdateDefaultUpdatedAt()
		buo.mutation.SetUpdatedAt(v)
	}
}

func (buo *BlocksUpdateOne) sqlSave(ctx context.Context) (_node *Blocks, err error) {
	_spec := sqlgraph.NewUpdateSpec(blocks.Table, blocks.Columns, sqlgraph.NewFieldSpec(blocks.FieldID, field.TypeInt))
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Blocks.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, blocks.FieldID)
		for _, f := range fields {
			if !blocks.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != blocks.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.UpdatedAt(); ok {
		_spec.SetField(blocks.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := buo.mutation.IP(); ok {
		_spec.SetField(blocks.FieldIP, field.TypeString, value)
	}
	if value, ok := buo.mutation.Domain(); ok {
		_spec.SetField(blocks.FieldDomain, field.TypeString, value)
	}
	_node = &Blocks{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{blocks.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	buo.mutation.done = true
	return _node, nil
}
