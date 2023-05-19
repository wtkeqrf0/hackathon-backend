// Code generated by ent, DO NOT EDIT.

package economicactivity

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the economicactivity type in the database.
	Label = "economic_activity"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "main"
	// FieldSubs holds the string denoting the subs field in the database.
	FieldSubs = "subs"
	// Table holds the table name of the economicactivity in the database.
	Table = "economic_activities"
)

// Columns holds all SQL columns for economicactivity fields.
var Columns = []string{
	FieldID,
	FieldSubs,
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
	// SubsValidator is a validator for the "subs" field. It is called by the builders before save.
	SubsValidator func(string) error
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// OrderOption defines the ordering options for the EconomicActivity queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// BySubs orders the results by the subs field.
func BySubs(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSubs, opts...).ToFunc()
}