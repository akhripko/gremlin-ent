// Code generated by entc, DO NOT EDIT.

package persona

import (
	"github.com/akhripko/gremlin-ent/ent/predicate"
	"github.com/facebook/ent/dialect/gremlin/graph/dsl"
	"github.com/facebook/ent/dialect/gremlin/graph/dsl/__"
	"github.com/facebook/ent/dialect/gremlin/graph/dsl/p"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Persona {
	return predicate.Persona(func(t *dsl.Traversal) {
		t.HasID(id)
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Persona {
	return predicate.Persona(func(t *dsl.Traversal) {
		t.HasID(p.EQ(id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Persona {
	return predicate.Persona(func(t *dsl.Traversal) {
		t.HasID(p.NEQ(id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Persona {
	return predicate.Persona(func(t *dsl.Traversal) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		t.HasID(p.Within(v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Persona {
	return predicate.Persona(func(t *dsl.Traversal) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		t.HasID(p.Without(v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Persona {
	return predicate.Persona(func(t *dsl.Traversal) {
		t.HasID(p.GT(id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Persona {
	return predicate.Persona(func(t *dsl.Traversal) {
		t.HasID(p.GTE(id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Persona {
	return predicate.Persona(func(t *dsl.Traversal) {
		t.HasID(p.LT(id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Persona {
	return predicate.Persona(func(t *dsl.Traversal) {
		t.HasID(p.LTE(id))
	})
}

// Age applies equality check predicate on the "age" field. It's identical to AgeEQ.
func Age(v int) predicate.Persona {
	return predicate.Persona(func(t *dsl.Traversal) {
		t.Has(Label, FieldAge, p.EQ(v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Persona {
	return predicate.Persona(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.EQ(v))
	})
}

// Gender applies equality check predicate on the "gender" field. It's identical to GenderEQ.
func Gender(v string) predicate.Persona {
	return predicate.Persona(func(t *dsl.Traversal) {
		t.Has(Label, FieldGender, p.EQ(v))
	})
}

// AgeEQ applies the EQ predicate on the "age" field.
func AgeEQ(v int) predicate.Persona {
	return predicate.Persona(func(t *dsl.Traversal) {
		t.Has(Label, FieldAge, p.EQ(v))
	})
}

// AgeNEQ applies the NEQ predicate on the "age" field.
func AgeNEQ(v int) predicate.Persona {
	return predicate.Persona(func(t *dsl.Traversal) {
		t.Has(Label, FieldAge, p.NEQ(v))
	})
}

// AgeIn applies the In predicate on the "age" field.
func AgeIn(vs ...int) predicate.Persona {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Persona(func(t *dsl.Traversal) {
		t.Has(Label, FieldAge, p.Within(v...))
	})
}

// AgeNotIn applies the NotIn predicate on the "age" field.
func AgeNotIn(vs ...int) predicate.Persona {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Persona(func(t *dsl.Traversal) {
		t.Has(Label, FieldAge, p.Without(v...))
	})
}

// AgeGT applies the GT predicate on the "age" field.
func AgeGT(v int) predicate.Persona {
	return predicate.Persona(func(t *dsl.Traversal) {
		t.Has(Label, FieldAge, p.GT(v))
	})
}

// AgeGTE applies the GTE predicate on the "age" field.
func AgeGTE(v int) predicate.Persona {
	return predicate.Persona(func(t *dsl.Traversal) {
		t.Has(Label, FieldAge, p.GTE(v))
	})
}

// AgeLT applies the LT predicate on the "age" field.
func AgeLT(v int) predicate.Persona {
	return predicate.Persona(func(t *dsl.Traversal) {
		t.Has(Label, FieldAge, p.LT(v))
	})
}

// AgeLTE applies the LTE predicate on the "age" field.
func AgeLTE(v int) predicate.Persona {
	return predicate.Persona(func(t *dsl.Traversal) {
		t.Has(Label, FieldAge, p.LTE(v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Persona {
	return predicate.Persona(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.EQ(v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Persona {
	return predicate.Persona(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.NEQ(v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Persona {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Persona(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.Within(v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Persona {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Persona(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.Without(v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Persona {
	return predicate.Persona(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.GT(v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Persona {
	return predicate.Persona(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.GTE(v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Persona {
	return predicate.Persona(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.LT(v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Persona {
	return predicate.Persona(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.LTE(v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Persona {
	return predicate.Persona(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.Containing(v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Persona {
	return predicate.Persona(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.StartingWith(v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Persona {
	return predicate.Persona(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.EndingWith(v))
	})
}

// GenderEQ applies the EQ predicate on the "gender" field.
func GenderEQ(v string) predicate.Persona {
	return predicate.Persona(func(t *dsl.Traversal) {
		t.Has(Label, FieldGender, p.EQ(v))
	})
}

// GenderNEQ applies the NEQ predicate on the "gender" field.
func GenderNEQ(v string) predicate.Persona {
	return predicate.Persona(func(t *dsl.Traversal) {
		t.Has(Label, FieldGender, p.NEQ(v))
	})
}

// GenderIn applies the In predicate on the "gender" field.
func GenderIn(vs ...string) predicate.Persona {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Persona(func(t *dsl.Traversal) {
		t.Has(Label, FieldGender, p.Within(v...))
	})
}

// GenderNotIn applies the NotIn predicate on the "gender" field.
func GenderNotIn(vs ...string) predicate.Persona {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Persona(func(t *dsl.Traversal) {
		t.Has(Label, FieldGender, p.Without(v...))
	})
}

// GenderGT applies the GT predicate on the "gender" field.
func GenderGT(v string) predicate.Persona {
	return predicate.Persona(func(t *dsl.Traversal) {
		t.Has(Label, FieldGender, p.GT(v))
	})
}

// GenderGTE applies the GTE predicate on the "gender" field.
func GenderGTE(v string) predicate.Persona {
	return predicate.Persona(func(t *dsl.Traversal) {
		t.Has(Label, FieldGender, p.GTE(v))
	})
}

// GenderLT applies the LT predicate on the "gender" field.
func GenderLT(v string) predicate.Persona {
	return predicate.Persona(func(t *dsl.Traversal) {
		t.Has(Label, FieldGender, p.LT(v))
	})
}

// GenderLTE applies the LTE predicate on the "gender" field.
func GenderLTE(v string) predicate.Persona {
	return predicate.Persona(func(t *dsl.Traversal) {
		t.Has(Label, FieldGender, p.LTE(v))
	})
}

// GenderContains applies the Contains predicate on the "gender" field.
func GenderContains(v string) predicate.Persona {
	return predicate.Persona(func(t *dsl.Traversal) {
		t.Has(Label, FieldGender, p.Containing(v))
	})
}

// GenderHasPrefix applies the HasPrefix predicate on the "gender" field.
func GenderHasPrefix(v string) predicate.Persona {
	return predicate.Persona(func(t *dsl.Traversal) {
		t.Has(Label, FieldGender, p.StartingWith(v))
	})
}

// GenderHasSuffix applies the HasSuffix predicate on the "gender" field.
func GenderHasSuffix(v string) predicate.Persona {
	return predicate.Persona(func(t *dsl.Traversal) {
		t.Has(Label, FieldGender, p.EndingWith(v))
	})
}

// HasGenderIs applies the HasEdge predicate on the "genderIs" edge.
func HasGenderIs() predicate.Persona {
	return predicate.Persona(func(t *dsl.Traversal) {
		t.OutE(GenderIsLabel).OutV()
	})
}

// HasGenderIsWith applies the HasEdge predicate on the "genderIs" edge with a given conditions (other predicates).
func HasGenderIsWith(preds ...predicate.Gender) predicate.Persona {
	return predicate.Persona(func(t *dsl.Traversal) {
		tr := __.InV()
		for _, p := range preds {
			p(tr)
		}
		t.OutE(GenderIsLabel).Where(tr).OutV()
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Persona) predicate.Persona {
	return predicate.Persona(func(tr *dsl.Traversal) {
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
func Or(predicates ...predicate.Persona) predicate.Persona {
	return predicate.Persona(func(tr *dsl.Traversal) {
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
func Not(p predicate.Persona) predicate.Persona {
	return predicate.Persona(func(tr *dsl.Traversal) {
		t := __.New()
		p(t)
		tr.Where(__.Not(t))
	})
}
