// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/akhripko/gremlin-ent/ent/persona"
	"github.com/akhripko/gremlin-ent/ent/predicate"
	"github.com/facebook/ent/dialect/gremlin"
	"github.com/facebook/ent/dialect/gremlin/graph/dsl"
	"github.com/facebook/ent/dialect/gremlin/graph/dsl/__"
	"github.com/facebook/ent/dialect/gremlin/graph/dsl/g"
)

// PersonaDelete is the builder for deleting a Persona entity.
type PersonaDelete struct {
	config
	hooks      []Hook
	mutation   *PersonaMutation
	predicates []predicate.Persona
}

// Where adds a new predicate to the delete builder.
func (pd *PersonaDelete) Where(ps ...predicate.Persona) *PersonaDelete {
	pd.predicates = append(pd.predicates, ps...)
	return pd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pd *PersonaDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pd.hooks) == 0 {
		affected, err = pd.gremlinExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PersonaMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pd.mutation = mutation
			affected, err = pd.gremlinExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pd.hooks) - 1; i >= 0; i-- {
			mut = pd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (pd *PersonaDelete) ExecX(ctx context.Context) int {
	n, err := pd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pd *PersonaDelete) gremlinExec(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := pd.gremlin().Query()
	if err := pd.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	return res.ReadInt()
}

func (pd *PersonaDelete) gremlin() *dsl.Traversal {
	t := g.V().HasLabel(persona.Label)
	for _, p := range pd.predicates {
		p(t)
	}
	return t.SideEffect(__.Drop()).Count()
}

// PersonaDeleteOne is the builder for deleting a single Persona entity.
type PersonaDeleteOne struct {
	pd *PersonaDelete
}

// Exec executes the deletion query.
func (pdo *PersonaDeleteOne) Exec(ctx context.Context) error {
	n, err := pdo.pd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{persona.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pdo *PersonaDeleteOne) ExecX(ctx context.Context) {
	pdo.pd.ExecX(ctx)
}
