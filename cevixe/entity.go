package cevixe

import (
	"context"
	"github.com/cevixe/core-sdk-go/core"
)

func Entity(ctx context.Context, typ string, id string) core.Entity {
	store := ctx.Value(CevixeEventStore).(core.EventStore)
	source := "/" + typ + "/" + id
	event := store.GetEvent(ctx, source, nil)
	return event.Source()
}
