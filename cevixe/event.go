package cevixe

import (
	"context"
	"github.com/cevixe/core-sdk-go/core"
)

func NewCommandEvent(ctx context.Context, data interface{}) core.Event {
	factory := ctx.Value(CevixeEventFactory).(core.EventFactory)
	return factory.NewCommandEvent(ctx, data)
}

func NewCommandEventWithCustomType(ctx context.Context, typ string, data interface{}) core.Event {
	factory := ctx.Value(CevixeEventFactory).(core.EventFactory)
	return factory.NewCommandEventWithCustomType(ctx, typ, data)
}

func NewDomainEvent(ctx context.Context, entity core.Entity, data interface{}, state interface{}) core.Event {
	factory := ctx.Value(CevixeEventFactory).(core.EventFactory)
	return factory.NewDomainEvent(ctx, data, entity, state)
}

func NewDomainEventWithCustomType(ctx context.Context, entity core.Entity, typ string, data interface{}, state interface{}) core.Event {
	factory := ctx.Value(CevixeEventFactory).(core.EventFactory)
	return factory.NewDomainEventWithCustomType(ctx, typ, data, entity, state)
}

func NewFirstDomainEvent(ctx context.Context, data interface{}, state interface{}) core.Event {
	factory := ctx.Value(CevixeEventFactory).(core.EventFactory)
	return factory.NewFirstDomainEvent(ctx, data, state)
}

func NewFirstDomainEventWithCustomType(ctx context.Context, typ string, data interface{}, state interface{}) core.Event {
	factory := ctx.Value(CevixeEventFactory).(core.EventFactory)
	return factory.NewFirstDomainEventWithCustomType(ctx, typ, data, state)
}

func NewFirstDomainEventWithCustomID(ctx context.Context, data interface{}, id string, state interface{}) core.Event {
	factory := ctx.Value(CevixeEventFactory).(core.EventFactory)
	return factory.NewFirstDomainEventWithCustomID(ctx, data, id, state)
}

func NewFirstDomainEventWithCustomIDAndCustomType(ctx context.Context, typ string, data interface{}, id string, state interface{}) core.Event {
	factory := ctx.Value(CevixeEventFactory).(core.EventFactory)
	return factory.NewFirstDomainEventWithCustomIDAndCustomType(ctx, typ, data, id, state)
}

func NewBusinessEvent(ctx context.Context, data interface{}) core.Event {
	factory := ctx.Value(CevixeEventFactory).(core.EventFactory)
	return factory.NewBusinessEvent(ctx, data)
}

func NewBusinessEventWithCustomID(ctx context.Context, data interface{}) core.Event {
	factory := ctx.Value(CevixeEventFactory).(core.EventFactory)
	return factory.NewBusinessEvent(ctx, data)
}
