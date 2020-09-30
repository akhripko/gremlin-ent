package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Persona holds the schema definition for the Persona entity.
type Persona struct {
	ent.Schema
}

// Fields of the Persona.
func (Persona) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Positive(),
		field.Int("age").
			Positive(),
		field.String("name").
			Default("unknown"),
		field.String("gender").
			Default("unknown"),
	}
}

// Edges of the Persona.
func (Persona) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("genderIs", Gender.Type),
	}
}
