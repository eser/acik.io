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
	"github.com/eser/acik.io/pkg/service/config"
	"github.com/eser/acik.io/pkg/service/routes/home"
	"github.com/eser/acik.io/pkg/service/routes/protected"
)

func LoadConfig(loader configfx.ConfigLoader) (*config.AppConfig, *logfx.Config, *httpfx.Config, *datafx.Config, error) { //nolint:lll
	appConfig := &config.AppConfig{} //nolint:exhaustruct

	err := loader.Load(
		appConfig,

		loader.FromJsonFile("config.json"),
		loader.FromEnvFile(".env"),
		loader.FromSystemEnv(),
	)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("failed to load config: %w", err)
	}

	return appConfig, &appConfig.Log, &appConfig.Http, &appConfig.Data, nil
}

func RegisterMiddlewares(routes httpfx.Router, httpMetrics *httpfx.Metrics, appConfig *config.AppConfig) error {
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

		RegisterMiddlewares,

		healthcheck.RegisterRoutes,
		openapi.RegisterRoutes,

		home.RegisterIndexRoute,
		protected.RegisterIndexRoute,
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
