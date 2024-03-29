package core

import "time"

type Entity interface {
	ID() string
	Type() string
	Version() uint64
	State(interface{})
	UpdatedAt() time.Time
	UpdatedBy() string
	CreatedAt() time.Time
	CreatedBy() string
	IsDeleted() bool
}

type EntityPage interface {
	Items() []Entity
	NextToken() string
}
