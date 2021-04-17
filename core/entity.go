package core

import "time"

type Entity interface {
	ID() string
	Type() string
	Time() time.Time
	Owner() string
	Version() uint64
	State(interface{})
}
