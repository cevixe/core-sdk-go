package core

import "context"

type EventStore interface {
	GetLastEvent(ctx context.Context, source string) Event
	GetEventByID(ctx context.Context, source string, id string) Event
	GetEntityEvents(ctx context.Context, source string, nextToken string) EventPage
	GetTransactionEvents(ctx context.Context, source string, nextToken string) EventPage
}
