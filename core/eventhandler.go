package core

import "context"

type EventHandler func(ctx context.Context, event Event) Event
