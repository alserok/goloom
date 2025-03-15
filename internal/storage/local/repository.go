package local

import (
	"context"
	"fmt"
	"github.com/alserok/goloom/internal/service/models"
	"github.com/alserok/goloom/internal/utils"
	"io"
	"os"
	"sync"
)

func NewRepository(dir string) *repository {
	return &repository{
		dir:          dir,
		stateStorage: make(map[string]int),
	}
}

type repository struct {
	dir string

	mu           sync.RWMutex
	stateStorage map[string]int
}

func (r *repository) UpdateStatus(ctx context.Context, data models.ServiceState) error {
	r.mu.Lock()
	r.stateStorage[data.Service] = data.Status
	r.mu.Unlock()

	return nil
}

func (r *repository) GetStatuses(ctx context.Context) ([]models.ServiceState, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	states := make([]models.ServiceState, 0, len(r.stateStorage))

	for srvc, status := range r.stateStorage {
		states = append(states, models.ServiceState{
			Service: srvc,
			Status:  status,
		})
	}

	return states, nil
}

func (r *repository) SaveFile(ctx context.Context, path string, data []byte) error {
	f, err := os.OpenFile(fmt.Sprintf("%s/%s", r.dir, path), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0777)
	if err != nil {
		// TODO
	}

	if _, err = f.Write(data); err != nil {
		// TODO
	}

	return nil
}

func (r *repository) DeleteFile(ctx context.Context, path string) error {
	if err := os.Remove(fmt.Sprintf("%s/%s", r.dir, path)); err != nil {
		// TODO
	}

	return nil
}

func (r *repository) GetFile(ctx context.Context, path string) ([]byte, error) {
	f, err := os.OpenFile(path, os.O_RDONLY, 0777)
	if err != nil {
		return nil, utils.NewError(err.Error(), utils.ErrInternal)
	}

	b, err := io.ReadAll(f)
	if err != nil {
		return nil, utils.NewError(err.Error(), utils.ErrInternal)
	}

	return b, nil
}
