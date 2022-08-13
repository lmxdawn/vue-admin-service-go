package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// RoleAdmin holds the schema definition for the RoleAdmin entity.
type RoleAdmin struct {
	ent.Schema
}

// Fields of the RoleAdmin.
func (RoleAdmin) Fields() []ent.Field {
	return []ent.Field{
		field.Int("role_id"),
		field.Int("admin_id"),
	}
}

// Edges of the RoleAdmin.
func (RoleAdmin) Edges() []ent.Edge {
	return nil
}
