package core

import "context"

type StateStore interface {
	UpdateState(ctx context.Context, entities []Entity)
}
