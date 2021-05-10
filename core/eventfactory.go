package core

import "context"

type EventFactory interface {
	NewCommandEvent(ctx context.Context, data interface{}) Event
	NewBusinessEvent(ctx context.Context, data interface{}) Event
	NewDomainEvent(ctx context.Context, data interface{}, entity Entity, state interface{}) Event
}
