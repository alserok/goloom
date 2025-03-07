package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/alserok/goloom/internal/service/models"
	"github.com/alserok/goloom/internal/storage"
)

type Service interface {
	ConfigService
	StatusService
}

type ConfigService interface {
	GetConfig(ctx context.Context, path string) (models.Config, error)
	UpdateConfig(ctx context.Context, path string, cfg models.Config) error
	CreateConfig(ctx context.Context, path string, cfg models.Config) error
	DeleteConfig(ctx context.Context, path string) error
}

type StatusService interface {
	UpdateStatus(ctx context.Context, data models.ServiceState) error
	GetStatuses(ctx context.Context) ([]models.ServiceState, error)
}

func New(repo storage.Storage) Service {
	return &service{
		repo: repo,
	}
}

type service struct {
	repo storage.Storage
}

func (s service) GetConfig(ctx context.Context, path string) (models.Config, error) {
	file, err := s.repo.GetFile(ctx, path)
	if err != nil {
		return models.Config{}, fmt.Errorf("failed to get file: %w", err)
	}

	var cfg models.Config
	if err = json.Unmarshal(file, &cfg); err != nil {
		// TODO
	}

	return cfg, nil
}

func (s service) UpdateConfig(ctx context.Context, path string, cfg models.Config) error {
	b, err := json.Marshal(cfg)
	if err != nil {
		// TODO
	}

	if err = s.repo.SaveFile(ctx, path, b); err != nil {
		return fmt.Errorf("failed to save file: %w", err)
	}

	return nil
}

func (s service) CreateConfig(ctx context.Context, path string, cfg models.Config) error {
	b, err := json.Marshal(cfg)
	if err != nil {
		// TODO
	}

	if err = s.repo.SaveFile(ctx, path, b); err != nil {
		return fmt.Errorf("failed to save file: %w", err)
	}

	return nil
}

func (s service) DeleteConfig(ctx context.Context, path string) error {
	if err := s.repo.DeleteFile(ctx, path); err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}

	return nil
}

func (s service) UpdateStatus(ctx context.Context, data models.ServiceState) error {
	if err := s.repo.UpdateStatus(ctx, data); err != nil {
		return fmt.Errorf("failed to update status: %w", err)
	}

	return nil
}

func (s service) GetStatuses(ctx context.Context) ([]models.ServiceState, error) {
	statuses, err := s.repo.GetStatuses(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get statuses: %w", err)
	}

	return statuses, nil
}
