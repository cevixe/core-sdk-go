package core

import "context"

type EventFactory interface {
	NewSystemEvent(ctx context.Context, data interface{}) Event
	NewCommandEvent(ctx context.Context, data interface{}) Event
	NewBusinessEvent(ctx context.Context, data interface{}) Event
	NewDomainEvent(ctx context.Context, data interface{}, entity Entity, state interface{}) Event
	NewFirstDomainEvent(ctx context.Context, data interface{}, state interface{}) Event
	NewFirstDomainEventWithCustomID(ctx context.Context, data interface{}, id string, state interface{}) Event
}
