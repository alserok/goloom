package broadcaster

import "context"

type Broadcaster interface {
	Broadcast(ctx context.Context, data any) error
	AddTarget(ctx context.Context, target string) error
	RemoveTarget(ctx context.Context, target string) error
}
