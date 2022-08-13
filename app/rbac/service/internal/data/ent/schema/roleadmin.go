package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// RoleAdmin holds the schema definition for the RoleAdmin entity.
type RoleAdmin struct {
	ent.Schema
}

// Annotations 用户实体的注解
func (RoleAdmin) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "role_admin"},
	}
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
