package core

import (
	"context"
	"time"
)

type StateStore interface {
	GetLastVersion(ctx context.Context, typ string, id string) Entity
	GetByVersion(ctx context.Context, typ string, id string, version uint64) Entity

	GetVersions(ctx context.Context, typ string, id string,
		after *uint64, before *uint64, nextToken *string, limit *int64) VersionPage
	GetByType(ctx context.Context, typ string,
		after *time.Time, nextToken *string, limit *int64) EntityPage
}
