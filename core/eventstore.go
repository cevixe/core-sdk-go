package core

import (
	"context"
	"time"
)

type EventStore interface {
	GetLastEvent(ctx context.Context, source string) Event
	GetEventByID(ctx context.Context, source string, id string) Event
	
	GetEntityEvents(ctx context.Context, typ string, id string,
		after *time.Time, before *time.Time, nextToken *string, limit *int64) EventPage
	GetDayEvents(ctx context.Context, day string,
		after *time.Time, before *time.Time, nextToken *string, limit *int64) EventPage
	GetTypeEvents(ctx context.Context, typ string,
		after *time.Time, before *time.Time, nextToken *string, limit *int64) EventPage
	GetAuthorEvents(ctx context.Context, author string,
		after *time.Time, before *time.Time, nextToken *string, limit *int64) EventPage
	GetTransactionEvents(ctx context.Context, transaction string,
		after *time.Time, before *time.Time, nextToken *string, limit *int64) EventPage
}
