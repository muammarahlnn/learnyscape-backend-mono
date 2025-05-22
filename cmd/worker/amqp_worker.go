package worker

import (
	"context"
	"learnyscape-backend-mono/internal/config"
	"learnyscape-backend-mono/internal/server"
)

func runAMQPWorker(ctx context.Context, cfg *config.Config) {
	srv := server.NewAMQPServer(cfg)
	go srv.Start()

	<-ctx.Done()
	srv.Shutdown()
}
