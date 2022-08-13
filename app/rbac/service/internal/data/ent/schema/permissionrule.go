package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// PermissionRule holds the schema definition for the PermissionRule entity.
type PermissionRule struct {
	ent.Schema
}

// Annotations 用户实体的注解
func (PermissionRule) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "permission_rule"},
	}
}

// Fields of the PermissionRule.
func (PermissionRule) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Int("pid"),
		field.String("name"),
		field.String("title"),
		field.Int("status"),
		field.String("condition"),
		field.Int("sort"),
		field.Time("create_time"),
		field.Time("update_time"),
	}
}

// Edges of the PermissionRule.
func (PermissionRule) Edges() []ent.Edge {
	return nil
}
