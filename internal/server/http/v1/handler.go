package v1

import (
	"encoding/json"
	"fmt"
	"github.com/alserok/goloom/internal/service"
	"github.com/alserok/goloom/internal/service/models"
	"github.com/alserok/goloom/internal/utils"
	"net"
	"net/http"
	"strings"
)

func newHandler(srvc service.Service) *handler {
	return &handler{
		service: services{
			services: srvc,
			pages:    srvc,
			config:   srvc,
		},
	}
}

type handler struct {
	service services
}

type services struct {
	services service.ServicesService
	pages    service.PageService
	config   service.ConfigService
}

func (h *handler) AddService(w http.ResponseWriter, r *http.Request) error {
	host, port, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return utils.NewError(err.Error(), utils.ErrInternal)
	}

	if queryPort := r.URL.Query().Get("port"); port != "" {
		port = queryPort
	}

	addr := fmt.Sprintf("%s:%s", host, port)

	if err := h.service.services.AddService(r.Context(), addr); err != nil {
		return fmt.Errorf("service failed to add service: %w", err)
	}

	return nil
}

func (h *handler) RemoveService(w http.ResponseWriter, r *http.Request) error {
	host, port, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return utils.NewError(err.Error(), utils.ErrInternal)
	}

	if queryPort := r.URL.Query().Get("port"); port != "" {
		port = queryPort
	}

	addr := fmt.Sprintf("%s:%s", host, port)

	if err := h.service.services.RemoveService(r.Context(), addr); err != nil {
		return fmt.Errorf("service failed to remove service: %w", err)
	}

	return nil
}

func (h *handler) UpdateConfig(w http.ResponseWriter, r *http.Request) error {
	var req models.UpdateConfigReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return utils.NewError(err.Error(), utils.ErrBadRequest)
	}

	path := strings.TrimPrefix(r.URL.Path, "/config/update")

	if err := h.service.config.UpdateConfig(r.Context(), path, req.Content); err != nil {
		return fmt.Errorf("service failed to update config: %w", err)
	}

	return nil
}
func (h *handler) GetConfig(w http.ResponseWriter, r *http.Request) error {
	path := r.PathValue("path")

	cfg, err := h.service.pages.GetConfigPage(r.Context(), path)
	if err != nil {
		return fmt.Errorf("service failed to delete config: %w", err)
	}

	if err = json.NewEncoder(w).Encode(cfg); err != nil {
		return utils.NewError(err.Error(), utils.ErrInternal)
	}

	return nil
}

func (h *handler) DeleteConfig(w http.ResponseWriter, r *http.Request) error {
	path := strings.TrimPrefix(r.URL.Path, "/config/delete")

	err := h.service.config.DeleteConfig(r.Context(), path)
	if err != nil {
		return fmt.Errorf("service failed to delete config: %w", err)
	}

	return nil
}

func (h *handler) GetDirPage(w http.ResponseWriter, r *http.Request) error {
	path := strings.TrimPrefix(r.URL.Path, "/web/config/dir/")

	page, err := h.service.pages.GetDirPage(r.Context(), path)
	if err != nil {
		return fmt.Errorf("service failed to get config: %w", err)
	}

	if _, err = w.Write(page); err != nil {
		return utils.NewError(err.Error(), utils.ErrInternal)
	}

	return nil
}

func (h *handler) GetConfigPage(w http.ResponseWriter, r *http.Request) error {
	path := strings.TrimPrefix(r.URL.Path, "/web/config/file/")

	page, err := h.service.pages.GetConfigPage(r.Context(), path)
	if err != nil {
		return fmt.Errorf("service failed to get config: %w", err)
	}

	if _, err = w.Write(page); err != nil {
		return utils.NewError(err.Error(), utils.ErrInternal)
	}

	return nil
}

func (h *handler) GetStatePage(w http.ResponseWriter, r *http.Request) error {
	page, err := h.service.pages.GetStatusesPage(r.Context())
	if err != nil {
		return fmt.Errorf("service failed to get statuses: %w", err)
	}

	if _, err = w.Write(page); err != nil {
		return utils.NewError(err.Error(), utils.ErrInternal)
	}

	return nil
}
