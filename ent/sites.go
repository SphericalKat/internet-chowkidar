// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/gnulinuxindia/internet-chowkidar/ent/sites"
)

// Sites is the model entity for the Sites schema.
type Sites struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Domain holds the value of the "domain" field.
	Domain string `json:"domain,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SitesQuery when eager-loading is set.
	Edges        SitesEdges `json:"edges"`
	selectValues sql.SelectValues
}

// SitesEdges holds the relations/edges for other nodes in the graph.
type SitesEdges struct {
	// Blocks holds the value of the blocks edge.
	Blocks []*Blocks `json:"blocks,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// BlocksOrErr returns the Blocks value or an error if the edge
// was not loaded in eager-loading.
func (e SitesEdges) BlocksOrErr() ([]*Blocks, error) {
	if e.loadedTypes[0] {
		return e.Blocks, nil
	}
	return nil, &NotLoadedError{edge: "blocks"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Sites) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case sites.FieldID:
			values[i] = new(sql.NullInt64)
		case sites.FieldDomain:
			values[i] = new(sql.NullString)
		case sites.FieldCreatedAt, sites.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Sites fields.
func (s *Sites) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sites.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case sites.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case sites.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = value.Time
			}
		case sites.FieldDomain:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field domain", values[i])
			} else if value.Valid {
				s.Domain = value.String
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Sites.
// This includes values selected through modifiers, order, etc.
func (s *Sites) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryBlocks queries the "blocks" edge of the Sites entity.
func (s *Sites) QueryBlocks() *BlocksQuery {
	return NewSitesClient(s.config).QueryBlocks(s)
}

// Update returns a builder for updating this Sites.
// Note that you need to call Sites.Unwrap() before calling this method if this Sites
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Sites) Update() *SitesUpdateOne {
	return NewSitesClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Sites entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Sites) Unwrap() *Sites {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Sites is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Sites) String() string {
	var builder strings.Builder
	builder.WriteString("Sites(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(s.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("domain=")
	builder.WriteString(s.Domain)
	builder.WriteByte(')')
	return builder.String()
}

// SitesSlice is a parsable slice of Sites.
type SitesSlice []*Sites
