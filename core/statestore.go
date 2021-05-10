package core

import "context"

type StateStore interface {
	GetLastVersion(ctx context.Context, typ string, id string) Entity
	GetByVersion(ctx context.Context, typ string, id string, version uint64) Entity
}
