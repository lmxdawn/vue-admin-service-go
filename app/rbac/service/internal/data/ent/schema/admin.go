package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Admin holds the schema definition for the Admin entity.
type Admin struct {
	ent.Schema
}

// Annotations 用户实体的注解
func (Admin) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "admin"},
	}
}

// Fields of the Admin.
func (Admin) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("username"),
		field.String("password"),
		field.String("tel"),
		field.String("email"),
		field.String("avatar"),
		field.Int("sex"),
		field.String("last_login_ip"),
		field.Time("last_login_time"),
		field.Int("status"),
		field.Time("create_time"),
		field.Time("update_time"),
	}
}

// Edges of the Admin.
func (Admin) Edges() []ent.Edge {
	return nil
}
