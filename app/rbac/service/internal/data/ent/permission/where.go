// Code generated by ent, DO NOT EDIT.

package permission

import (
	"vue-admin/app/rbac/service/internal/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// RoleID applies equality check predicate on the "role_id" field. It's identical to RoleIDEQ.
func RoleID(v int) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRoleID), v))
	})
}

// PermissionRuleID applies equality check predicate on the "permission_rule_id" field. It's identical to PermissionRuleIDEQ.
func PermissionRuleID(v int) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPermissionRuleID), v))
	})
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// RoleIDEQ applies the EQ predicate on the "role_id" field.
func RoleIDEQ(v int) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRoleID), v))
	})
}

// RoleIDNEQ applies the NEQ predicate on the "role_id" field.
func RoleIDNEQ(v int) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRoleID), v))
	})
}

// RoleIDIn applies the In predicate on the "role_id" field.
func RoleIDIn(vs ...int) predicate.Permission {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldRoleID), v...))
	})
}

// RoleIDNotIn applies the NotIn predicate on the "role_id" field.
func RoleIDNotIn(vs ...int) predicate.Permission {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldRoleID), v...))
	})
}

// RoleIDGT applies the GT predicate on the "role_id" field.
func RoleIDGT(v int) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRoleID), v))
	})
}

// RoleIDGTE applies the GTE predicate on the "role_id" field.
func RoleIDGTE(v int) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRoleID), v))
	})
}

// RoleIDLT applies the LT predicate on the "role_id" field.
func RoleIDLT(v int) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRoleID), v))
	})
}

// RoleIDLTE applies the LTE predicate on the "role_id" field.
func RoleIDLTE(v int) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRoleID), v))
	})
}

// PermissionRuleIDEQ applies the EQ predicate on the "permission_rule_id" field.
func PermissionRuleIDEQ(v int) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPermissionRuleID), v))
	})
}

// PermissionRuleIDNEQ applies the NEQ predicate on the "permission_rule_id" field.
func PermissionRuleIDNEQ(v int) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPermissionRuleID), v))
	})
}

// PermissionRuleIDIn applies the In predicate on the "permission_rule_id" field.
func PermissionRuleIDIn(vs ...int) predicate.Permission {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPermissionRuleID), v...))
	})
}

// PermissionRuleIDNotIn applies the NotIn predicate on the "permission_rule_id" field.
func PermissionRuleIDNotIn(vs ...int) predicate.Permission {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPermissionRuleID), v...))
	})
}

// PermissionRuleIDGT applies the GT predicate on the "permission_rule_id" field.
func PermissionRuleIDGT(v int) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPermissionRuleID), v))
	})
}

// PermissionRuleIDGTE applies the GTE predicate on the "permission_rule_id" field.
func PermissionRuleIDGTE(v int) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPermissionRuleID), v))
	})
}

// PermissionRuleIDLT applies the LT predicate on the "permission_rule_id" field.
func PermissionRuleIDLT(v int) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPermissionRuleID), v))
	})
}

// PermissionRuleIDLTE applies the LTE predicate on the "permission_rule_id" field.
func PermissionRuleIDLTE(v int) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPermissionRuleID), v))
	})
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldType), v))
	})
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.Permission {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldType), v...))
	})
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.Permission {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldType), v...))
	})
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldType), v))
	})
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldType), v))
	})
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldType), v))
	})
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldType), v))
	})
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldType), v))
	})
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldType), v))
	})
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldType), v))
	})
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldType), v))
	})
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldType), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Permission) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Permission) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Permission) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		p(s.Not())
	})
}