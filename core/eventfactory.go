package core

import "context"

type EventFactory interface {
	NewEvent(ctx context.Context, entity Entity, payload interface{}, state interface{}) Event
}
