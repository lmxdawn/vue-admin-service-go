// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"vue-admin/app/rbac/service/internal/data/ent/permissionrule"
	"vue-admin/app/rbac/service/internal/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PermissionRuleQuery is the builder for querying PermissionRule entities.
type PermissionRuleQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.PermissionRule
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PermissionRuleQuery builder.
func (prq *PermissionRuleQuery) Where(ps ...predicate.PermissionRule) *PermissionRuleQuery {
	prq.predicates = append(prq.predicates, ps...)
	return prq
}

// Limit adds a limit step to the query.
func (prq *PermissionRuleQuery) Limit(limit int) *PermissionRuleQuery {
	prq.limit = &limit
	return prq
}

// Offset adds an offset step to the query.
func (prq *PermissionRuleQuery) Offset(offset int) *PermissionRuleQuery {
	prq.offset = &offset
	return prq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (prq *PermissionRuleQuery) Unique(unique bool) *PermissionRuleQuery {
	prq.unique = &unique
	return prq
}

// Order adds an order step to the query.
func (prq *PermissionRuleQuery) Order(o ...OrderFunc) *PermissionRuleQuery {
	prq.order = append(prq.order, o...)
	return prq
}

// First returns the first PermissionRule entity from the query.
// Returns a *NotFoundError when no PermissionRule was found.
func (prq *PermissionRuleQuery) First(ctx context.Context) (*PermissionRule, error) {
	nodes, err := prq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{permissionrule.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (prq *PermissionRuleQuery) FirstX(ctx context.Context) *PermissionRule {
	node, err := prq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PermissionRule ID from the query.
// Returns a *NotFoundError when no PermissionRule ID was found.
func (prq *PermissionRuleQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = prq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{permissionrule.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (prq *PermissionRuleQuery) FirstIDX(ctx context.Context) int {
	id, err := prq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PermissionRule entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PermissionRule entity is found.
// Returns a *NotFoundError when no PermissionRule entities are found.
func (prq *PermissionRuleQuery) Only(ctx context.Context) (*PermissionRule, error) {
	nodes, err := prq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{permissionrule.Label}
	default:
		return nil, &NotSingularError{permissionrule.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (prq *PermissionRuleQuery) OnlyX(ctx context.Context) *PermissionRule {
	node, err := prq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PermissionRule ID in the query.
// Returns a *NotSingularError when more than one PermissionRule ID is found.
// Returns a *NotFoundError when no entities are found.
func (prq *PermissionRuleQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = prq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{permissionrule.Label}
	default:
		err = &NotSingularError{permissionrule.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (prq *PermissionRuleQuery) OnlyIDX(ctx context.Context) int {
	id, err := prq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PermissionRules.
func (prq *PermissionRuleQuery) All(ctx context.Context) ([]*PermissionRule, error) {
	if err := prq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return prq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (prq *PermissionRuleQuery) AllX(ctx context.Context) []*PermissionRule {
	nodes, err := prq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PermissionRule IDs.
func (prq *PermissionRuleQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := prq.Select(permissionrule.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (prq *PermissionRuleQuery) IDsX(ctx context.Context) []int {
	ids, err := prq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (prq *PermissionRuleQuery) Count(ctx context.Context) (int, error) {
	if err := prq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return prq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (prq *PermissionRuleQuery) CountX(ctx context.Context) int {
	count, err := prq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (prq *PermissionRuleQuery) Exist(ctx context.Context) (bool, error) {
	if err := prq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return prq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (prq *PermissionRuleQuery) ExistX(ctx context.Context) bool {
	exist, err := prq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PermissionRuleQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (prq *PermissionRuleQuery) Clone() *PermissionRuleQuery {
	if prq == nil {
		return nil
	}
	return &PermissionRuleQuery{
		config:     prq.config,
		limit:      prq.limit,
		offset:     prq.offset,
		order:      append([]OrderFunc{}, prq.order...),
		predicates: append([]predicate.PermissionRule{}, prq.predicates...),
		// clone intermediate query.
		sql:    prq.sql.Clone(),
		path:   prq.path,
		unique: prq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Pid int `json:"pid,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PermissionRule.Query().
//		GroupBy(permissionrule.FieldPid).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (prq *PermissionRuleQuery) GroupBy(field string, fields ...string) *PermissionRuleGroupBy {
	grbuild := &PermissionRuleGroupBy{config: prq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := prq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return prq.sqlQuery(ctx), nil
	}
	grbuild.label = permissionrule.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Pid int `json:"pid,omitempty"`
//	}
//
//	client.PermissionRule.Query().
//		Select(permissionrule.FieldPid).
//		Scan(ctx, &v)
//
func (prq *PermissionRuleQuery) Select(fields ...string) *PermissionRuleSelect {
	prq.fields = append(prq.fields, fields...)
	selbuild := &PermissionRuleSelect{PermissionRuleQuery: prq}
	selbuild.label = permissionrule.Label
	selbuild.flds, selbuild.scan = &prq.fields, selbuild.Scan
	return selbuild
}

func (prq *PermissionRuleQuery) prepareQuery(ctx context.Context) error {
	for _, f := range prq.fields {
		if !permissionrule.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if prq.path != nil {
		prev, err := prq.path(ctx)
		if err != nil {
			return err
		}
		prq.sql = prev
	}
	return nil
}

func (prq *PermissionRuleQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PermissionRule, error) {
	var (
		nodes = []*PermissionRule{}
		_spec = prq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*PermissionRule).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &PermissionRule{config: prq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, prq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (prq *PermissionRuleQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := prq.querySpec()
	_spec.Node.Columns = prq.fields
	if len(prq.fields) > 0 {
		_spec.Unique = prq.unique != nil && *prq.unique
	}
	return sqlgraph.CountNodes(ctx, prq.driver, _spec)
}

func (prq *PermissionRuleQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := prq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (prq *PermissionRuleQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   permissionrule.Table,
			Columns: permissionrule.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: permissionrule.FieldID,
			},
		},
		From:   prq.sql,
		Unique: true,
	}
	if unique := prq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := prq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, permissionrule.FieldID)
		for i := range fields {
			if fields[i] != permissionrule.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := prq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := prq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := prq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := prq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (prq *PermissionRuleQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(prq.driver.Dialect())
	t1 := builder.Table(permissionrule.Table)
	columns := prq.fields
	if len(columns) == 0 {
		columns = permissionrule.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if prq.sql != nil {
		selector = prq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if prq.unique != nil && *prq.unique {
		selector.Distinct()
	}
	for _, p := range prq.predicates {
		p(selector)
	}
	for _, p := range prq.order {
		p(selector)
	}
	if offset := prq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := prq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PermissionRuleGroupBy is the group-by builder for PermissionRule entities.
type PermissionRuleGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (prgb *PermissionRuleGroupBy) Aggregate(fns ...AggregateFunc) *PermissionRuleGroupBy {
	prgb.fns = append(prgb.fns, fns...)
	return prgb
}

// Scan applies the group-by query and scans the result into the given value.
func (prgb *PermissionRuleGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := prgb.path(ctx)
	if err != nil {
		return err
	}
	prgb.sql = query
	return prgb.sqlScan(ctx, v)
}

func (prgb *PermissionRuleGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range prgb.fields {
		if !permissionrule.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := prgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := prgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (prgb *PermissionRuleGroupBy) sqlQuery() *sql.Selector {
	selector := prgb.sql.Select()
	aggregation := make([]string, 0, len(prgb.fns))
	for _, fn := range prgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(prgb.fields)+len(prgb.fns))
		for _, f := range prgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(prgb.fields...)...)
}

// PermissionRuleSelect is the builder for selecting fields of PermissionRule entities.
type PermissionRuleSelect struct {
	*PermissionRuleQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (prs *PermissionRuleSelect) Scan(ctx context.Context, v interface{}) error {
	if err := prs.prepareQuery(ctx); err != nil {
		return err
	}
	prs.sql = prs.PermissionRuleQuery.sqlQuery(ctx)
	return prs.sqlScan(ctx, v)
}

func (prs *PermissionRuleSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := prs.sql.Query()
	if err := prs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
