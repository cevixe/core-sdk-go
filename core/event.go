package core

import "time"

type EventClass string

const (
	CommandEvent  EventClass = "C"
	DomainEvent   EventClass = "D"
	BusinessEvent EventClass = "B"
	SystemEvent   EventClass = "S"
)

type Event interface {
	ID() string
	Source() string
	Class() EventClass
	Type() string
	Time() time.Time
	Author() string
	Data(interface{})
	Entity() Entity
	Trigger() Event
	Transaction() string
}

type EventPage interface {
	Items() []Event
	NextToken() string
}
