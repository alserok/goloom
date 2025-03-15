package workers

import (
	"context"
	"github.com/alserok/goloom/pkg/logger"
	"sync"
)

type Worker interface {
	Start(ctx context.Context)
}

type Launcher interface {
	Stop()
	Launch()
}

type launcher struct {
	// this context controls workers execution time
	ctx   context.Context
	close func()

	workers []Worker

	mu sync.Mutex
}

func NewLauncher(log logger.Logger, workers ...Worker) Launcher {
	ctx, cancel := context.WithCancel(context.Background())

	ctx = logger.WrapLogger(ctx, log)

	return &launcher{
		ctx:     ctx,
		close:   cancel,
		workers: workers,
	}
}

func (l *launcher) Launch() {
	for _, w := range l.workers {
		go w.Start(l.ctx)
	}
}

func (l *launcher) Stop() {
	l.close()
}
