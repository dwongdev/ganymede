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
	"github.com/zibbp/ganymede/ent/livetitleregex"
	"github.com/zibbp/ganymede/ent/predicate"
)

// LiveTitleRegexQuery is the builder for querying LiveTitleRegex entities.
type LiveTitleRegexQuery struct {
	config
	ctx        *QueryContext
	order      []livetitleregex.OrderOption
	inters     []Interceptor
	predicates []predicate.LiveTitleRegex
	withLive   *LiveQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LiveTitleRegexQuery builder.
func (ltrq *LiveTitleRegexQuery) Where(ps ...predicate.LiveTitleRegex) *LiveTitleRegexQuery {
	ltrq.predicates = append(ltrq.predicates, ps...)
	return ltrq
}

// Limit the number of records to be returned by this query.
func (ltrq *LiveTitleRegexQuery) Limit(limit int) *LiveTitleRegexQuery {
	ltrq.ctx.Limit = &limit
	return ltrq
}

// Offset to start from.
func (ltrq *LiveTitleRegexQuery) Offset(offset int) *LiveTitleRegexQuery {
	ltrq.ctx.Offset = &offset
	return ltrq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ltrq *LiveTitleRegexQuery) Unique(unique bool) *LiveTitleRegexQuery {
	ltrq.ctx.Unique = &unique
	return ltrq
}

// Order specifies how the records should be ordered.
func (ltrq *LiveTitleRegexQuery) Order(o ...livetitleregex.OrderOption) *LiveTitleRegexQuery {
	ltrq.order = append(ltrq.order, o...)
	return ltrq
}

// QueryLive chains the current query on the "live" edge.
func (ltrq *LiveTitleRegexQuery) QueryLive() *LiveQuery {
	query := (&LiveClient{config: ltrq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ltrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ltrq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(livetitleregex.Table, livetitleregex.FieldID, selector),
			sqlgraph.To(live.Table, live.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, livetitleregex.LiveTable, livetitleregex.LiveColumn),
		)
		fromU = sqlgraph.SetNeighbors(ltrq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first LiveTitleRegex entity from the query.
// Returns a *NotFoundError when no LiveTitleRegex was found.
func (ltrq *LiveTitleRegexQuery) First(ctx context.Context) (*LiveTitleRegex, error) {
	nodes, err := ltrq.Limit(1).All(setContextOp(ctx, ltrq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{livetitleregex.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ltrq *LiveTitleRegexQuery) FirstX(ctx context.Context) *LiveTitleRegex {
	node, err := ltrq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first LiveTitleRegex ID from the query.
// Returns a *NotFoundError when no LiveTitleRegex ID was found.
func (ltrq *LiveTitleRegexQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ltrq.Limit(1).IDs(setContextOp(ctx, ltrq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{livetitleregex.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ltrq *LiveTitleRegexQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := ltrq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single LiveTitleRegex entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one LiveTitleRegex entity is found.
// Returns a *NotFoundError when no LiveTitleRegex entities are found.
func (ltrq *LiveTitleRegexQuery) Only(ctx context.Context) (*LiveTitleRegex, error) {
	nodes, err := ltrq.Limit(2).All(setContextOp(ctx, ltrq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{livetitleregex.Label}
	default:
		return nil, &NotSingularError{livetitleregex.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ltrq *LiveTitleRegexQuery) OnlyX(ctx context.Context) *LiveTitleRegex {
	node, err := ltrq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only LiveTitleRegex ID in the query.
// Returns a *NotSingularError when more than one LiveTitleRegex ID is found.
// Returns a *NotFoundError when no entities are found.
func (ltrq *LiveTitleRegexQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ltrq.Limit(2).IDs(setContextOp(ctx, ltrq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{livetitleregex.Label}
	default:
		err = &NotSingularError{livetitleregex.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ltrq *LiveTitleRegexQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := ltrq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of LiveTitleRegexes.
func (ltrq *LiveTitleRegexQuery) All(ctx context.Context) ([]*LiveTitleRegex, error) {
	ctx = setContextOp(ctx, ltrq.ctx, ent.OpQueryAll)
	if err := ltrq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*LiveTitleRegex, *LiveTitleRegexQuery]()
	return withInterceptors[[]*LiveTitleRegex](ctx, ltrq, qr, ltrq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ltrq *LiveTitleRegexQuery) AllX(ctx context.Context) []*LiveTitleRegex {
	nodes, err := ltrq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of LiveTitleRegex IDs.
func (ltrq *LiveTitleRegexQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if ltrq.ctx.Unique == nil && ltrq.path != nil {
		ltrq.Unique(true)
	}
	ctx = setContextOp(ctx, ltrq.ctx, ent.OpQueryIDs)
	if err = ltrq.Select(livetitleregex.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ltrq *LiveTitleRegexQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := ltrq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ltrq *LiveTitleRegexQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ltrq.ctx, ent.OpQueryCount)
	if err := ltrq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ltrq, querierCount[*LiveTitleRegexQuery](), ltrq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ltrq *LiveTitleRegexQuery) CountX(ctx context.Context) int {
	count, err := ltrq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ltrq *LiveTitleRegexQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ltrq.ctx, ent.OpQueryExist)
	switch _, err := ltrq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ltrq *LiveTitleRegexQuery) ExistX(ctx context.Context) bool {
	exist, err := ltrq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LiveTitleRegexQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ltrq *LiveTitleRegexQuery) Clone() *LiveTitleRegexQuery {
	if ltrq == nil {
		return nil
	}
	return &LiveTitleRegexQuery{
		config:     ltrq.config,
		ctx:        ltrq.ctx.Clone(),
		order:      append([]livetitleregex.OrderOption{}, ltrq.order...),
		inters:     append([]Interceptor{}, ltrq.inters...),
		predicates: append([]predicate.LiveTitleRegex{}, ltrq.predicates...),
		withLive:   ltrq.withLive.Clone(),
		// clone intermediate query.
		sql:  ltrq.sql.Clone(),
		path: ltrq.path,
	}
}

// WithLive tells the query-builder to eager-load the nodes that are connected to
// the "live" edge. The optional arguments are used to configure the query builder of the edge.
func (ltrq *LiveTitleRegexQuery) WithLive(opts ...func(*LiveQuery)) *LiveTitleRegexQuery {
	query := (&LiveClient{config: ltrq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ltrq.withLive = query
	return ltrq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Negative bool `json:"negative"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.LiveTitleRegex.Query().
//		GroupBy(livetitleregex.FieldNegative).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ltrq *LiveTitleRegexQuery) GroupBy(field string, fields ...string) *LiveTitleRegexGroupBy {
	ltrq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &LiveTitleRegexGroupBy{build: ltrq}
	grbuild.flds = &ltrq.ctx.Fields
	grbuild.label = livetitleregex.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Negative bool `json:"negative"`
//	}
//
//	client.LiveTitleRegex.Query().
//		Select(livetitleregex.FieldNegative).
//		Scan(ctx, &v)
func (ltrq *LiveTitleRegexQuery) Select(fields ...string) *LiveTitleRegexSelect {
	ltrq.ctx.Fields = append(ltrq.ctx.Fields, fields...)
	sbuild := &LiveTitleRegexSelect{LiveTitleRegexQuery: ltrq}
	sbuild.label = livetitleregex.Label
	sbuild.flds, sbuild.scan = &ltrq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a LiveTitleRegexSelect configured with the given aggregations.
func (ltrq *LiveTitleRegexQuery) Aggregate(fns ...AggregateFunc) *LiveTitleRegexSelect {
	return ltrq.Select().Aggregate(fns...)
}

func (ltrq *LiveTitleRegexQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ltrq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ltrq); err != nil {
				return err
			}
		}
	}
	for _, f := range ltrq.ctx.Fields {
		if !livetitleregex.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ltrq.path != nil {
		prev, err := ltrq.path(ctx)
		if err != nil {
			return err
		}
		ltrq.sql = prev
	}
	return nil
}

func (ltrq *LiveTitleRegexQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*LiveTitleRegex, error) {
	var (
		nodes       = []*LiveTitleRegex{}
		withFKs     = ltrq.withFKs
		_spec       = ltrq.querySpec()
		loadedTypes = [1]bool{
			ltrq.withLive != nil,
		}
	)
	if ltrq.withLive != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, livetitleregex.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*LiveTitleRegex).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &LiveTitleRegex{config: ltrq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ltrq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ltrq.withLive; query != nil {
		if err := ltrq.loadLive(ctx, query, nodes, nil,
			func(n *LiveTitleRegex, e *Live) { n.Edges.Live = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ltrq *LiveTitleRegexQuery) loadLive(ctx context.Context, query *LiveQuery, nodes []*LiveTitleRegex, init func(*LiveTitleRegex), assign func(*LiveTitleRegex, *Live)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*LiveTitleRegex)
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

func (ltrq *LiveTitleRegexQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ltrq.querySpec()
	_spec.Node.Columns = ltrq.ctx.Fields
	if len(ltrq.ctx.Fields) > 0 {
		_spec.Unique = ltrq.ctx.Unique != nil && *ltrq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ltrq.driver, _spec)
}

func (ltrq *LiveTitleRegexQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(livetitleregex.Table, livetitleregex.Columns, sqlgraph.NewFieldSpec(livetitleregex.FieldID, field.TypeUUID))
	_spec.From = ltrq.sql
	if unique := ltrq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ltrq.path != nil {
		_spec.Unique = true
	}
	if fields := ltrq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, livetitleregex.FieldID)
		for i := range fields {
			if fields[i] != livetitleregex.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ltrq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ltrq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ltrq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ltrq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ltrq *LiveTitleRegexQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ltrq.driver.Dialect())
	t1 := builder.Table(livetitleregex.Table)
	columns := ltrq.ctx.Fields
	if len(columns) == 0 {
		columns = livetitleregex.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ltrq.sql != nil {
		selector = ltrq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ltrq.ctx.Unique != nil && *ltrq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ltrq.predicates {
		p(selector)
	}
	for _, p := range ltrq.order {
		p(selector)
	}
	if offset := ltrq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ltrq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// LiveTitleRegexGroupBy is the group-by builder for LiveTitleRegex entities.
type LiveTitleRegexGroupBy struct {
	selector
	build *LiveTitleRegexQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ltrgb *LiveTitleRegexGroupBy) Aggregate(fns ...AggregateFunc) *LiveTitleRegexGroupBy {
	ltrgb.fns = append(ltrgb.fns, fns...)
	return ltrgb
}

// Scan applies the selector query and scans the result into the given value.
func (ltrgb *LiveTitleRegexGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ltrgb.build.ctx, ent.OpQueryGroupBy)
	if err := ltrgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LiveTitleRegexQuery, *LiveTitleRegexGroupBy](ctx, ltrgb.build, ltrgb, ltrgb.build.inters, v)
}

func (ltrgb *LiveTitleRegexGroupBy) sqlScan(ctx context.Context, root *LiveTitleRegexQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ltrgb.fns))
	for _, fn := range ltrgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ltrgb.flds)+len(ltrgb.fns))
		for _, f := range *ltrgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ltrgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ltrgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// LiveTitleRegexSelect is the builder for selecting fields of LiveTitleRegex entities.
type LiveTitleRegexSelect struct {
	*LiveTitleRegexQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ltrs *LiveTitleRegexSelect) Aggregate(fns ...AggregateFunc) *LiveTitleRegexSelect {
	ltrs.fns = append(ltrs.fns, fns...)
	return ltrs
}

// Scan applies the selector query and scans the result into the given value.
func (ltrs *LiveTitleRegexSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ltrs.ctx, ent.OpQuerySelect)
	if err := ltrs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LiveTitleRegexQuery, *LiveTitleRegexSelect](ctx, ltrs.LiveTitleRegexQuery, ltrs, ltrs.inters, v)
}

func (ltrs *LiveTitleRegexSelect) sqlScan(ctx context.Context, root *LiveTitleRegexQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ltrs.fns))
	for _, fn := range ltrs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ltrs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ltrs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
