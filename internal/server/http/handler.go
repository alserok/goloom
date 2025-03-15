package http

import (
	"encoding/json"
	"fmt"
	"github.com/alserok/goloom/internal/service"
	"github.com/alserok/goloom/internal/service/models"
	"github.com/alserok/goloom/internal/utils"
	"net/http"
	"strings"
)

func newHandler(srvc service.Service) *handler {
	return &handler{
		srvc: srvc,
	}
}

type handler struct {
	srvc service.Service
}

func (h *handler) UpdateConfig(w http.ResponseWriter, r *http.Request) error {
	var req models.UpdateConfigReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return utils.NewError(err.Error(), utils.ErrBadRequest)
	}

	path := strings.TrimPrefix(r.URL.Path, "/config/update")

	if err := h.srvc.UpdateConfig(r.Context(), path, req.Content); err != nil {
		return fmt.Errorf("service failed to update config: %w", err)
	}

	return nil
}
func (h *handler) GetConfig(w http.ResponseWriter, r *http.Request) error {
	path := r.PathValue("path")

	cfg, err := h.srvc.GetConfigPage(r.Context(), path)
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

	err := h.srvc.DeleteConfig(r.Context(), path)
	if err != nil {
		return fmt.Errorf("service failed to delete config: %w", err)
	}

	return nil
}

func (h *handler) GetDirPage(w http.ResponseWriter, r *http.Request) error {
	path := strings.TrimPrefix(r.URL.Path, "/web/config/dir/")

	page, err := h.srvc.GetDirPage(r.Context(), path)
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

	page, err := h.srvc.GetConfigPage(r.Context(), path)
	if err != nil {
		return fmt.Errorf("service failed to get config: %w", err)
	}

	if _, err = w.Write(page); err != nil {
		return utils.NewError(err.Error(), utils.ErrInternal)
	}

	return nil
}

func (h *handler) GetStatePage(w http.ResponseWriter, r *http.Request) error {
	page, err := h.srvc.GetStatusesPage(r.Context())
	if err != nil {
		return fmt.Errorf("service failed to get statuses: %w", err)
	}

	if _, err = w.Write(page); err != nil {
		return utils.NewError(err.Error(), utils.ErrInternal)
	}

	return nil
}
