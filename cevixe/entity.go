package cevixe

import (
	"context"
	"github.com/cevixe/core-sdk-go/core"
)

func Entity(ctx context.Context, typ string, id string) core.Entity {
	store := ctx.Value(CevixeStateStore).(core.StateStore)
	entity := store.GetLastVersion(ctx, typ, id)
	return entity
}

func EntityVersion(ctx context.Context, typ string, id string, version uint64) core.Entity {
	store := ctx.Value(CevixeStateStore).(core.StateStore)
	entity := store.GetByVersion(ctx, typ, id, version)
	return entity
}
