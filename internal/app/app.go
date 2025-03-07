package app

import (
	"context"
	"fmt"
	"github.com/alserok/goloom/internal/config"
	"github.com/alserok/goloom/internal/server"
	"github.com/alserok/goloom/internal/service"
	"github.com/alserok/goloom/internal/storage/files"
	"github.com/alserok/goloom/internal/workers"
	state "github.com/alserok/goloom/internal/workers/health_state"
	"github.com/alserok/goloom/pkg/logger"
	"os/signal"
	"syscall"

	_ "github.com/joho/godotenv/autoload"
)

func MustStart(cfg *config.Config) {
	log := logger.NewLogger(logger.Slog, cfg.Env)
	defer func() {
		_ = log.Close()
	}()

	log.Info("starting goloom 🚧", logger.WithArg("port", cfg.Port))

	repo := files.NewRepository(files.MustSetup(cfg.Storage.Dir))
	srvc := service.New(repo)
	serv := server.New(server.HTTP, srvc)

	launcher := workers.NewLauncher(log, state.New(cfg.Services, srvc))
	defer launcher.Stop()
	launcher.Launch()

	log.Info("goloom is running ✅ ", logger.WithArg("web", fmt.Sprintf("http://127.0.0.1:%s/web/state", cfg.Port)))

	run(serv, cfg.Port)

	log.Info("goloom is closed ❌ ")
}

func run(s server.Server, port string) {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	go s.MustServe(port)

	<-ctx.Done()

	_ = s.Shutdown()
}
