package core

import "time"

type Version interface {
	ID() uint64
	Time() time.Time
	Author() string
}

type VersionPage interface {
	Items() []Version
	NextToken() string
}
