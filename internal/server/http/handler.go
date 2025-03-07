package http

import (
	"encoding/json"
	"fmt"
	"github.com/alserok/goloom/internal/service"
	"github.com/alserok/goloom/internal/service/models"
	"github.com/alserok/goloom/static/pages"
	"net/http"
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
		// TODO
	}

	if err := h.srvc.UpdateConfig(r.Context(), req.Path, req.Config); err != nil {
		// TODO
	}

	return nil
}
func (h *handler) GetConfig(w http.ResponseWriter, r *http.Request) error {
	path := r.PathValue("path")

	cfg, err := h.srvc.GetConfig(r.Context(), path)
	if err != nil {
		// TODO
	}

	if err = json.NewEncoder(w).Encode(cfg); err != nil {
		// TODO
	}

	return nil
}

func (h *handler) DeleteConfig(w http.ResponseWriter, r *http.Request) error {
	path := r.PathValue("path")

	err := h.srvc.DeleteConfig(r.Context(), path)
	if err != nil {
		// TODO
	}

	return nil
}

func (h *handler) GetConfigPage(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h *handler) GetStatePage(w http.ResponseWriter, r *http.Request) error {
	//states, err := h.srvc.GetStatuses(r.Context())
	//if err != nil {
	//	return fmt.Errorf("service failed to get statuses: %w", err)
	//}
	states := []models.ServiceState{
		{
			Status:  200,
			Service: "a",
		},
		{
			Status:  500,
			Service: "b",
		},
		{
			Status:  200,
			Service: "c",
		},
		{
			Status:  500,
			Service: "d",
		},
	}

	page, err := pages.NewStatePage(r.Context(), states)
	if err != nil {
		return fmt.Errorf("failed to generate state page: %w", err)
	}

	_, err = w.Write([]byte(page))
	if err != nil {
		// TODO
	}

	return nil
}
