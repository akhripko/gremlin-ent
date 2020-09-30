// Code generated by entc, DO NOT EDIT.

package gender

import (
	"github.com/akhripko/gremlin-ent/ent/predicate"
	"github.com/facebook/ent/dialect/gremlin/graph/dsl"
	"github.com/facebook/ent/dialect/gremlin/graph/dsl/__"
	"github.com/facebook/ent/dialect/gremlin/graph/dsl/p"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Gender {
	return predicate.Gender(func(t *dsl.Traversal) {
		t.HasID(id)
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Gender {
	return predicate.Gender(func(t *dsl.Traversal) {
		t.HasID(p.EQ(id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Gender {
	return predicate.Gender(func(t *dsl.Traversal) {
		t.HasID(p.NEQ(id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Gender {
	return predicate.Gender(func(t *dsl.Traversal) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		t.HasID(p.Within(v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Gender {
	return predicate.Gender(func(t *dsl.Traversal) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		t.HasID(p.Without(v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Gender {
	return predicate.Gender(func(t *dsl.Traversal) {
		t.HasID(p.GT(id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Gender {
	return predicate.Gender(func(t *dsl.Traversal) {
		t.HasID(p.GTE(id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Gender {
	return predicate.Gender(func(t *dsl.Traversal) {
		t.HasID(p.LT(id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Gender {
	return predicate.Gender(func(t *dsl.Traversal) {
		t.HasID(p.LTE(id))
	})
}

// Value applies equality check predicate on the "value" field. It's identical to ValueEQ.
func Value(v string) predicate.Gender {
	return predicate.Gender(func(t *dsl.Traversal) {
		t.Has(Label, FieldValue, p.EQ(v))
	})
}

// ValueEQ applies the EQ predicate on the "value" field.
func ValueEQ(v string) predicate.Gender {
	return predicate.Gender(func(t *dsl.Traversal) {
		t.Has(Label, FieldValue, p.EQ(v))
	})
}

// ValueNEQ applies the NEQ predicate on the "value" field.
func ValueNEQ(v string) predicate.Gender {
	return predicate.Gender(func(t *dsl.Traversal) {
		t.Has(Label, FieldValue, p.NEQ(v))
	})
}

// ValueIn applies the In predicate on the "value" field.
func ValueIn(vs ...string) predicate.Gender {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Gender(func(t *dsl.Traversal) {
		t.Has(Label, FieldValue, p.Within(v...))
	})
}

// ValueNotIn applies the NotIn predicate on the "value" field.
func ValueNotIn(vs ...string) predicate.Gender {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Gender(func(t *dsl.Traversal) {
		t.Has(Label, FieldValue, p.Without(v...))
	})
}

// ValueGT applies the GT predicate on the "value" field.
func ValueGT(v string) predicate.Gender {
	return predicate.Gender(func(t *dsl.Traversal) {
		t.Has(Label, FieldValue, p.GT(v))
	})
}

// ValueGTE applies the GTE predicate on the "value" field.
func ValueGTE(v string) predicate.Gender {
	return predicate.Gender(func(t *dsl.Traversal) {
		t.Has(Label, FieldValue, p.GTE(v))
	})
}

// ValueLT applies the LT predicate on the "value" field.
func ValueLT(v string) predicate.Gender {
	return predicate.Gender(func(t *dsl.Traversal) {
		t.Has(Label, FieldValue, p.LT(v))
	})
}

// ValueLTE applies the LTE predicate on the "value" field.
func ValueLTE(v string) predicate.Gender {
	return predicate.Gender(func(t *dsl.Traversal) {
		t.Has(Label, FieldValue, p.LTE(v))
	})
}

// ValueContains applies the Contains predicate on the "value" field.
func ValueContains(v string) predicate.Gender {
	return predicate.Gender(func(t *dsl.Traversal) {
		t.Has(Label, FieldValue, p.Containing(v))
	})
}

// ValueHasPrefix applies the HasPrefix predicate on the "value" field.
func ValueHasPrefix(v string) predicate.Gender {
	return predicate.Gender(func(t *dsl.Traversal) {
		t.Has(Label, FieldValue, p.StartingWith(v))
	})
}

// ValueHasSuffix applies the HasSuffix predicate on the "value" field.
func ValueHasSuffix(v string) predicate.Gender {
	return predicate.Gender(func(t *dsl.Traversal) {
		t.Has(Label, FieldValue, p.EndingWith(v))
	})
}

// HasPersonas applies the HasEdge predicate on the "personas" edge.
func HasPersonas() predicate.Gender {
	return predicate.Gender(func(t *dsl.Traversal) {
		t.InE(PersonasInverseLabel).InV()
	})
}

// HasPersonasWith applies the HasEdge predicate on the "personas" edge with a given conditions (other predicates).
func HasPersonasWith(preds ...predicate.Persona) predicate.Gender {
	return predicate.Gender(func(t *dsl.Traversal) {
		tr := __.OutV()
		for _, p := range preds {
			p(tr)
		}
		t.InE(PersonasInverseLabel).Where(tr).InV()
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Gender) predicate.Gender {
	return predicate.Gender(func(tr *dsl.Traversal) {
		trs := make([]interface{}, 0, len(predicates))
		for _, p := range predicates {
			t := __.New()
			p(t)
			trs = append(trs, t)
		}
		tr.Where(__.And(trs...))
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Gender) predicate.Gender {
	return predicate.Gender(func(tr *dsl.Traversal) {
		trs := make([]interface{}, 0, len(predicates))
		for _, p := range predicates {
			t := __.New()
			p(t)
			trs = append(trs, t)
		}
		tr.Where(__.Or(trs...))
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Gender) predicate.Gender {
	return predicate.Gender(func(tr *dsl.Traversal) {
		t := __.New()
		p(t)
		tr.Where(__.Not(t))
	})
}
