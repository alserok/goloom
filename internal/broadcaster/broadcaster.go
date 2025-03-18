package broadcaster

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/alserok/goloom/internal/utils"
	"net/http"
	"sync"
)

type Broadcaster interface {
	Broadcast(ctx context.Context, data any) error
	AddTargets(ctx context.Context, target ...string)
	RemoveTarget(ctx context.Context, target string)
}

func New() Broadcaster {
	return &broadcaster{
		targets: make(map[string]struct{}),
		cl:      http.DefaultClient,
	}
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
		req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%s/provide", srvc), bytes.NewReader(body))
		if err != nil {
			return utils.NewError(err.Error(), utils.ErrInternal)
		}

		_, _ = b.cl.Do(req)
	}

	return nil
}

func (b *broadcaster) AddTargets(ctx context.Context, targets ...string) {
	b.mu.Lock()
	for _, target := range targets {
		b.targets[target] = struct{}{}
	}
	b.mu.Unlock()
}

func (b *broadcaster) RemoveTarget(ctx context.Context, target string) {
	b.mu.Lock()
	delete(b.targets, target)
	b.mu.Unlock()
}

type Body map[string]any
