package worker

import (
	"context"
	"learnyscape-backend-mono/internal/config"
	"learnyscape-backend-mono/internal/server"
)

func runHttpWorker(ctx context.Context, cfg *config.Config) {
	srv := server.NewHttpServer(cfg)
	go srv.Start()

	<-ctx.Done()
	srv.Shutdown()
}
