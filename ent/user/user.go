// Code generated by ent, DO NOT EDIT.

package user

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldSurname holds the string denoting the surname field in the database.
	FieldSurname = "surname"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// EdgeUserRoles holds the string denoting the user_roles edge name in mutations.
	EdgeUserRoles = "user_roles"
	// EdgeUserPermissions holds the string denoting the user_permissions edge name in mutations.
	EdgeUserPermissions = "user_permissions"
	// Table holds the table name of the user in the database.
	Table = "users"
	// UserRolesTable is the table that holds the user_roles relation/edge.
	UserRolesTable = "user_roles"
	// UserRolesInverseTable is the table name for the UserRole entity.
	// It exists in this package in order to avoid circular dependency with the "userrole" package.
	UserRolesInverseTable = "user_roles"
	// UserRolesColumn is the table column denoting the user_roles relation/edge.
	UserRolesColumn = "user_id"
	// UserPermissionsTable is the table that holds the user_permissions relation/edge.
	UserPermissionsTable = "user_permissions"
	// UserPermissionsInverseTable is the table name for the UserPermission entity.
	// It exists in this package in order to avoid circular dependency with the "userpermission" package.
	UserPermissionsInverseTable = "user_permissions"
	// UserPermissionsColumn is the table column denoting the user_permissions relation/edge.
	UserPermissionsColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldSurname,
	FieldUsername,
	FieldEmail,
	FieldPhone,
	FieldPassword,
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

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// SurnameValidator is a validator for the "surname" field. It is called by the builders before save.
	SurnameValidator func(string) error
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// BySurname orders the results by the surname field.
func BySurname(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSurname, opts...).ToFunc()
}

// ByUsername orders the results by the username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByPhone orders the results by the phone field.
func ByPhone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhone, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByUserRolesCount orders the results by user_roles count.
func ByUserRolesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUserRolesStep(), opts...)
	}
}

// ByUserRoles orders the results by user_roles terms.
func ByUserRoles(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserRolesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByUserPermissionsCount orders the results by user_permissions count.
func ByUserPermissionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUserPermissionsStep(), opts...)
	}
}

// ByUserPermissions orders the results by user_permissions terms.
func ByUserPermissions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserPermissionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newUserRolesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserRolesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, UserRolesTable, UserRolesColumn),
	)
}
func newUserPermissionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserPermissionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, UserPermissionsTable, UserPermissionsColumn),
	)
}
