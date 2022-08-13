// Code generated by ent, DO NOT EDIT.

package admin

const (
	// Label holds the string label denoting the admin type in the database.
	Label = "admin"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldTel holds the string denoting the tel field in the database.
	FieldTel = "tel"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldAvatar holds the string denoting the avatar field in the database.
	FieldAvatar = "avatar"
	// FieldSex holds the string denoting the sex field in the database.
	FieldSex = "sex"
	// FieldLastLoginIP holds the string denoting the last_login_ip field in the database.
	FieldLastLoginIP = "last_login_ip"
	// FieldLastLoginTime holds the string denoting the last_login_time field in the database.
	FieldLastLoginTime = "last_login_time"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// Table holds the table name of the admin in the database.
	Table = "admins"
)

// Columns holds all SQL columns for admin fields.
var Columns = []string{
	FieldID,
	FieldUsername,
	FieldPassword,
	FieldTel,
	FieldEmail,
	FieldAvatar,
	FieldSex,
	FieldLastLoginIP,
	FieldLastLoginTime,
	FieldStatus,
	FieldCreateTime,
	FieldUpdateTime,
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
