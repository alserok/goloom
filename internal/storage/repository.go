package storage

import (
	"context"
	"github.com/alserok/goloom/internal/service/models"
)

type Storage interface {
	FileStorage
	ServicesStorage
}

type FileStorage interface {
	SaveFile(ctx context.Context, path string, data []byte) error
	DeleteFile(ctx context.Context, path string) error
	GetFile(ctx context.Context, path string) ([]byte, error)
}

type ServicesStorage interface {
	UpdateServiceStatus(ctx context.Context, data models.ServiceState) error
	GetServicesInfo(ctx context.Context) ([]models.ServiceState, error)
	RemoveService(ctx context.Context, addr string) error
	AddService(ctx context.Context, addr string) error
}
