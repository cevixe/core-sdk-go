package core

import (
	"context"
	"time"
)

type StateStore interface {
	GetLastVersion(ctx context.Context, typ string, id string) Entity
	GetByVersion(ctx context.Context, typ string, id string, version uint64) Entity

	GetVersions(ctx context.Context, typ string, id string,
		after *time.Time, before *time.Time, nextToken *string, limit *int64) VersionPage
	GetByType(ctx context.Context, typ string,
		after *time.Time, before *time.Time, nextToken *string, limit *int64) EntityPage
}
