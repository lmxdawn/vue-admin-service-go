package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Permission holds the schema definition for the Permission entity.
type Permission struct {
	ent.Schema
}

// Fields of the Permission.
func (Permission) Fields() []ent.Field {
	return []ent.Field{
		field.Int("role_id"),
		field.Int("permission_rule_id"),
		field.String("type"),
	}
}

// Edges of the Permission.
func (Permission) Edges() []ent.Edge {
	return nil
}
