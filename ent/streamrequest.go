// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/skkugoon/strattonight/ent/streamrequest"
)

// StreamRequest is the model entity for the StreamRequest schema.
type StreamRequest struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// RequestID holds the value of the "request_id" field.
	RequestID int `json:"request_id,omitempty"`
	// RequestType holds the value of the "request_type" field.
	RequestType string `json:"request_type,omitempty"`
	// IsActive holds the value of the "is_active" field.
	IsActive     bool `json:"is_active,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*StreamRequest) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case streamrequest.FieldIsActive:
			values[i] = new(sql.NullBool)
		case streamrequest.FieldID, streamrequest.FieldRequestID:
			values[i] = new(sql.NullInt64)
		case streamrequest.FieldRequestType:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the StreamRequest fields.
func (sr *StreamRequest) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case streamrequest.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sr.ID = int(value.Int64)
		case streamrequest.FieldRequestID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field request_id", values[i])
			} else if value.Valid {
				sr.RequestID = int(value.Int64)
			}
		case streamrequest.FieldRequestType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field request_type", values[i])
			} else if value.Valid {
				sr.RequestType = value.String
			}
		case streamrequest.FieldIsActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_active", values[i])
			} else if value.Valid {
				sr.IsActive = value.Bool
			}
		default:
			sr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the StreamRequest.
// This includes values selected through modifiers, order, etc.
func (sr *StreamRequest) Value(name string) (ent.Value, error) {
	return sr.selectValues.Get(name)
}

// Update returns a builder for updating this StreamRequest.
// Note that you need to call StreamRequest.Unwrap() before calling this method if this StreamRequest
// was returned from a transaction, and the transaction was committed or rolled back.
func (sr *StreamRequest) Update() *StreamRequestUpdateOne {
	return NewStreamRequestClient(sr.config).UpdateOne(sr)
}

// Unwrap unwraps the StreamRequest entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sr *StreamRequest) Unwrap() *StreamRequest {
	_tx, ok := sr.config.driver.(*txDriver)
	if !ok {
		panic("ent: StreamRequest is not a transactional entity")
	}
	sr.config.driver = _tx.drv
	return sr
}

// String implements the fmt.Stringer.
func (sr *StreamRequest) String() string {
	var builder strings.Builder
	builder.WriteString("StreamRequest(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sr.ID))
	builder.WriteString("request_id=")
	builder.WriteString(fmt.Sprintf("%v", sr.RequestID))
	builder.WriteString(", ")
	builder.WriteString("request_type=")
	builder.WriteString(sr.RequestType)
	builder.WriteString(", ")
	builder.WriteString("is_active=")
	builder.WriteString(fmt.Sprintf("%v", sr.IsActive))
	builder.WriteByte(')')
	return builder.String()
}

// StreamRequests is a parsable slice of StreamRequest.
type StreamRequests []*StreamRequest