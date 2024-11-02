package httpfx

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"

	"github.com/eser/acik.io/pkg/bliss/metricsfx"
)

type HttpService interface {
	Server() *http.Server
	Router() Router

	Start(ctx context.Context) (func(), error)
}

type HttpServiceImpl struct {
	InnerServer  *http.Server
	InnerRouter  Router
	InnerMetrics *Metrics

	Config *Config
}

var _ HttpService = (*HttpServiceImpl)(nil)

func NewHttpService(config *Config, router Router, mp metricsfx.MetricsProvider) *HttpServiceImpl {
	server := &http.Server{ //nolint:exhaustruct
		ReadHeaderTimeout: config.ReadHeaderTimeout,
		ReadTimeout:       config.ReadTimeout,
		WriteTimeout:      config.WriteTimeout,
		IdleTimeout:       config.IdleTimeout,

		Addr: config.Addr,

		Handler: router.GetMux(),
	}

	return &HttpServiceImpl{
		InnerServer:  server,
		InnerRouter:  router,
		InnerMetrics: NewMetrics(mp),
		Config:       config,
	}
}

func (hs *HttpServiceImpl) Server() *http.Server {
	return hs.InnerServer
}

func (hs *HttpServiceImpl) Router() Router { //nolint:ireturn
	return hs.InnerRouter
}

func (hs *HttpServiceImpl) Start(ctx context.Context) (func(), error) {
	slog.InfoContext(ctx, "HttpService is starting...", slog.String("addr", hs.Config.Addr))

	listener, lnErr := net.Listen("tcp", hs.InnerServer.Addr)
	if lnErr != nil {
		return nil, fmt.Errorf("HttpService Net Listen error: %w", lnErr)
	}

	serverErrChan := make(chan error, 1)

	go func() {
		if sErr := hs.Server().Serve(listener); sErr != nil && !errors.Is(sErr, http.ErrServerClosed) {
			serverErrChan <- fmt.Errorf("HttpService Serve error: %w", sErr)
		}

		close(serverErrChan)
	}()

	if err := <-serverErrChan; err != nil {
		listener.Close() //nolint:errcheck,gosec

		return nil, err
	}

	cleanup := func() {
		slog.InfoContext(ctx, "Shutting down server...")

		newCtx, cancel := context.WithTimeout(ctx, hs.Config.GracefulShutdownTimeout)
		defer cancel()

		if err := hs.InnerServer.Shutdown(newCtx); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.ErrorContext(ctx, "HttpService forced to shutdown", slog.Any("error", err))

			return
		}

		slog.InfoContext(ctx, "HttpService has gracefully stopped.")
	}

	return cleanup, nil
}
