package schema

import "entgo.io/ent"

// Admin holds the schema definition for the Admin entity.
type Admin struct {
	ent.Schema
}

// Fields of the Admin.
func (Admin) Fields() []ent.Field {
	return nil
}

// Edges of the Admin.
func (Admin) Edges() []ent.Edge {
	return nil
}
