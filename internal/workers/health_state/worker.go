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

func New(tickPeriod time.Duration, broadcaster broadcaster.Broadcaster, srvc service.ServicesService) *worker {
	return &worker{
		tickPeriod:  tickPeriod,
		cl:          http.DefaultClient,
		srvc:        srvc,
		broadcaster: broadcaster,
	}
}

type worker struct {
	tickPeriod time.Duration

	srvc service.ServicesService

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

			services, err := w.srvc.GetServices(context.TODO())
			if err != nil {
				log.Error("failed to get statuses", logger.WithArg("error", err.Error()))
				continue
			}

			for _, srvc := range services {
				req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%s/%s", srvc.Addr, healthRoute), nil)
				if err != nil {
					log.Error("failed to init request", logger.WithArg("error", err.Error()))
					continue
				}

				var status int

				res, err := w.cl.Do(req)
				if err != nil {
					w.broadcaster.RemoveTarget(ctx, srvc.Addr)

					status = http.StatusInternalServerError
					failedReqs = append(failedReqs, srvc.Addr)
				} else {
					w.broadcaster.AddTargets(ctx, srvc.Addr)

					status = res.StatusCode
					succeededReqs = append(succeededReqs, srvc.Addr)
				}

				if err = w.srvc.UpdateServiceStatus(context.Background(), models.ServiceState{
					Addr:   srvc.Addr,
					Status: status,
				}); err != nil {
					log.Error("failed to update status", logger.WithArg("error", err.Error()))
					continue
				}
			}

			log.Info("worker checked states 🛠️",
				logger.WithArg("services", len(services)),
				logger.WithArg("succeeded", succeededReqs),
				logger.WithArg("failed", failedReqs),
			)
		case <-ctx.Done():
			return
		}
	}
}
