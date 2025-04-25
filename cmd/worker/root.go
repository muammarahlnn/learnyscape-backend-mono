package worker

import (
	"context"
	"learnyscape-backend-mono/internal/config"
	"learnyscape-backend-mono/internal/provider"
	"os"
	"os/signal"
	"syscall"
)

func Start() {
	cfg := config.InitConfig()
	provider.BootstrapGlobal(cfg)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	runHttpWorker(ctx, cfg)
}
