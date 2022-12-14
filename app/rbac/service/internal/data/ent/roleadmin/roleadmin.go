// Code generated by ent, DO NOT EDIT.

package roleadmin

const (
	// Label holds the string label denoting the roleadmin type in the database.
	Label = "role_admin"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRoleID holds the string denoting the role_id field in the database.
	FieldRoleID = "role_id"
	// FieldAdminID holds the string denoting the admin_id field in the database.
	FieldAdminID = "admin_id"
	// Table holds the table name of the roleadmin in the database.
	Table = "role_admin"
)

// Columns holds all SQL columns for roleadmin fields.
var Columns = []string{
	FieldID,
	FieldRoleID,
	FieldAdminID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
