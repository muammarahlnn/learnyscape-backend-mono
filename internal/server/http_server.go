package server

import (
	"context"
	"errors"
	"fmt"
	"learnyscape-backend-mono/internal/config"
	"learnyscape-backend-mono/internal/log"
	"learnyscape-backend-mono/internal/provider"
	"learnyscape-backend-mono/pkg/middleware"
	validationutil "learnyscape-backend-mono/pkg/util/validation"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type HttpServer struct {
	cfg *config.Config
	srv *http.Server
}

func NewHttpServer(cfg *config.Config) *HttpServer {
	gin.SetMode(cfg.App.Envinronment)

	router := gin.New()
	router.ContextWithFallback = true
	router.HandleMethodNotAllowed = true

	registerValidators()
	registerMiddleware(router, cfg)
	provider.BootstrapHttp(cfg, router)

	return &HttpServer{
		cfg: cfg,
		srv: &http.Server{
			Addr:    fmt.Sprintf("%s:%d", cfg.HttpServer.Host, cfg.HttpServer.Port),
			Handler: router,
		},
	}
}

func (s *HttpServer) Start() {
	log.Logger.Info("Running HTTP server on port:", s.cfg.HttpServer.Port)

	err := s.srv.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Logger.Fatal("Error while HTTP server listening:", err)
	}

	log.Logger.Info("HTTP server is not receiving new requests...:")
}

func (s *HttpServer) Shutdown() {
	timeout := time.Duration(s.cfg.HttpServer.GracePeriod) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	log.Logger.Info("Attempting  to shutdown HTTP server...")
	if err := s.srv.Shutdown(ctx); err != nil {
		log.Logger.Fatal("Error shutting down HTTP server:", err)
	}

	log.Logger.Info("HTTP server shudown gracefully")
}

func registerValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(validationutil.TagNameFormatter)
		v.RegisterValidation("password", validationutil.Password)
	}
}

func registerMiddleware(router *gin.Engine, cfg *config.Config) {
	middlewares := []gin.HandlerFunc{
		gin.Recovery(),
		middleware.LoggerMiddleware(log.Logger),
		gzip.Gzip(gzip.BestSpeed),
		middleware.ErrorMiddleware(),
		middleware.TimeoutCancelMiddleware(cfg.HttpServer.RequestTimeoutPeriod),
		cors.New(cors.Config{
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTION", "PATCH", "HEAD"},
			AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
			AllowAllOrigins:  true,
			AllowCredentials: true,
		}),
	}

	router.Use(middlewares...)
}
