package worker

import (
	"context"
	"learnyscape-backend-mono/internal/config"
	"os"
	"os/signal"
	"syscall"
)

func Start() {
	cfg := config.InitConfig()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	runHttpWorker(ctx, cfg)
}
