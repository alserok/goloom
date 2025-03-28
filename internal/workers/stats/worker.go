package stats

import (
	"context"
	"github.com/alserok/goloom/pkg/logger"
	"runtime"
	"runtime/pprof"
	"time"
)

func New() *worker {
	return &worker{
		tickPeriod: time.Minute,
	}
}

type worker struct {
	tickPeriod time.Duration

	pprof.LabelSet
}

func (w worker) Start(ctx context.Context) {
	tick := time.NewTicker(w.tickPeriod)
	defer tick.Stop()

	log := logger.UnwrapLogger(ctx)

	log.Info("starting stats ✅ ")
	defer log.Info("closing stats ☑️ ")

	type Stats struct {
		Mem runtime.MemStats
	}

	for {
		select {
		case <-ctx.Done():
			return
		case <-tick.C:
			var stats Stats
			runtime.ReadMemStats(&stats.Mem)

			log.Info("stats",
				logger.WithArg("alloc", stats.Mem.Alloc/1024/1024),
				logger.WithArg("total_alloc", stats.Mem.TotalAlloc/1024/1024),
				logger.WithArg("sys", stats.Mem.Sys/1024/1024),
				logger.WithArg("num_gc", stats.Mem.NumGC),
				logger.WithArg("goroutines", runtime.NumGoroutine()),
				logger.WithArg("cpu", runtime.NumCPU()),
			)
		}
	}
}
