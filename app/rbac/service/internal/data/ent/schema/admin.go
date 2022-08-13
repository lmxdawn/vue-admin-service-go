package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Admin holds the schema definition for the Admin entity.
type Admin struct {
	ent.Schema
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
