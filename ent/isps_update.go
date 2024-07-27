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
	"github.com/gnulinuxindia/internet-chowkidar/ent/isps"
	"github.com/gnulinuxindia/internet-chowkidar/ent/predicate"
)

// IspsUpdate is the builder for updating Isps entities.
type IspsUpdate struct {
	config
	hooks    []Hook
	mutation *IspsMutation
}

// Where appends a list predicates to the IspsUpdate builder.
func (iu *IspsUpdate) Where(ps ...predicate.Isps) *IspsUpdate {
	iu.mutation.Where(ps...)
	return iu
}

// SetUpdatedAt sets the "updated_at" field.
func (iu *IspsUpdate) SetUpdatedAt(t time.Time) *IspsUpdate {
	iu.mutation.SetUpdatedAt(t)
	return iu
}

// SetLatitude sets the "latitude" field.
func (iu *IspsUpdate) SetLatitude(f float64) *IspsUpdate {
	iu.mutation.ResetLatitude()
	iu.mutation.SetLatitude(f)
	return iu
}

// SetNillableLatitude sets the "latitude" field if the given value is not nil.
func (iu *IspsUpdate) SetNillableLatitude(f *float64) *IspsUpdate {
	if f != nil {
		iu.SetLatitude(*f)
	}
	return iu
}

// AddLatitude adds f to the "latitude" field.
func (iu *IspsUpdate) AddLatitude(f float64) *IspsUpdate {
	iu.mutation.AddLatitude(f)
	return iu
}

// SetLongitude sets the "longitude" field.
func (iu *IspsUpdate) SetLongitude(f float64) *IspsUpdate {
	iu.mutation.ResetLongitude()
	iu.mutation.SetLongitude(f)
	return iu
}

// SetNillableLongitude sets the "longitude" field if the given value is not nil.
func (iu *IspsUpdate) SetNillableLongitude(f *float64) *IspsUpdate {
	if f != nil {
		iu.SetLongitude(*f)
	}
	return iu
}

// AddLongitude adds f to the "longitude" field.
func (iu *IspsUpdate) AddLongitude(f float64) *IspsUpdate {
	iu.mutation.AddLongitude(f)
	return iu
}

// SetName sets the "name" field.
func (iu *IspsUpdate) SetName(s string) *IspsUpdate {
	iu.mutation.SetName(s)
	return iu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (iu *IspsUpdate) SetNillableName(s *string) *IspsUpdate {
	if s != nil {
		iu.SetName(*s)
	}
	return iu
}

// AddIspBlockIDs adds the "isp_blocks" edge to the Blocks entity by IDs.
func (iu *IspsUpdate) AddIspBlockIDs(ids ...int) *IspsUpdate {
	iu.mutation.AddIspBlockIDs(ids...)
	return iu
}

// AddIspBlocks adds the "isp_blocks" edges to the Blocks entity.
func (iu *IspsUpdate) AddIspBlocks(b ...*Blocks) *IspsUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return iu.AddIspBlockIDs(ids...)
}

// Mutation returns the IspsMutation object of the builder.
func (iu *IspsUpdate) Mutation() *IspsMutation {
	return iu.mutation
}

// ClearIspBlocks clears all "isp_blocks" edges to the Blocks entity.
func (iu *IspsUpdate) ClearIspBlocks() *IspsUpdate {
	iu.mutation.ClearIspBlocks()
	return iu
}

// RemoveIspBlockIDs removes the "isp_blocks" edge to Blocks entities by IDs.
func (iu *IspsUpdate) RemoveIspBlockIDs(ids ...int) *IspsUpdate {
	iu.mutation.RemoveIspBlockIDs(ids...)
	return iu
}

// RemoveIspBlocks removes "isp_blocks" edges to Blocks entities.
func (iu *IspsUpdate) RemoveIspBlocks(b ...*Blocks) *IspsUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return iu.RemoveIspBlockIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *IspsUpdate) Save(ctx context.Context) (int, error) {
	iu.defaults()
	return withHooks(ctx, iu.sqlSave, iu.mutation, iu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iu *IspsUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *IspsUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *IspsUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iu *IspsUpdate) defaults() {
	if _, ok := iu.mutation.UpdatedAt(); !ok {
		v := isps.UpdateDefaultUpdatedAt()
		iu.mutation.SetUpdatedAt(v)
	}
}

func (iu *IspsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(isps.Table, isps.Columns, sqlgraph.NewFieldSpec(isps.FieldID, field.TypeInt))
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.UpdatedAt(); ok {
		_spec.SetField(isps.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := iu.mutation.Latitude(); ok {
		_spec.SetField(isps.FieldLatitude, field.TypeFloat64, value)
	}
	if value, ok := iu.mutation.AddedLatitude(); ok {
		_spec.AddField(isps.FieldLatitude, field.TypeFloat64, value)
	}
	if value, ok := iu.mutation.Longitude(); ok {
		_spec.SetField(isps.FieldLongitude, field.TypeFloat64, value)
	}
	if value, ok := iu.mutation.AddedLongitude(); ok {
		_spec.AddField(isps.FieldLongitude, field.TypeFloat64, value)
	}
	if value, ok := iu.mutation.Name(); ok {
		_spec.SetField(isps.FieldName, field.TypeString, value)
	}
	if iu.mutation.IspBlocksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   isps.IspBlocksTable,
			Columns: []string{isps.IspBlocksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blocks.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.RemovedIspBlocksIDs(); len(nodes) > 0 && !iu.mutation.IspBlocksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   isps.IspBlocksTable,
			Columns: []string{isps.IspBlocksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blocks.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.IspBlocksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   isps.IspBlocksTable,
			Columns: []string{isps.IspBlocksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blocks.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{isps.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	iu.mutation.done = true
	return n, nil
}

// IspsUpdateOne is the builder for updating a single Isps entity.
type IspsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *IspsMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (iuo *IspsUpdateOne) SetUpdatedAt(t time.Time) *IspsUpdateOne {
	iuo.mutation.SetUpdatedAt(t)
	return iuo
}

// SetLatitude sets the "latitude" field.
func (iuo *IspsUpdateOne) SetLatitude(f float64) *IspsUpdateOne {
	iuo.mutation.ResetLatitude()
	iuo.mutation.SetLatitude(f)
	return iuo
}

// SetNillableLatitude sets the "latitude" field if the given value is not nil.
func (iuo *IspsUpdateOne) SetNillableLatitude(f *float64) *IspsUpdateOne {
	if f != nil {
		iuo.SetLatitude(*f)
	}
	return iuo
}

// AddLatitude adds f to the "latitude" field.
func (iuo *IspsUpdateOne) AddLatitude(f float64) *IspsUpdateOne {
	iuo.mutation.AddLatitude(f)
	return iuo
}

// SetLongitude sets the "longitude" field.
func (iuo *IspsUpdateOne) SetLongitude(f float64) *IspsUpdateOne {
	iuo.mutation.ResetLongitude()
	iuo.mutation.SetLongitude(f)
	return iuo
}

// SetNillableLongitude sets the "longitude" field if the given value is not nil.
func (iuo *IspsUpdateOne) SetNillableLongitude(f *float64) *IspsUpdateOne {
	if f != nil {
		iuo.SetLongitude(*f)
	}
	return iuo
}

// AddLongitude adds f to the "longitude" field.
func (iuo *IspsUpdateOne) AddLongitude(f float64) *IspsUpdateOne {
	iuo.mutation.AddLongitude(f)
	return iuo
}

// SetName sets the "name" field.
func (iuo *IspsUpdateOne) SetName(s string) *IspsUpdateOne {
	iuo.mutation.SetName(s)
	return iuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (iuo *IspsUpdateOne) SetNillableName(s *string) *IspsUpdateOne {
	if s != nil {
		iuo.SetName(*s)
	}
	return iuo
}

// AddIspBlockIDs adds the "isp_blocks" edge to the Blocks entity by IDs.
func (iuo *IspsUpdateOne) AddIspBlockIDs(ids ...int) *IspsUpdateOne {
	iuo.mutation.AddIspBlockIDs(ids...)
	return iuo
}

// AddIspBlocks adds the "isp_blocks" edges to the Blocks entity.
func (iuo *IspsUpdateOne) AddIspBlocks(b ...*Blocks) *IspsUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return iuo.AddIspBlockIDs(ids...)
}

// Mutation returns the IspsMutation object of the builder.
func (iuo *IspsUpdateOne) Mutation() *IspsMutation {
	return iuo.mutation
}

// ClearIspBlocks clears all "isp_blocks" edges to the Blocks entity.
func (iuo *IspsUpdateOne) ClearIspBlocks() *IspsUpdateOne {
	iuo.mutation.ClearIspBlocks()
	return iuo
}

// RemoveIspBlockIDs removes the "isp_blocks" edge to Blocks entities by IDs.
func (iuo *IspsUpdateOne) RemoveIspBlockIDs(ids ...int) *IspsUpdateOne {
	iuo.mutation.RemoveIspBlockIDs(ids...)
	return iuo
}

// RemoveIspBlocks removes "isp_blocks" edges to Blocks entities.
func (iuo *IspsUpdateOne) RemoveIspBlocks(b ...*Blocks) *IspsUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return iuo.RemoveIspBlockIDs(ids...)
}

// Where appends a list predicates to the IspsUpdate builder.
func (iuo *IspsUpdateOne) Where(ps ...predicate.Isps) *IspsUpdateOne {
	iuo.mutation.Where(ps...)
	return iuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iuo *IspsUpdateOne) Select(field string, fields ...string) *IspsUpdateOne {
	iuo.fields = append([]string{field}, fields...)
	return iuo
}

// Save executes the query and returns the updated Isps entity.
func (iuo *IspsUpdateOne) Save(ctx context.Context) (*Isps, error) {
	iuo.defaults()
	return withHooks(ctx, iuo.sqlSave, iuo.mutation, iuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *IspsUpdateOne) SaveX(ctx context.Context) *Isps {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *IspsUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *IspsUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iuo *IspsUpdateOne) defaults() {
	if _, ok := iuo.mutation.UpdatedAt(); !ok {
		v := isps.UpdateDefaultUpdatedAt()
		iuo.mutation.SetUpdatedAt(v)
	}
}

func (iuo *IspsUpdateOne) sqlSave(ctx context.Context) (_node *Isps, err error) {
	_spec := sqlgraph.NewUpdateSpec(isps.Table, isps.Columns, sqlgraph.NewFieldSpec(isps.FieldID, field.TypeInt))
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Isps.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := iuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, isps.FieldID)
		for _, f := range fields {
			if !isps.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != isps.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iuo.mutation.UpdatedAt(); ok {
		_spec.SetField(isps.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := iuo.mutation.Latitude(); ok {
		_spec.SetField(isps.FieldLatitude, field.TypeFloat64, value)
	}
	if value, ok := iuo.mutation.AddedLatitude(); ok {
		_spec.AddField(isps.FieldLatitude, field.TypeFloat64, value)
	}
	if value, ok := iuo.mutation.Longitude(); ok {
		_spec.SetField(isps.FieldLongitude, field.TypeFloat64, value)
	}
	if value, ok := iuo.mutation.AddedLongitude(); ok {
		_spec.AddField(isps.FieldLongitude, field.TypeFloat64, value)
	}
	if value, ok := iuo.mutation.Name(); ok {
		_spec.SetField(isps.FieldName, field.TypeString, value)
	}
	if iuo.mutation.IspBlocksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   isps.IspBlocksTable,
			Columns: []string{isps.IspBlocksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blocks.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.RemovedIspBlocksIDs(); len(nodes) > 0 && !iuo.mutation.IspBlocksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   isps.IspBlocksTable,
			Columns: []string{isps.IspBlocksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blocks.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.IspBlocksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   isps.IspBlocksTable,
			Columns: []string{isps.IspBlocksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blocks.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Isps{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{isps.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	iuo.mutation.done = true
	return _node, nil
}
