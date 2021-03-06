// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/akhripko/gremlin-ent/ent/gender"
	"github.com/akhripko/gremlin-ent/ent/persona"
	"github.com/akhripko/gremlin-ent/ent/predicate"
	"github.com/facebook/ent/dialect/gremlin"
	"github.com/facebook/ent/dialect/gremlin/graph/dsl"
	"github.com/facebook/ent/dialect/gremlin/graph/dsl/__"
	"github.com/facebook/ent/dialect/gremlin/graph/dsl/g"
)

// GenderQuery is the builder for querying Gender entities.
type GenderQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.Gender
	// eager-loading edges.
	withPersonas *PersonaQuery
	// intermediate query (i.e. traversal path).
	gremlin *dsl.Traversal
	path    func(context.Context) (*dsl.Traversal, error)
}

// Where adds a new predicate for the builder.
func (gq *GenderQuery) Where(ps ...predicate.Gender) *GenderQuery {
	gq.predicates = append(gq.predicates, ps...)
	return gq
}

// Limit adds a limit step to the query.
func (gq *GenderQuery) Limit(limit int) *GenderQuery {
	gq.limit = &limit
	return gq
}

// Offset adds an offset step to the query.
func (gq *GenderQuery) Offset(offset int) *GenderQuery {
	gq.offset = &offset
	return gq
}

// Order adds an order step to the query.
func (gq *GenderQuery) Order(o ...OrderFunc) *GenderQuery {
	gq.order = append(gq.order, o...)
	return gq
}

// QueryPersonas chains the current query on the personas edge.
func (gq *GenderQuery) QueryPersonas() *PersonaQuery {
	query := &PersonaQuery{config: gq.config}
	query.path = func(ctx context.Context) (fromU *dsl.Traversal, err error) {
		if err := gq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		gremlin := gq.gremlinQuery()
		fromU = gremlin.InE(persona.GenderIsLabel).OutV()
		return fromU, nil
	}
	return query
}

// First returns the first Gender entity in the query. Returns *NotFoundError when no gender was found.
func (gq *GenderQuery) First(ctx context.Context) (*Gender, error) {
	nodes, err := gq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{gender.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gq *GenderQuery) FirstX(ctx context.Context) *Gender {
	node, err := gq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Gender id in the query. Returns *NotFoundError when no id was found.
func (gq *GenderQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{gender.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (gq *GenderQuery) FirstXID(ctx context.Context) int {
	id, err := gq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Gender entity in the query, returns an error if not exactly one entity was returned.
func (gq *GenderQuery) Only(ctx context.Context) (*Gender, error) {
	nodes, err := gq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{gender.Label}
	default:
		return nil, &NotSingularError{gender.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gq *GenderQuery) OnlyX(ctx context.Context) *Gender {
	node, err := gq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID returns the only Gender id in the query, returns an error if not exactly one id was returned.
func (gq *GenderQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{gender.Label}
	default:
		err = &NotSingularError{gender.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gq *GenderQuery) OnlyIDX(ctx context.Context) int {
	id, err := gq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Genders.
func (gq *GenderQuery) All(ctx context.Context) ([]*Gender, error) {
	if err := gq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return gq.gremlinAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (gq *GenderQuery) AllX(ctx context.Context) []*Gender {
	nodes, err := gq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Gender ids.
func (gq *GenderQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := gq.Select(gender.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gq *GenderQuery) IDsX(ctx context.Context) []int {
	ids, err := gq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gq *GenderQuery) Count(ctx context.Context) (int, error) {
	if err := gq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return gq.gremlinCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (gq *GenderQuery) CountX(ctx context.Context) int {
	count, err := gq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gq *GenderQuery) Exist(ctx context.Context) (bool, error) {
	if err := gq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return gq.gremlinExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (gq *GenderQuery) ExistX(ctx context.Context) bool {
	exist, err := gq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gq *GenderQuery) Clone() *GenderQuery {
	return &GenderQuery{
		config:     gq.config,
		limit:      gq.limit,
		offset:     gq.offset,
		order:      append([]OrderFunc{}, gq.order...),
		unique:     append([]string{}, gq.unique...),
		predicates: append([]predicate.Gender{}, gq.predicates...),
		// clone intermediate query.
		gremlin: gq.gremlin.Clone(),
		path:    gq.path,
	}
}

//  WithPersonas tells the query-builder to eager-loads the nodes that are connected to
// the "personas" edge. The optional arguments used to configure the query builder of the edge.
func (gq *GenderQuery) WithPersonas(opts ...func(*PersonaQuery)) *GenderQuery {
	query := &PersonaQuery{config: gq.config}
	for _, opt := range opts {
		opt(query)
	}
	gq.withPersonas = query
	return gq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Value string `json:"value,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Gender.Query().
//		GroupBy(gender.FieldValue).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (gq *GenderQuery) GroupBy(field string, fields ...string) *GenderGroupBy {
	group := &GenderGroupBy{config: gq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *dsl.Traversal, err error) {
		if err := gq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return gq.gremlinQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Value string `json:"value,omitempty"`
//	}
//
//	client.Gender.Query().
//		Select(gender.FieldValue).
//		Scan(ctx, &v)
//
func (gq *GenderQuery) Select(field string, fields ...string) *GenderSelect {
	selector := &GenderSelect{config: gq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *dsl.Traversal, err error) {
		if err := gq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return gq.gremlinQuery(), nil
	}
	return selector
}

func (gq *GenderQuery) prepareQuery(ctx context.Context) error {
	if gq.path != nil {
		prev, err := gq.path(ctx)
		if err != nil {
			return err
		}
		gq.gremlin = prev
	}
	return nil
}

func (gq *GenderQuery) gremlinAll(ctx context.Context) ([]*Gender, error) {
	res := &gremlin.Response{}
	query, bindings := gq.gremlinQuery().ValueMap(true).Query()
	if err := gq.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	var ges Genders
	if err := ges.FromResponse(res); err != nil {
		return nil, err
	}
	ges.config(gq.config)
	return ges, nil
}

func (gq *GenderQuery) gremlinCount(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := gq.gremlinQuery().Count().Query()
	if err := gq.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	return res.ReadInt()
}

func (gq *GenderQuery) gremlinExist(ctx context.Context) (bool, error) {
	res := &gremlin.Response{}
	query, bindings := gq.gremlinQuery().HasNext().Query()
	if err := gq.driver.Exec(ctx, query, bindings, res); err != nil {
		return false, err
	}
	return res.ReadBool()
}

func (gq *GenderQuery) gremlinQuery() *dsl.Traversal {
	v := g.V().HasLabel(gender.Label)
	if gq.gremlin != nil {
		v = gq.gremlin.Clone()
	}
	for _, p := range gq.predicates {
		p(v)
	}
	if len(gq.order) > 0 {
		v.Order()
		for _, p := range gq.order {
			p(v)
		}
	}
	switch limit, offset := gq.limit, gq.offset; {
	case limit != nil && offset != nil:
		v.Range(*offset, *offset+*limit)
	case offset != nil:
		v.Range(*offset, math.MaxInt32)
	case limit != nil:
		v.Limit(*limit)
	}
	if unique := gq.unique; len(unique) == 0 {
		v.Dedup()
	}
	return v
}

// GenderGroupBy is the builder for group-by Gender entities.
type GenderGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	gremlin *dsl.Traversal
	path    func(context.Context) (*dsl.Traversal, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ggb *GenderGroupBy) Aggregate(fns ...AggregateFunc) *GenderGroupBy {
	ggb.fns = append(ggb.fns, fns...)
	return ggb
}

// Scan applies the group-by query and scan the result into the given value.
func (ggb *GenderGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ggb.path(ctx)
	if err != nil {
		return err
	}
	ggb.gremlin = query
	return ggb.gremlinScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ggb *GenderGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ggb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (ggb *GenderGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ggb.fields) > 1 {
		return nil, errors.New("ent: GenderGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ggb *GenderGroupBy) StringsX(ctx context.Context) []string {
	v, err := ggb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (ggb *GenderGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ggb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{gender.Label}
	default:
		err = fmt.Errorf("ent: GenderGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ggb *GenderGroupBy) StringX(ctx context.Context) string {
	v, err := ggb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (ggb *GenderGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ggb.fields) > 1 {
		return nil, errors.New("ent: GenderGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ggb *GenderGroupBy) IntsX(ctx context.Context) []int {
	v, err := ggb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (ggb *GenderGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ggb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{gender.Label}
	default:
		err = fmt.Errorf("ent: GenderGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ggb *GenderGroupBy) IntX(ctx context.Context) int {
	v, err := ggb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (ggb *GenderGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ggb.fields) > 1 {
		return nil, errors.New("ent: GenderGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ggb *GenderGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ggb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (ggb *GenderGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ggb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{gender.Label}
	default:
		err = fmt.Errorf("ent: GenderGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ggb *GenderGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ggb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (ggb *GenderGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ggb.fields) > 1 {
		return nil, errors.New("ent: GenderGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ggb *GenderGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ggb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (ggb *GenderGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ggb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{gender.Label}
	default:
		err = fmt.Errorf("ent: GenderGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ggb *GenderGroupBy) BoolX(ctx context.Context) bool {
	v, err := ggb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ggb *GenderGroupBy) gremlinScan(ctx context.Context, v interface{}) error {
	res := &gremlin.Response{}
	query, bindings := ggb.gremlinQuery().Query()
	if err := ggb.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(ggb.fields)+len(ggb.fns) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}

func (ggb *GenderGroupBy) gremlinQuery() *dsl.Traversal {
	var (
		trs   []interface{}
		names []interface{}
	)
	for _, fn := range ggb.fns {
		name, tr := fn("p", "")
		trs = append(trs, tr)
		names = append(names, name)
	}
	for _, f := range ggb.fields {
		names = append(names, f)
		trs = append(trs, __.As("p").Unfold().Values(f).As(f))
	}
	return ggb.gremlin.Group().
		By(__.Values(ggb.fields...).Fold()).
		By(__.Fold().Match(trs...).Select(names...)).
		Select(dsl.Values).
		Next()
}

// GenderSelect is the builder for select fields of Gender entities.
type GenderSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	gremlin *dsl.Traversal
	path    func(context.Context) (*dsl.Traversal, error)
}

// Scan applies the selector query and scan the result into the given value.
func (gs *GenderSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := gs.path(ctx)
	if err != nil {
		return err
	}
	gs.gremlin = query
	return gs.gremlinScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (gs *GenderSelect) ScanX(ctx context.Context, v interface{}) {
	if err := gs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (gs *GenderSelect) Strings(ctx context.Context) ([]string, error) {
	if len(gs.fields) > 1 {
		return nil, errors.New("ent: GenderSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := gs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (gs *GenderSelect) StringsX(ctx context.Context) []string {
	v, err := gs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (gs *GenderSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = gs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{gender.Label}
	default:
		err = fmt.Errorf("ent: GenderSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (gs *GenderSelect) StringX(ctx context.Context) string {
	v, err := gs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (gs *GenderSelect) Ints(ctx context.Context) ([]int, error) {
	if len(gs.fields) > 1 {
		return nil, errors.New("ent: GenderSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := gs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (gs *GenderSelect) IntsX(ctx context.Context) []int {
	v, err := gs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (gs *GenderSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = gs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{gender.Label}
	default:
		err = fmt.Errorf("ent: GenderSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (gs *GenderSelect) IntX(ctx context.Context) int {
	v, err := gs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (gs *GenderSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(gs.fields) > 1 {
		return nil, errors.New("ent: GenderSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := gs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (gs *GenderSelect) Float64sX(ctx context.Context) []float64 {
	v, err := gs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (gs *GenderSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = gs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{gender.Label}
	default:
		err = fmt.Errorf("ent: GenderSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (gs *GenderSelect) Float64X(ctx context.Context) float64 {
	v, err := gs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (gs *GenderSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(gs.fields) > 1 {
		return nil, errors.New("ent: GenderSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := gs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (gs *GenderSelect) BoolsX(ctx context.Context) []bool {
	v, err := gs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (gs *GenderSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = gs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{gender.Label}
	default:
		err = fmt.Errorf("ent: GenderSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (gs *GenderSelect) BoolX(ctx context.Context) bool {
	v, err := gs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gs *GenderSelect) gremlinScan(ctx context.Context, v interface{}) error {
	var (
		traversal *dsl.Traversal
		res       = &gremlin.Response{}
	)
	if len(gs.fields) == 1 {
		if gs.fields[0] != gender.FieldID {
			traversal = gs.gremlin.Values(gs.fields...)
		} else {
			traversal = gs.gremlin.ID()
		}
	} else {
		fields := make([]interface{}, len(gs.fields))
		for i, f := range gs.fields {
			fields[i] = f
		}
		traversal = gs.gremlin.ValueMap(fields...)
	}
	query, bindings := traversal.Query()
	if err := gs.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(gs.fields) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}
