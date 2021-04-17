package core

import "context"

type EventStore interface {
	GetEvent(ctx context.Context, source string, id *uint64) Event
	SaveEvent(ctx context.Context, event Event)
}
