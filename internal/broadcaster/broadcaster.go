package broadcaster

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/alserok/goloom/internal/utils"
	"net/http"
	"sync"
)

type Broadcaster interface {
	Broadcast(ctx context.Context, data any) error
	AddTarget(ctx context.Context, target string) error
	RemoveTarget(ctx context.Context, target string) error
}

func New() Broadcaster {
	return &broadcaster{}
}

type broadcaster struct {
	targets map[string]struct{}
	mu      sync.RWMutex

	cl *http.Client
}

func (b *broadcaster) Broadcast(ctx context.Context, data any) error {
	body, err := json.Marshal(data)
	if err != nil {
		return utils.NewError(err.Error(), utils.ErrInternal)
	}

	b.mu.RLock()
	defer b.mu.RUnlock()

	for srvc, _ := range b.targets {
		req, err := http.NewRequest(http.MethodPost, srvc, bytes.NewReader(body))
		if err != nil {
			return utils.NewError(err.Error(), utils.ErrInternal)
		}

		_, err = b.cl.Do(req)
		if err != nil {
			return utils.NewError(err.Error(), utils.ErrInternal)
		}
	}

	return nil
}

func (b *broadcaster) AddTarget(ctx context.Context, target string) error {
	b.mu.Lock()
	b.targets[target] = struct{}{}
	b.mu.Unlock()

	return nil
}

func (b *broadcaster) RemoveTarget(ctx context.Context, target string) error {
	b.mu.Lock()
	delete(b.targets, target)
	b.mu.Unlock()

	return nil
}
