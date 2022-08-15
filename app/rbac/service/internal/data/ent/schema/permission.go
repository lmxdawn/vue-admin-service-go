package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Permission holds the schema definition for the Permission entity.
type Permission struct {
	ent.Schema
}

// Annotations 用户实体的注解
func (Permission) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "permission"},
	}
}

// Fields of the Permission.
func (Permission) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Int("role_id"),
		field.Int("permission_rule_id"),
		field.String("type"),
	}
}

// Edges of the Permission.
func (Permission) Edges() []ent.Edge {
	return nil
}
