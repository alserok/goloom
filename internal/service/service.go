package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/alserok/goloom/internal/broadcaster"
	"github.com/alserok/goloom/internal/service/models"
	"github.com/alserok/goloom/internal/storage"
	"github.com/alserok/goloom/internal/utils"
	"github.com/alserok/goloom/static/pages"
	"gopkg.in/yaml.v3"
	"os"
	p "path"
	"strings"
	"time"
)

type Service interface {
	ConfigService
	PageService
	ServicesService
}

type ServicesService interface {
	AddService(ctx context.Context, addr string) error
	RemoveService(ctx context.Context, addr string) error
	UpdateServiceStatus(ctx context.Context, data models.ServiceState) error
	GetServices(ctx context.Context) ([]models.ServiceState, error)
}

type PageService interface {
	GetConfigPage(ctx context.Context, path string) ([]byte, error)
	GetDirPage(ctx context.Context, path string) ([]byte, error)
	GetStatusesPage(ctx context.Context) ([]byte, error)
}

type ConfigService interface {
	UpdateConfig(ctx context.Context, path string, content string) error
	CreateConfig(ctx context.Context, path string, cfg models.Config) error
	DeleteConfig(ctx context.Context, path string) error
}

func New(repo storage.Storage, pagesConstructor pages.HTMLConstructor, broadcaster broadcaster.Broadcaster) Service {
	return &service{
		repo:             repo,
		pagesConstructor: pagesConstructor,
		broadcaster:      broadcaster,
	}
}

type service struct {
	repo storage.Storage

	pagesConstructor pages.HTMLConstructor

	broadcaster broadcaster.Broadcaster
}

func (s service) GetServices(ctx context.Context) ([]models.ServiceState, error) {
	info, err := s.repo.GetServicesInfo(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get statuses: %w", err)
	}

	return info, nil
}

func (s service) AddService(ctx context.Context, addr string) error {
	if err := s.repo.AddService(ctx, addr); err != nil {
		return fmt.Errorf("failed to add service: %w", err)
	}

	return nil
}

func (s service) RemoveService(ctx context.Context, addr string) error {
	if err := s.repo.RemoveService(ctx, addr); err != nil {
		return fmt.Errorf("failed to remove service: %w", err)
	}

	return nil
}

func (s service) GetDirPage(ctx context.Context, path string) ([]byte, error) {
	absPath := fmt.Sprintf("./data/%s", path)

	info, err := os.Stat(absPath)
	if err != nil {
		return nil, utils.NewError(err.Error(), utils.ErrNotFound)
	}

	dir := models.Dir{
		Path: path,
		Name: info.Name(),
	}

	if info.IsDir() {
		entries, err := os.ReadDir(absPath)
		if err != nil {
			return nil, utils.NewError(err.Error(), utils.ErrNotFound)
		}

		for _, entry := range entries {
			content := models.Content{
				IsDir: entry.IsDir(),
				Name:  entry.Name(),
			}

			if content.IsDir {
				size := utils.CountSize(fmt.Sprintf("%s/%s", absPath, entry.Name()))
				content.Size = utils.PrettifySize(size)
			} else {
				entryInfo, _ := entry.Info()
				content.Size = utils.PrettifySize(entryInfo.Size())
			}

			dir.Content = append(dir.Content, content)
		}
	}

	data := pages.Data{
		"dir":  dir,
		"path": fmt.Sprintf("./%s", dir.Path),
	}

	page, err := s.pagesConstructor.Render(pages.PageDir, data)
	if err != nil {
		return nil, fmt.Errorf("failed to generate dir page: %w", err)
	}

	return page, nil
}

func (s service) GetConfigPage(ctx context.Context, path string) ([]byte, error) {
	file, err := s.repo.GetFile(ctx, path)
	if err != nil {
		return nil, fmt.Errorf("failed to get file: %w", err)
	}

	var content []byte
	if len(file) > 0 {
		switch {
		case strings.HasSuffix(path, ".json"):
			var cfg models.Config

			if err = json.Unmarshal(file, &cfg); err != nil {
				return nil, utils.NewError(err.Error(), utils.ErrInternal)
			}

			content, err = json.MarshalIndent(cfg, " ", " ")
			if err != nil {
				return nil, utils.NewError(err.Error(), utils.ErrInternal)
			}
		case strings.HasSuffix(path, ".yaml"), strings.HasSuffix(path, ".yml"):
			var cfg models.Config

			if err = yaml.Unmarshal(file, &cfg); err != nil {
				return nil, utils.NewError(err.Error(), utils.ErrInternal)
			}

			content, err = yaml.Marshal(cfg)
			if err != nil {
				return nil, utils.NewError(err.Error(), utils.ErrInternal)
			}
		}
	}

	_, filename := p.Split(path)

	data := pages.Data{
		"content":  string(content),
		"path":     path,
		"filename": filename,
	}

	page, err := s.pagesConstructor.Render(pages.PageConfig, data)
	if err != nil {
		return nil, fmt.Errorf("failed to generate config page: %w", err)
	}

	return page, nil
}

func (s service) UpdateConfig(ctx context.Context, path string, content string) error {
	if err := s.repo.SaveFile(ctx, path, []byte(content)); err != nil {
		return fmt.Errorf("failed to save file: %w", err)
	}

	cfg, err := s.repo.GetFile(ctx, path)
	if err != nil {
		return fmt.Errorf("failed to get file: %w", err)
	}

	if err = s.broadcaster.Broadcast(ctx, broadcaster.Body{"path": path, "contentBytes": cfg}); err != nil {
		return utils.NewError(err.Error(), utils.ErrInternal)
	}

	return nil
}

func (s service) CreateConfig(ctx context.Context, path string, cfg models.Config) error {
	b, err := json.Marshal(cfg)
	if err != nil {
		return utils.NewError(err.Error(), utils.ErrInternal)
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

func (s service) UpdateServiceStatus(ctx context.Context, data models.ServiceState) error {
	if err := s.repo.UpdateServiceStatus(ctx, data); err != nil {
		return fmt.Errorf("failed to update status: %w", err)
	}

	return nil
}

func (s service) GetStatusesPage(ctx context.Context) ([]byte, error) {
	statesInfo, err := s.repo.GetServicesInfo(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get statuses: %w", err)
	}

	data := pages.Data{
		"states": statesInfo,
		"time":   time.Now().Format("2006-01-02 15:04:05"),
	}

	page, err := s.pagesConstructor.Render(pages.PageState, data)
	if err != nil {
		return nil, fmt.Errorf("failed to generate state page: %w", err)
	}

	return page, nil
}
