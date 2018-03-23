package msg

import (
	"fmt"
	"strings"
	"time"
)

//Represents a generic event
type Event interface {
	Type() string
	Package() string
	Value() interface{}
	Timestamp() time.Time
}

//event represents an event
type event struct {
	timestamp time.Time
	value     interface{}
}

func (e *event) Value() interface{} {
	return e.value
}

func (e *event) Timestamp() time.Time {
	return e.timestamp
}

func (e *event) Package() string {
	var eventType = fmt.Sprintf("%T", e.value)
	var fragments = strings.Split(eventType, ".")
	return strings.Replace(fragments[0], "*", "", 1)
}

//Type returns event type (simple package and struct name)
func (e *event) Type() string {
	var eventType = fmt.Sprintf("%T", e.value)
	eventType = strings.Replace(eventType, "*", "", len(eventType))
	var fragments = strings.Split(eventType, ".")
	if len(fragments) > 2 {
		fragments = fragments[len(fragments)-2:]
	}
	return strings.Join(fragments, "_")
}

//NewEvent creates a new event
func NewEvent(value interface{}) *event {
	return &event{
		timestamp: time.Now(),
		value:     value,
	}
}

//NewEvent creates a new event
func NewEventWithTimestamp(value interface{}, timestamp time.Time) Event {
	return &event{
		timestamp: timestamp,
		value:     value,
	}
}
