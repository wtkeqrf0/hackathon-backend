// Code generated by ent, DO NOT EDIT.

package company

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the company type in the database.
	Label = "company"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "inn"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldWebsite holds the string denoting the website field in the database.
	FieldWebsite = "website"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// UserFieldID holds the string denoting the ID field of the User.
	UserFieldID = "id"
	// Table holds the table name of the company in the database.
	Table = "companies"
	// UsersTable is the table that holds the users relation/edge.
	UsersTable = "users"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
	// UsersColumn is the table column denoting the users relation/edge.
	UsersColumn = "company_inn"
)

// Columns holds all SQL columns for company fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldWebsite,
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
	// WebsiteValidator is a validator for the "website" field. It is called by the builders before save.
	WebsiteValidator func(string) error
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// OrderOption defines the ordering options for the Company queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByWebsite orders the results by the website field.
func ByWebsite(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWebsite, opts...).ToFunc()
}

// ByUsersField orders the results by users field.
func ByUsersField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUsersStep(), sql.OrderByField(field, opts...))
	}
}
func newUsersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UsersInverseTable, UserFieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, UsersTable, UsersColumn),
	)
}
