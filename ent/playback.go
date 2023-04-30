// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/zibbp/ganymede/ent/playback"
	"github.com/zibbp/ganymede/internal/utils"
)

// Playback is the model entity for the Playback schema.
type Playback struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// VodID holds the value of the "vod_id" field.
	VodID uuid.UUID `json:"vod_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// Time holds the value of the "time" field.
	Time int `json:"time,omitempty"`
	// Status holds the value of the "status" field.
	Status utils.PlaybackStatus `json:"status,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt    time.Time `json:"created_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Playback) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case playback.FieldTime:
			values[i] = new(sql.NullInt64)
		case playback.FieldStatus:
			values[i] = new(sql.NullString)
		case playback.FieldUpdatedAt, playback.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case playback.FieldID, playback.FieldVodID, playback.FieldUserID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Playback fields.
func (pl *Playback) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case playback.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pl.ID = *value
			}
		case playback.FieldVodID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field vod_id", values[i])
			} else if value != nil {
				pl.VodID = *value
			}
		case playback.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				pl.UserID = *value
			}
		case playback.FieldTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field time", values[i])
			} else if value.Valid {
				pl.Time = int(value.Int64)
			}
		case playback.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				pl.Status = utils.PlaybackStatus(value.String)
			}
		case playback.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pl.UpdatedAt = value.Time
			}
		case playback.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pl.CreatedAt = value.Time
			}
		default:
			pl.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Playback.
// This includes values selected through modifiers, order, etc.
func (pl *Playback) Value(name string) (ent.Value, error) {
	return pl.selectValues.Get(name)
}

// Update returns a builder for updating this Playback.
// Note that you need to call Playback.Unwrap() before calling this method if this Playback
// was returned from a transaction, and the transaction was committed or rolled back.
func (pl *Playback) Update() *PlaybackUpdateOne {
	return NewPlaybackClient(pl.config).UpdateOne(pl)
}

// Unwrap unwraps the Playback entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pl *Playback) Unwrap() *Playback {
	_tx, ok := pl.config.driver.(*txDriver)
	if !ok {
		panic("ent: Playback is not a transactional entity")
	}
	pl.config.driver = _tx.drv
	return pl
}

// String implements the fmt.Stringer.
func (pl *Playback) String() string {
	var builder strings.Builder
	builder.WriteString("Playback(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pl.ID))
	builder.WriteString("vod_id=")
	builder.WriteString(fmt.Sprintf("%v", pl.VodID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", pl.UserID))
	builder.WriteString(", ")
	builder.WriteString("time=")
	builder.WriteString(fmt.Sprintf("%v", pl.Time))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", pl.Status))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pl.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(pl.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Playbacks is a parsable slice of Playback.
type Playbacks []*Playback
