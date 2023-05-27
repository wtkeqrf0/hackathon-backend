// Code generated by ent, DO NOT EDIT.

package businessactivity

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the businessactivity type in the database.
	Label = "business_activity"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldSubType holds the string denoting the sub_type field in the database.
	FieldSubType = "sub_type"
	// FieldTotal holds the string denoting the total field in the database.
	FieldTotal = "total"
	// EdgeHistories holds the string denoting the histories edge name in mutations.
	EdgeHistories = "histories"
	// Table holds the table name of the businessactivity in the database.
	Table = "business_activities"
	// HistoriesTable is the table that holds the histories relation/edge.
	HistoriesTable = "histories"
	// HistoriesInverseTable is the table name for the History entity.
	// It exists in this package in order to avoid circular dependency with the "history" package.
	HistoriesInverseTable = "histories"
	// HistoriesColumn is the table column denoting the histories relation/edge.
	HistoriesColumn = "business_activity_id"
)

// Columns holds all SQL columns for businessactivity fields.
var Columns = []string{
	FieldID,
	FieldType,
	FieldSubType,
	FieldTotal,
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
	// TotalValidator is a validator for the "total" field. It is called by the builders before save.
	TotalValidator func(float64) error
)

// OrderOption defines the ordering options for the BusinessActivity queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// BySubType orders the results by the sub_type field.
func BySubType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSubType, opts...).ToFunc()
}

// ByTotal orders the results by the total field.
func ByTotal(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotal, opts...).ToFunc()
}

// ByHistoriesCount orders the results by histories count.
func ByHistoriesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newHistoriesStep(), opts...)
	}
}

// ByHistories orders the results by histories terms.
func ByHistories(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newHistoriesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newHistoriesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(HistoriesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, HistoriesTable, HistoriesColumn),
	)
}