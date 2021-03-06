// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/akhripko/gremlin-ent/ent/gender"
	"github.com/akhripko/gremlin-ent/ent/persona"
	"github.com/akhripko/gremlin-ent/ent/predicate"
	"github.com/facebook/ent/dialect/gremlin"
	"github.com/facebook/ent/dialect/gremlin/graph/dsl"
	"github.com/facebook/ent/dialect/gremlin/graph/dsl/__"
	"github.com/facebook/ent/dialect/gremlin/graph/dsl/g"
)

// GenderUpdate is the builder for updating Gender entities.
type GenderUpdate struct {
	config
	hooks      []Hook
	mutation   *GenderMutation
	predicates []predicate.Gender
}

// Where adds a new predicate for the builder.
func (gu *GenderUpdate) Where(ps ...predicate.Gender) *GenderUpdate {
	gu.predicates = append(gu.predicates, ps...)
	return gu
}

// SetValue sets the value field.
func (gu *GenderUpdate) SetValue(s string) *GenderUpdate {
	gu.mutation.SetValue(s)
	return gu
}

// AddPersonaIDs adds the personas edge to Persona by ids.
func (gu *GenderUpdate) AddPersonaIDs(ids ...int) *GenderUpdate {
	gu.mutation.AddPersonaIDs(ids...)
	return gu
}

// AddPersonas adds the personas edges to Persona.
func (gu *GenderUpdate) AddPersonas(p ...*Persona) *GenderUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return gu.AddPersonaIDs(ids...)
}

// Mutation returns the GenderMutation object of the builder.
func (gu *GenderUpdate) Mutation() *GenderMutation {
	return gu.mutation
}

// ClearPersonas clears all "personas" edges to type Persona.
func (gu *GenderUpdate) ClearPersonas() *GenderUpdate {
	gu.mutation.ClearPersonas()
	return gu
}

// RemovePersonaIDs removes the personas edge to Persona by ids.
func (gu *GenderUpdate) RemovePersonaIDs(ids ...int) *GenderUpdate {
	gu.mutation.RemovePersonaIDs(ids...)
	return gu
}

// RemovePersonas removes personas edges to Persona.
func (gu *GenderUpdate) RemovePersonas(p ...*Persona) *GenderUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return gu.RemovePersonaIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (gu *GenderUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(gu.hooks) == 0 {
		affected, err = gu.gremlinSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GenderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gu.mutation = mutation
			affected, err = gu.gremlinSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gu.hooks) - 1; i >= 0; i-- {
			mut = gu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GenderUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GenderUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GenderUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (gu *GenderUpdate) gremlinSave(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := gu.gremlin().Query()
	if err := gu.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	if err, ok := isConstantError(res); ok {
		return 0, err
	}
	return res.ReadInt()
}

func (gu *GenderUpdate) gremlin() *dsl.Traversal {
	v := g.V().HasLabel(gender.Label)
	for _, p := range gu.predicates {
		p(v)
	}
	var (
		rv = v.Clone()
		_  = rv

		trs []*dsl.Traversal
	)
	if value, ok := gu.mutation.Value(); ok {
		v.Property(dsl.Single, gender.FieldValue, value)
	}
	for _, id := range gu.mutation.RemovedPersonasIDs() {
		tr := rv.Clone().InE(persona.GenderIsLabel).Where(__.OtherV().HasID(id)).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range gu.mutation.PersonasIDs() {
		v.AddE(persona.GenderIsLabel).From(g.V(id)).InV()
	}
	v.Count()
	trs = append(trs, v)
	return dsl.Join(trs...)
}

// GenderUpdateOne is the builder for updating a single Gender entity.
type GenderUpdateOne struct {
	config
	hooks    []Hook
	mutation *GenderMutation
}

// SetValue sets the value field.
func (guo *GenderUpdateOne) SetValue(s string) *GenderUpdateOne {
	guo.mutation.SetValue(s)
	return guo
}

// AddPersonaIDs adds the personas edge to Persona by ids.
func (guo *GenderUpdateOne) AddPersonaIDs(ids ...int) *GenderUpdateOne {
	guo.mutation.AddPersonaIDs(ids...)
	return guo
}

// AddPersonas adds the personas edges to Persona.
func (guo *GenderUpdateOne) AddPersonas(p ...*Persona) *GenderUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return guo.AddPersonaIDs(ids...)
}

// Mutation returns the GenderMutation object of the builder.
func (guo *GenderUpdateOne) Mutation() *GenderMutation {
	return guo.mutation
}

// ClearPersonas clears all "personas" edges to type Persona.
func (guo *GenderUpdateOne) ClearPersonas() *GenderUpdateOne {
	guo.mutation.ClearPersonas()
	return guo
}

// RemovePersonaIDs removes the personas edge to Persona by ids.
func (guo *GenderUpdateOne) RemovePersonaIDs(ids ...int) *GenderUpdateOne {
	guo.mutation.RemovePersonaIDs(ids...)
	return guo
}

// RemovePersonas removes personas edges to Persona.
func (guo *GenderUpdateOne) RemovePersonas(p ...*Persona) *GenderUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return guo.RemovePersonaIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (guo *GenderUpdateOne) Save(ctx context.Context) (*Gender, error) {
	var (
		err  error
		node *Gender
	)
	if len(guo.hooks) == 0 {
		node, err = guo.gremlinSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GenderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			guo.mutation = mutation
			node, err = guo.gremlinSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(guo.hooks) - 1; i >= 0; i-- {
			mut = guo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, guo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (guo *GenderUpdateOne) SaveX(ctx context.Context) *Gender {
	node, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (guo *GenderUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GenderUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (guo *GenderUpdateOne) gremlinSave(ctx context.Context) (*Gender, error) {
	res := &gremlin.Response{}
	id, ok := guo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Gender.ID for update")}
	}
	query, bindings := guo.gremlin(id).Query()
	if err := guo.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	ge := &Gender{config: guo.config}
	if err := ge.FromResponse(res); err != nil {
		return nil, err
	}
	return ge, nil
}

func (guo *GenderUpdateOne) gremlin(id int) *dsl.Traversal {
	v := g.V(id)
	var (
		rv = v.Clone()
		_  = rv

		trs []*dsl.Traversal
	)
	if value, ok := guo.mutation.Value(); ok {
		v.Property(dsl.Single, gender.FieldValue, value)
	}
	for _, id := range guo.mutation.RemovedPersonasIDs() {
		tr := rv.Clone().InE(persona.GenderIsLabel).Where(__.OtherV().HasID(id)).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range guo.mutation.PersonasIDs() {
		v.AddE(persona.GenderIsLabel).From(g.V(id)).InV()
	}
	v.ValueMap(true)
	trs = append(trs, v)
	return dsl.Join(trs...)
}
