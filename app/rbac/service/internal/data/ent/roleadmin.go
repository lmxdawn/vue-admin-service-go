// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"vue-admin/app/rbac/service/internal/data/ent/roleadmin"

	"entgo.io/ent/dialect/sql"
)

// RoleAdmin is the model entity for the RoleAdmin schema.
type RoleAdmin struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// RoleID holds the value of the "role_id" field.
	RoleID int `json:"role_id,omitempty"`
	// AdminID holds the value of the "admin_id" field.
	AdminID int64 `json:"admin_id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RoleAdmin) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case roleadmin.FieldID, roleadmin.FieldRoleID, roleadmin.FieldAdminID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type RoleAdmin", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RoleAdmin fields.
func (ra *RoleAdmin) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case roleadmin.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ra.ID = int64(value.Int64)
		case roleadmin.FieldRoleID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field role_id", values[i])
			} else if value.Valid {
				ra.RoleID = int(value.Int64)
			}
		case roleadmin.FieldAdminID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field admin_id", values[i])
			} else if value.Valid {
				ra.AdminID = value.Int64
			}
		}
	}
	return nil
}

// Update returns a builder for updating this RoleAdmin.
// Note that you need to call RoleAdmin.Unwrap() before calling this method if this RoleAdmin
// was returned from a transaction, and the transaction was committed or rolled back.
func (ra *RoleAdmin) Update() *RoleAdminUpdateOne {
	return (&RoleAdminClient{config: ra.config}).UpdateOne(ra)
}

// Unwrap unwraps the RoleAdmin entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ra *RoleAdmin) Unwrap() *RoleAdmin {
	_tx, ok := ra.config.driver.(*txDriver)
	if !ok {
		panic("ent: RoleAdmin is not a transactional entity")
	}
	ra.config.driver = _tx.drv
	return ra
}

// String implements the fmt.Stringer.
func (ra *RoleAdmin) String() string {
	var builder strings.Builder
	builder.WriteString("RoleAdmin(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ra.ID))
	builder.WriteString("role_id=")
	builder.WriteString(fmt.Sprintf("%v", ra.RoleID))
	builder.WriteString(", ")
	builder.WriteString("admin_id=")
	builder.WriteString(fmt.Sprintf("%v", ra.AdminID))
	builder.WriteByte(')')
	return builder.String()
}

// RoleAdmins is a parsable slice of RoleAdmin.
type RoleAdmins []*RoleAdmin

func (ra RoleAdmins) config(cfg config) {
	for _i := range ra {
		ra[_i].config = cfg
	}
}
