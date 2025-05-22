package worker

import (
	"context"
	"learnyscape-backend-mono/internal/config"
	"learnyscape-backend-mono/internal/log"
	"learnyscape-backend-mono/internal/provider"
	"learnyscape-backend-mono/pkg/logger"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func Start() {
	cfg := config.InitConfig()
	log.SetLogger(logger.NewZeroLogLogger(cfg.Logger.Level))
	provider.BootstrapGlobal(cfg)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		runHttpWorker(ctx, cfg)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		runAMQPWorker(ctx, cfg)
	}()

	<-ctx.Done()
	wg.Wait()
}
