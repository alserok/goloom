package health_state

import (
	"context"
	"fmt"
	"github.com/alserok/goloom/internal/service"
	"github.com/alserok/goloom/internal/service/models"
	"github.com/alserok/goloom/pkg/logger"
	"net/http"
	"time"
)

const (
	healthRoute = "/health"
)

func New(targets []string, srvc service.StatusService) *worker {
	return &worker{
		targets: targets,
		cl:      http.DefaultClient,
		srvc:    srvc,
	}
}

type worker struct {
	targets []string

	srvc service.StatusService

	cl *http.Client
}

func (w *worker) Start(ctx context.Context) {
	tick := time.NewTicker(time.Second * 5)
	defer tick.Stop()

	log := logger.UnwrapLogger(ctx)

	for {
		select {
		case <-tick.C:
			for _, target := range w.targets {
				req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", target, healthRoute), nil)
				if err != nil {
					log.Error("failed to init request", logger.WithArg("error", err.Error()))
					continue
				}

				res, err := w.cl.Do(req)
				if err != nil {
					log.Error("failed to do request", logger.WithArg("error", err.Error()))
					continue
				}

				log.Info("checked status", logger.WithArg("service", target), logger.WithArg("status", res.StatusCode))

				if err = w.srvc.UpdateStatus(context.Background(), models.ServiceState{
					Service: target,
					Status:  res.StatusCode,
				}); err != nil {
					log.Error("failed to update status", logger.WithArg("error", err.Error()))
					continue
				}
			}
		case <-ctx.Done():
			return
		}
	}
}
