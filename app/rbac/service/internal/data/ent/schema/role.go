package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Role holds the schema definition for the Role entity.
type Role struct {
	ent.Schema
}

// Fields of the Role.
func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("name"),
		field.Int("pid"),
		field.Int("status"),
		field.String("remark"),
		field.Int("sort"),
		field.Time("create_time"),
		field.Time("update_time"),
	}
}

// Edges of the Role.
func (Role) Edges() []ent.Edge {
	return nil
}
