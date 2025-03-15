package storage

import (
	"context"
	"github.com/alserok/goloom/internal/service/models"
)

type Storage interface {
	FileStorage
	StateStorage
}

type FileStorage interface {
	SaveFile(ctx context.Context, path string, data []byte) error
	DeleteFile(ctx context.Context, path string) error
	GetFile(ctx context.Context, path string) ([]byte, error)
}

type StateStorage interface {
	UpdateStatus(ctx context.Context, data models.ServiceState) error
	GetStatuses(ctx context.Context) ([]models.ServiceState, error)
}
