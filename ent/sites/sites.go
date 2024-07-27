// Code generated by ent, DO NOT EDIT.

package sites

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the sites type in the database.
	Label = "sites"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDomain holds the string denoting the domain field in the database.
	FieldDomain = "domain"
	// EdgeBlocks holds the string denoting the blocks edge name in mutations.
	EdgeBlocks = "blocks"
	// Table holds the table name of the sites in the database.
	Table = "sites"
	// BlocksTable is the table that holds the blocks relation/edge.
	BlocksTable = "blocks"
	// BlocksInverseTable is the table name for the Blocks entity.
	// It exists in this package in order to avoid circular dependency with the "blocks" package.
	BlocksInverseTable = "blocks"
	// BlocksColumn is the table column denoting the blocks relation/edge.
	BlocksColumn = "site_id"
)

// Columns holds all SQL columns for sites fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDomain,
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the Sites queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDomain orders the results by the domain field.
func ByDomain(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDomain, opts...).ToFunc()
}

// ByBlocksCount orders the results by blocks count.
func ByBlocksCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newBlocksStep(), opts...)
	}
}

// ByBlocks orders the results by blocks terms.
func ByBlocks(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBlocksStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newBlocksStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BlocksInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, BlocksTable, BlocksColumn),
	)
}
