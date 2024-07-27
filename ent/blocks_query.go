// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gnulinuxindia/internet-chowkidar/ent/blocks"
	"github.com/gnulinuxindia/internet-chowkidar/ent/isps"
	"github.com/gnulinuxindia/internet-chowkidar/ent/predicate"
	"github.com/gnulinuxindia/internet-chowkidar/ent/sites"
)

// BlocksQuery is the builder for querying Blocks entities.
type BlocksQuery struct {
	config
	ctx        *QueryContext
	order      []blocks.OrderOption
	inters     []Interceptor
	predicates []predicate.Blocks
	withSite   *SitesQuery
	withIsp    *IspsQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BlocksQuery builder.
func (bq *BlocksQuery) Where(ps ...predicate.Blocks) *BlocksQuery {
	bq.predicates = append(bq.predicates, ps...)
	return bq
}

// Limit the number of records to be returned by this query.
func (bq *BlocksQuery) Limit(limit int) *BlocksQuery {
	bq.ctx.Limit = &limit
	return bq
}

// Offset to start from.
func (bq *BlocksQuery) Offset(offset int) *BlocksQuery {
	bq.ctx.Offset = &offset
	return bq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bq *BlocksQuery) Unique(unique bool) *BlocksQuery {
	bq.ctx.Unique = &unique
	return bq
}

// Order specifies how the records should be ordered.
func (bq *BlocksQuery) Order(o ...blocks.OrderOption) *BlocksQuery {
	bq.order = append(bq.order, o...)
	return bq
}

// QuerySite chains the current query on the "site" edge.
func (bq *BlocksQuery) QuerySite() *SitesQuery {
	query := (&SitesClient{config: bq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(blocks.Table, blocks.FieldID, selector),
			sqlgraph.To(sites.Table, sites.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, blocks.SiteTable, blocks.SiteColumn),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryIsp chains the current query on the "isp" edge.
func (bq *BlocksQuery) QueryIsp() *IspsQuery {
	query := (&IspsClient{config: bq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(blocks.Table, blocks.FieldID, selector),
			sqlgraph.To(isps.Table, isps.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, blocks.IspTable, blocks.IspColumn),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Blocks entity from the query.
// Returns a *NotFoundError when no Blocks was found.
func (bq *BlocksQuery) First(ctx context.Context) (*Blocks, error) {
	nodes, err := bq.Limit(1).All(setContextOp(ctx, bq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{blocks.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bq *BlocksQuery) FirstX(ctx context.Context) *Blocks {
	node, err := bq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Blocks ID from the query.
// Returns a *NotFoundError when no Blocks ID was found.
func (bq *BlocksQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bq.Limit(1).IDs(setContextOp(ctx, bq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{blocks.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bq *BlocksQuery) FirstIDX(ctx context.Context) int {
	id, err := bq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Blocks entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Blocks entity is found.
// Returns a *NotFoundError when no Blocks entities are found.
func (bq *BlocksQuery) Only(ctx context.Context) (*Blocks, error) {
	nodes, err := bq.Limit(2).All(setContextOp(ctx, bq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{blocks.Label}
	default:
		return nil, &NotSingularError{blocks.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bq *BlocksQuery) OnlyX(ctx context.Context) *Blocks {
	node, err := bq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Blocks ID in the query.
// Returns a *NotSingularError when more than one Blocks ID is found.
// Returns a *NotFoundError when no entities are found.
func (bq *BlocksQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bq.Limit(2).IDs(setContextOp(ctx, bq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{blocks.Label}
	default:
		err = &NotSingularError{blocks.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bq *BlocksQuery) OnlyIDX(ctx context.Context) int {
	id, err := bq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BlocksSlice.
func (bq *BlocksQuery) All(ctx context.Context) ([]*Blocks, error) {
	ctx = setContextOp(ctx, bq.ctx, "All")
	if err := bq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Blocks, *BlocksQuery]()
	return withInterceptors[[]*Blocks](ctx, bq, qr, bq.inters)
}

// AllX is like All, but panics if an error occurs.
func (bq *BlocksQuery) AllX(ctx context.Context) []*Blocks {
	nodes, err := bq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Blocks IDs.
func (bq *BlocksQuery) IDs(ctx context.Context) (ids []int, err error) {
	if bq.ctx.Unique == nil && bq.path != nil {
		bq.Unique(true)
	}
	ctx = setContextOp(ctx, bq.ctx, "IDs")
	if err = bq.Select(blocks.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bq *BlocksQuery) IDsX(ctx context.Context) []int {
	ids, err := bq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bq *BlocksQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, bq.ctx, "Count")
	if err := bq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, bq, querierCount[*BlocksQuery](), bq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (bq *BlocksQuery) CountX(ctx context.Context) int {
	count, err := bq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bq *BlocksQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, bq.ctx, "Exist")
	switch _, err := bq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (bq *BlocksQuery) ExistX(ctx context.Context) bool {
	exist, err := bq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BlocksQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bq *BlocksQuery) Clone() *BlocksQuery {
	if bq == nil {
		return nil
	}
	return &BlocksQuery{
		config:     bq.config,
		ctx:        bq.ctx.Clone(),
		order:      append([]blocks.OrderOption{}, bq.order...),
		inters:     append([]Interceptor{}, bq.inters...),
		predicates: append([]predicate.Blocks{}, bq.predicates...),
		withSite:   bq.withSite.Clone(),
		withIsp:    bq.withIsp.Clone(),
		// clone intermediate query.
		sql:  bq.sql.Clone(),
		path: bq.path,
	}
}

// WithSite tells the query-builder to eager-load the nodes that are connected to
// the "site" edge. The optional arguments are used to configure the query builder of the edge.
func (bq *BlocksQuery) WithSite(opts ...func(*SitesQuery)) *BlocksQuery {
	query := (&SitesClient{config: bq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bq.withSite = query
	return bq
}

// WithIsp tells the query-builder to eager-load the nodes that are connected to
// the "isp" edge. The optional arguments are used to configure the query builder of the edge.
func (bq *BlocksQuery) WithIsp(opts ...func(*IspsQuery)) *BlocksQuery {
	query := (&IspsClient{config: bq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bq.withIsp = query
	return bq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Blocks.Query().
//		GroupBy(blocks.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (bq *BlocksQuery) GroupBy(field string, fields ...string) *BlocksGroupBy {
	bq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &BlocksGroupBy{build: bq}
	grbuild.flds = &bq.ctx.Fields
	grbuild.label = blocks.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.Blocks.Query().
//		Select(blocks.FieldCreatedAt).
//		Scan(ctx, &v)
func (bq *BlocksQuery) Select(fields ...string) *BlocksSelect {
	bq.ctx.Fields = append(bq.ctx.Fields, fields...)
	sbuild := &BlocksSelect{BlocksQuery: bq}
	sbuild.label = blocks.Label
	sbuild.flds, sbuild.scan = &bq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BlocksSelect configured with the given aggregations.
func (bq *BlocksQuery) Aggregate(fns ...AggregateFunc) *BlocksSelect {
	return bq.Select().Aggregate(fns...)
}

func (bq *BlocksQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range bq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, bq); err != nil {
				return err
			}
		}
	}
	for _, f := range bq.ctx.Fields {
		if !blocks.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if bq.path != nil {
		prev, err := bq.path(ctx)
		if err != nil {
			return err
		}
		bq.sql = prev
	}
	return nil
}

func (bq *BlocksQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Blocks, error) {
	var (
		nodes       = []*Blocks{}
		_spec       = bq.querySpec()
		loadedTypes = [2]bool{
			bq.withSite != nil,
			bq.withIsp != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Blocks).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Blocks{config: bq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := bq.withSite; query != nil {
		if err := bq.loadSite(ctx, query, nodes, nil,
			func(n *Blocks, e *Sites) { n.Edges.Site = e }); err != nil {
			return nil, err
		}
	}
	if query := bq.withIsp; query != nil {
		if err := bq.loadIsp(ctx, query, nodes, nil,
			func(n *Blocks, e *Isps) { n.Edges.Isp = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (bq *BlocksQuery) loadSite(ctx context.Context, query *SitesQuery, nodes []*Blocks, init func(*Blocks), assign func(*Blocks, *Sites)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Blocks)
	for i := range nodes {
		fk := nodes[i].SiteID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(sites.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "site_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (bq *BlocksQuery) loadIsp(ctx context.Context, query *IspsQuery, nodes []*Blocks, init func(*Blocks), assign func(*Blocks, *Isps)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Blocks)
	for i := range nodes {
		fk := nodes[i].IspID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(isps.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "isp_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (bq *BlocksQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bq.querySpec()
	_spec.Node.Columns = bq.ctx.Fields
	if len(bq.ctx.Fields) > 0 {
		_spec.Unique = bq.ctx.Unique != nil && *bq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, bq.driver, _spec)
}

func (bq *BlocksQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(blocks.Table, blocks.Columns, sqlgraph.NewFieldSpec(blocks.FieldID, field.TypeInt))
	_spec.From = bq.sql
	if unique := bq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if bq.path != nil {
		_spec.Unique = true
	}
	if fields := bq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, blocks.FieldID)
		for i := range fields {
			if fields[i] != blocks.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if bq.withSite != nil {
			_spec.Node.AddColumnOnce(blocks.FieldSiteID)
		}
		if bq.withIsp != nil {
			_spec.Node.AddColumnOnce(blocks.FieldIspID)
		}
	}
	if ps := bq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bq *BlocksQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bq.driver.Dialect())
	t1 := builder.Table(blocks.Table)
	columns := bq.ctx.Fields
	if len(columns) == 0 {
		columns = blocks.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bq.sql != nil {
		selector = bq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bq.ctx.Unique != nil && *bq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range bq.predicates {
		p(selector)
	}
	for _, p := range bq.order {
		p(selector)
	}
	if offset := bq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BlocksGroupBy is the group-by builder for Blocks entities.
type BlocksGroupBy struct {
	selector
	build *BlocksQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bgb *BlocksGroupBy) Aggregate(fns ...AggregateFunc) *BlocksGroupBy {
	bgb.fns = append(bgb.fns, fns...)
	return bgb
}

// Scan applies the selector query and scans the result into the given value.
func (bgb *BlocksGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bgb.build.ctx, "GroupBy")
	if err := bgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BlocksQuery, *BlocksGroupBy](ctx, bgb.build, bgb, bgb.build.inters, v)
}

func (bgb *BlocksGroupBy) sqlScan(ctx context.Context, root *BlocksQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(bgb.fns))
	for _, fn := range bgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*bgb.flds)+len(bgb.fns))
		for _, f := range *bgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*bgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// BlocksSelect is the builder for selecting fields of Blocks entities.
type BlocksSelect struct {
	*BlocksQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (bs *BlocksSelect) Aggregate(fns ...AggregateFunc) *BlocksSelect {
	bs.fns = append(bs.fns, fns...)
	return bs
}

// Scan applies the selector query and scans the result into the given value.
func (bs *BlocksSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bs.ctx, "Select")
	if err := bs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BlocksQuery, *BlocksSelect](ctx, bs.BlocksQuery, bs, bs.inters, v)
}

func (bs *BlocksSelect) sqlScan(ctx context.Context, root *BlocksQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(bs.fns))
	for _, fn := range bs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*bs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
