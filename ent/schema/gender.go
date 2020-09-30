package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Gender holds the schema definition for the Gender entity.
type Gender struct {
	ent.Schema
}

// Fields of the Gender.
func (Gender) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Positive(),
		field.String("value"),
	}
}

// Edges of the Gender.
func (Gender) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("personas", Persona.Type).
			Ref("genderIs"),
	}
}
