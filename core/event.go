package core

import "time"

type Event interface {
	ID() uint64
	Type() string
	Time() time.Time
	Author() string
	Payload(interface{})
	Source() Entity
	Trigger() Event
	Transaction() string
}
