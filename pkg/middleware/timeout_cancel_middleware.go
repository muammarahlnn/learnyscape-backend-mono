package middleware

import (
	"context"
	"errors"
	"learnyscape-backend-mono/pkg/httperror"
	"time"

	"github.com/gin-gonic/gin"
)

func TimeoutCancelMiddleware(timeoutPeriod int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		duration := time.Duration(timeoutPeriod) * time.Second
		timeoutCtx, cancel := context.WithTimeout(ctx.Request.Context(), duration)
		defer cancel()

		ctx.Request = ctx.Request.WithContext(timeoutCtx)

		done := make(chan struct{})
		go next(ctx, done)

		select {
		case <-timeoutCtx.Done():
			if errors.Is(timeoutCtx.Err(), context.Canceled) {
				ctx.Error(httperror.NewCanceledError())
				ctx.Abort()
			} else if errors.Is(timeoutCtx.Err(), context.DeadlineExceeded) {
				ctx.Error(httperror.NewTimeoutError())
				ctx.Abort()
			}
		case <-done:
		}
	}
}

func next(ctx *gin.Context, done chan struct{}) {
	defer func() {
		close(done)
		if err, ok := recover().(error); ok && err != nil {
			ctx.Error(err)
			ctx.Abort()
		}
	}()

	ctx.Next()
}
