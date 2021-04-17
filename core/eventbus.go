package core

import "context"

type EventBus interface {
	PublishEvent(ctx context.Context, event Event)
}
