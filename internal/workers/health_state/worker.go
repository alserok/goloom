package health_state

import (
	"context"
	"fmt"
	"github.com/alserok/goloom/internal/broadcaster"
	"github.com/alserok/goloom/internal/service"
	"github.com/alserok/goloom/internal/service/models"
	"github.com/alserok/goloom/pkg/logger"
	"net/http"
	"time"
)

const (
	healthRoute = "health"
)

func New(broadcaster broadcaster.Broadcaster, targets []string, tickPeriod time.Duration, srvc service.StatusService) *worker {
	return &worker{
		tickPeriod:  tickPeriod,
		targets:     targets,
		cl:          http.DefaultClient,
		srvc:        srvc,
		broadcaster: broadcaster,
	}
}

type worker struct {
	tickPeriod time.Duration

	targets []string

	srvc service.StatusService

	cl *http.Client

	broadcaster broadcaster.Broadcaster
}

func (w *worker) Start(ctx context.Context) {
	tick := time.NewTicker(w.tickPeriod)
	defer tick.Stop()

	log := logger.UnwrapLogger(ctx)

	log.Info("starting 'health state' worker ✳️")
	defer log.Info("closing 'health state' worker ✳️")

	for {
		select {
		case <-tick.C:
			var failedReqs, succeededReqs []string

			for _, target := range w.targets {
				req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%s/%s", target, healthRoute), nil)
				if err != nil {
					log.Error("failed to init request", logger.WithArg("error", err.Error()))
					continue
				}

				var status int

				res, err := w.cl.Do(req)
				if err != nil {
					_ = w.broadcaster.RemoveTarget(ctx, target)

					status = http.StatusInternalServerError
					failedReqs = append(failedReqs, target)
				} else {
					_ = w.broadcaster.AddTarget(ctx, target)

					status = res.StatusCode
					succeededReqs = append(succeededReqs, target)
				}

				if err = w.srvc.UpdateStatus(context.Background(), models.ServiceState{
					Service: target,
					Status:  status,
				}); err != nil {
					log.Error("failed to update status", logger.WithArg("error", err.Error()))
					continue
				}
			}

			log.Info("worker checked states 🛠️",
				logger.WithArg("services", len(w.targets)),
				logger.WithArg("succeeded", succeededReqs),
				logger.WithArg("failed", failedReqs),
			)
		case <-ctx.Done():
			return
		}
	}
}
