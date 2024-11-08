package main

import (
	"context"
	"fmt"

	"github.com/eser/acik.io/pkg/bliss/configfx"
	"github.com/eser/acik.io/pkg/bliss/datafx"
	"github.com/eser/acik.io/pkg/bliss/di"
	"github.com/eser/acik.io/pkg/bliss/httpfx"
	"github.com/eser/acik.io/pkg/bliss/httpfx/middlewares"
	"github.com/eser/acik.io/pkg/bliss/httpfx/modules/healthcheck"
	"github.com/eser/acik.io/pkg/bliss/httpfx/modules/openapi"
	"github.com/eser/acik.io/pkg/bliss/lib"
	"github.com/eser/acik.io/pkg/bliss/logfx"
	"github.com/eser/acik.io/pkg/bliss/metricsfx"
	"github.com/eser/acik.io/pkg/service"
	"github.com/eser/acik.io/pkg/service/broadcast"
)

func LoadConfig(loader configfx.ConfigLoader) (*service.AppConfig, *logfx.Config, *httpfx.Config, *datafx.Config, error) { //nolint:lll
	appConfig := &service.AppConfig{} //nolint:exhaustruct

	err := loader.LoadDefaults(appConfig)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("failed to load config: %w", err)
	}

	return appConfig, &appConfig.Log, &appConfig.Http, &appConfig.Data, nil
}

func RegisterHttpMiddlewares(routes httpfx.Router, httpMetrics *httpfx.Metrics, appConfig *service.AppConfig) error {
	routes.Use(middlewares.ErrorHandlerMiddleware())
	routes.Use(middlewares.ResolveAddressMiddleware())
	routes.Use(middlewares.ResponseTimeMiddleware())
	routes.Use(middlewares.CorrelationIdMiddleware())
	routes.Use(middlewares.CorsMiddleware())
	routes.Use(middlewares.MetricsMiddleware(httpMetrics))

	return nil
}

func main() {
	err := di.RegisterFn(
		di.Default,
		configfx.RegisterDependencies,
		LoadConfig,

		logfx.RegisterDependencies,
		metricsfx.RegisterDependencies,
		httpfx.RegisterDependencies,
		datafx.RegisterDependencies,

		RegisterHttpMiddlewares,

		healthcheck.RegisterHttpRoutes,
		openapi.RegisterHttpRoutes,

		broadcast.RegisterHttpRoutes,
	)
	if err != nil {
		panic(err)
	}

	run := di.CreateInvoker(
		di.Default,
		func(
			httpService httpfx.HttpService,
		) error {
			ctx := context.Background()

			cleanup, err := httpService.Start(ctx)
			if err != nil {
				return err //nolint:wrapcheck
			}

			lib.WaitForSignal()

			cleanup()

			return nil
		},
	)

	di.Seal(di.Default)

	err = run()
	if err != nil {
		panic(err)
	}
}
