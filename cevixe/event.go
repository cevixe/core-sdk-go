package cevixe

import (
	"context"
	"github.com/cevixe/core-sdk-go/core"
)

func NewCommandEvent(ctx context.Context, data interface{}) core.Event {
	factory := ctx.Value(CevixeEventFactory).(core.EventFactory)
	return factory.NewCommandEvent(ctx, data)
}

func NewDomainEvent(ctx context.Context, entity core.Entity, data interface{}, state interface{}) core.Event {
	factory := ctx.Value(CevixeEventFactory).(core.EventFactory)
	return factory.NewDomainEvent(ctx, data, entity, state)
}

func NewBusinessEvent(ctx context.Context, data interface{}) core.Event {
	factory := ctx.Value(CevixeEventFactory).(core.EventFactory)
	return factory.NewBusinessEvent(ctx, data)
}
