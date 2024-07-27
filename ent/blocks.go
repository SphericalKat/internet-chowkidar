// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/gnulinuxindia/internet-chowkidar/ent/blocks"
)

// Blocks is the model entity for the Blocks schema.
type Blocks struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// IP holds the value of the "ip" field.
	IP string `json:"ip,omitempty"`
	// Domain holds the value of the "domain" field.
	Domain       string `json:"domain,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Blocks) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case blocks.FieldID:
			values[i] = new(sql.NullInt64)
		case blocks.FieldIP, blocks.FieldDomain:
			values[i] = new(sql.NullString)
		case blocks.FieldCreatedAt, blocks.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Blocks fields.
func (b *Blocks) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case blocks.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			b.ID = int(value.Int64)
		case blocks.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				b.CreatedAt = value.Time
			}
		case blocks.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				b.UpdatedAt = value.Time
			}
		case blocks.FieldIP:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ip", values[i])
			} else if value.Valid {
				b.IP = value.String
			}
		case blocks.FieldDomain:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field domain", values[i])
			} else if value.Valid {
				b.Domain = value.String
			}
		default:
			b.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Blocks.
// This includes values selected through modifiers, order, etc.
func (b *Blocks) Value(name string) (ent.Value, error) {
	return b.selectValues.Get(name)
}

// Update returns a builder for updating this Blocks.
// Note that you need to call Blocks.Unwrap() before calling this method if this Blocks
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Blocks) Update() *BlocksUpdateOne {
	return NewBlocksClient(b.config).UpdateOne(b)
}

// Unwrap unwraps the Blocks entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Blocks) Unwrap() *Blocks {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Blocks is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Blocks) String() string {
	var builder strings.Builder
	builder.WriteString("Blocks(")
	builder.WriteString(fmt.Sprintf("id=%v, ", b.ID))
	builder.WriteString("created_at=")
	builder.WriteString(b.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(b.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("ip=")
	builder.WriteString(b.IP)
	builder.WriteString(", ")
	builder.WriteString("domain=")
	builder.WriteString(b.Domain)
	builder.WriteByte(')')
	return builder.String()
}

// BlocksSlice is a parsable slice of Blocks.
type BlocksSlice []*Blocks
