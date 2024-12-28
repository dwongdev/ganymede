// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/zibbp/ganymede/ent/live"
	"github.com/zibbp/ganymede/ent/livecategory"
	"github.com/zibbp/ganymede/ent/predicate"
)

// LiveCategoryQuery is the builder for querying LiveCategory entities.
type LiveCategoryQuery struct {
	config
	ctx        *QueryContext
	order      []livecategory.OrderOption
	inters     []Interceptor
	predicates []predicate.LiveCategory
	withLive   *LiveQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LiveCategoryQuery builder.
func (lcq *LiveCategoryQuery) Where(ps ...predicate.LiveCategory) *LiveCategoryQuery {
	lcq.predicates = append(lcq.predicates, ps...)
	return lcq
}

// Limit the number of records to be returned by this query.
func (lcq *LiveCategoryQuery) Limit(limit int) *LiveCategoryQuery {
	lcq.ctx.Limit = &limit
	return lcq
}

// Offset to start from.
func (lcq *LiveCategoryQuery) Offset(offset int) *LiveCategoryQuery {
	lcq.ctx.Offset = &offset
	return lcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (lcq *LiveCategoryQuery) Unique(unique bool) *LiveCategoryQuery {
	lcq.ctx.Unique = &unique
	return lcq
}

// Order specifies how the records should be ordered.
func (lcq *LiveCategoryQuery) Order(o ...livecategory.OrderOption) *LiveCategoryQuery {
	lcq.order = append(lcq.order, o...)
	return lcq
}

// QueryLive chains the current query on the "live" edge.
func (lcq *LiveCategoryQuery) QueryLive() *LiveQuery {
	query := (&LiveClient{config: lcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(livecategory.Table, livecategory.FieldID, selector),
			sqlgraph.To(live.Table, live.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, livecategory.LiveTable, livecategory.LiveColumn),
		)
		fromU = sqlgraph.SetNeighbors(lcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first LiveCategory entity from the query.
// Returns a *NotFoundError when no LiveCategory was found.
func (lcq *LiveCategoryQuery) First(ctx context.Context) (*LiveCategory, error) {
	nodes, err := lcq.Limit(1).All(setContextOp(ctx, lcq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{livecategory.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (lcq *LiveCategoryQuery) FirstX(ctx context.Context) *LiveCategory {
	node, err := lcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first LiveCategory ID from the query.
// Returns a *NotFoundError when no LiveCategory ID was found.
func (lcq *LiveCategoryQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = lcq.Limit(1).IDs(setContextOp(ctx, lcq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{livecategory.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (lcq *LiveCategoryQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := lcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single LiveCategory entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one LiveCategory entity is found.
// Returns a *NotFoundError when no LiveCategory entities are found.
func (lcq *LiveCategoryQuery) Only(ctx context.Context) (*LiveCategory, error) {
	nodes, err := lcq.Limit(2).All(setContextOp(ctx, lcq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{livecategory.Label}
	default:
		return nil, &NotSingularError{livecategory.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (lcq *LiveCategoryQuery) OnlyX(ctx context.Context) *LiveCategory {
	node, err := lcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only LiveCategory ID in the query.
// Returns a *NotSingularError when more than one LiveCategory ID is found.
// Returns a *NotFoundError when no entities are found.
func (lcq *LiveCategoryQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = lcq.Limit(2).IDs(setContextOp(ctx, lcq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{livecategory.Label}
	default:
		err = &NotSingularError{livecategory.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (lcq *LiveCategoryQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := lcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of LiveCategories.
func (lcq *LiveCategoryQuery) All(ctx context.Context) ([]*LiveCategory, error) {
	ctx = setContextOp(ctx, lcq.ctx, ent.OpQueryAll)
	if err := lcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*LiveCategory, *LiveCategoryQuery]()
	return withInterceptors[[]*LiveCategory](ctx, lcq, qr, lcq.inters)
}

// AllX is like All, but panics if an error occurs.
func (lcq *LiveCategoryQuery) AllX(ctx context.Context) []*LiveCategory {
	nodes, err := lcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of LiveCategory IDs.
func (lcq *LiveCategoryQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if lcq.ctx.Unique == nil && lcq.path != nil {
		lcq.Unique(true)
	}
	ctx = setContextOp(ctx, lcq.ctx, ent.OpQueryIDs)
	if err = lcq.Select(livecategory.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (lcq *LiveCategoryQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := lcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (lcq *LiveCategoryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, lcq.ctx, ent.OpQueryCount)
	if err := lcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, lcq, querierCount[*LiveCategoryQuery](), lcq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (lcq *LiveCategoryQuery) CountX(ctx context.Context) int {
	count, err := lcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (lcq *LiveCategoryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, lcq.ctx, ent.OpQueryExist)
	switch _, err := lcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (lcq *LiveCategoryQuery) ExistX(ctx context.Context) bool {
	exist, err := lcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LiveCategoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (lcq *LiveCategoryQuery) Clone() *LiveCategoryQuery {
	if lcq == nil {
		return nil
	}
	return &LiveCategoryQuery{
		config:     lcq.config,
		ctx:        lcq.ctx.Clone(),
		order:      append([]livecategory.OrderOption{}, lcq.order...),
		inters:     append([]Interceptor{}, lcq.inters...),
		predicates: append([]predicate.LiveCategory{}, lcq.predicates...),
		withLive:   lcq.withLive.Clone(),
		// clone intermediate query.
		sql:  lcq.sql.Clone(),
		path: lcq.path,
	}
}

// WithLive tells the query-builder to eager-load the nodes that are connected to
// the "live" edge. The optional arguments are used to configure the query builder of the edge.
func (lcq *LiveCategoryQuery) WithLive(opts ...func(*LiveQuery)) *LiveCategoryQuery {
	query := (&LiveClient{config: lcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	lcq.withLive = query
	return lcq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.LiveCategory.Query().
//		GroupBy(livecategory.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (lcq *LiveCategoryQuery) GroupBy(field string, fields ...string) *LiveCategoryGroupBy {
	lcq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &LiveCategoryGroupBy{build: lcq}
	grbuild.flds = &lcq.ctx.Fields
	grbuild.label = livecategory.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.LiveCategory.Query().
//		Select(livecategory.FieldName).
//		Scan(ctx, &v)
func (lcq *LiveCategoryQuery) Select(fields ...string) *LiveCategorySelect {
	lcq.ctx.Fields = append(lcq.ctx.Fields, fields...)
	sbuild := &LiveCategorySelect{LiveCategoryQuery: lcq}
	sbuild.label = livecategory.Label
	sbuild.flds, sbuild.scan = &lcq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a LiveCategorySelect configured with the given aggregations.
func (lcq *LiveCategoryQuery) Aggregate(fns ...AggregateFunc) *LiveCategorySelect {
	return lcq.Select().Aggregate(fns...)
}

func (lcq *LiveCategoryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range lcq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, lcq); err != nil {
				return err
			}
		}
	}
	for _, f := range lcq.ctx.Fields {
		if !livecategory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if lcq.path != nil {
		prev, err := lcq.path(ctx)
		if err != nil {
			return err
		}
		lcq.sql = prev
	}
	return nil
}

func (lcq *LiveCategoryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*LiveCategory, error) {
	var (
		nodes       = []*LiveCategory{}
		withFKs     = lcq.withFKs
		_spec       = lcq.querySpec()
		loadedTypes = [1]bool{
			lcq.withLive != nil,
		}
	)
	if lcq.withLive != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, livecategory.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*LiveCategory).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &LiveCategory{config: lcq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, lcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := lcq.withLive; query != nil {
		if err := lcq.loadLive(ctx, query, nodes, nil,
			func(n *LiveCategory, e *Live) { n.Edges.Live = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (lcq *LiveCategoryQuery) loadLive(ctx context.Context, query *LiveQuery, nodes []*LiveCategory, init func(*LiveCategory), assign func(*LiveCategory, *Live)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*LiveCategory)
	for i := range nodes {
		if nodes[i].live_id == nil {
			continue
		}
		fk := *nodes[i].live_id
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(live.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "live_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (lcq *LiveCategoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := lcq.querySpec()
	_spec.Node.Columns = lcq.ctx.Fields
	if len(lcq.ctx.Fields) > 0 {
		_spec.Unique = lcq.ctx.Unique != nil && *lcq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, lcq.driver, _spec)
}

func (lcq *LiveCategoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(livecategory.Table, livecategory.Columns, sqlgraph.NewFieldSpec(livecategory.FieldID, field.TypeUUID))
	_spec.From = lcq.sql
	if unique := lcq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if lcq.path != nil {
		_spec.Unique = true
	}
	if fields := lcq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, livecategory.FieldID)
		for i := range fields {
			if fields[i] != livecategory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := lcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := lcq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := lcq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := lcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (lcq *LiveCategoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(lcq.driver.Dialect())
	t1 := builder.Table(livecategory.Table)
	columns := lcq.ctx.Fields
	if len(columns) == 0 {
		columns = livecategory.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if lcq.sql != nil {
		selector = lcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if lcq.ctx.Unique != nil && *lcq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range lcq.predicates {
		p(selector)
	}
	for _, p := range lcq.order {
		p(selector)
	}
	if offset := lcq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := lcq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// LiveCategoryGroupBy is the group-by builder for LiveCategory entities.
type LiveCategoryGroupBy struct {
	selector
	build *LiveCategoryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lcgb *LiveCategoryGroupBy) Aggregate(fns ...AggregateFunc) *LiveCategoryGroupBy {
	lcgb.fns = append(lcgb.fns, fns...)
	return lcgb
}

// Scan applies the selector query and scans the result into the given value.
func (lcgb *LiveCategoryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lcgb.build.ctx, ent.OpQueryGroupBy)
	if err := lcgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LiveCategoryQuery, *LiveCategoryGroupBy](ctx, lcgb.build, lcgb, lcgb.build.inters, v)
}

func (lcgb *LiveCategoryGroupBy) sqlScan(ctx context.Context, root *LiveCategoryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(lcgb.fns))
	for _, fn := range lcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*lcgb.flds)+len(lcgb.fns))
		for _, f := range *lcgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*lcgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lcgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// LiveCategorySelect is the builder for selecting fields of LiveCategory entities.
type LiveCategorySelect struct {
	*LiveCategoryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (lcs *LiveCategorySelect) Aggregate(fns ...AggregateFunc) *LiveCategorySelect {
	lcs.fns = append(lcs.fns, fns...)
	return lcs
}

// Scan applies the selector query and scans the result into the given value.
func (lcs *LiveCategorySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lcs.ctx, ent.OpQuerySelect)
	if err := lcs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LiveCategoryQuery, *LiveCategorySelect](ctx, lcs.LiveCategoryQuery, lcs, lcs.inters, v)
}

func (lcs *LiveCategorySelect) sqlScan(ctx context.Context, root *LiveCategoryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(lcs.fns))
	for _, fn := range lcs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*lcs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
