package worker

import (
	"context"
	"learnyscape-backend-mono/internal/config"
	"learnyscape-backend-mono/internal/log"
	"learnyscape-backend-mono/internal/provider"
	"learnyscape-backend-mono/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

func Start() {
	cfg := config.InitConfig()
	log.SetLogger(logger.NewZeroLogLogger(cfg.Logger.Level))
	provider.BootstrapGlobal(cfg)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	runHttpWorker(ctx, cfg)
}
