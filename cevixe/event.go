package cevixe

import (
	"context"
	"github.com/cevixe/core-sdk-go/core"
)

func NewEvent(ctx context.Context, entity core.Entity, state interface{}, payload interface{}) core.Event {
	factory := ctx.Value(CevixeEventFactory).(core.EventFactory)
	return factory.NewEvent(ctx, entity, state, payload)
}
