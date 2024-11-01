package service

import (
	"fmt"

	"github.com/eser/acik.io/pkg/bliss/configfx"
	"github.com/eser/acik.io/pkg/bliss/datafx"
	"github.com/eser/acik.io/pkg/bliss/di"
	"github.com/eser/acik.io/pkg/bliss/httpfx"
	"github.com/eser/acik.io/pkg/bliss/httpfx/middlewares"
	"github.com/eser/acik.io/pkg/bliss/logfx"
	"github.com/eser/acik.io/pkg/bliss/metricsfx"
	"github.com/eser/acik.io/pkg/service/config"
)

func LoadConfig(loader configfx.ConfigLoader) (*config.AppConfig, *logfx.Config, *httpfx.Config, error) {
	appConfig := &config.AppConfig{} //nolint:exhaustruct

	err := loader.Load(
		appConfig,

		loader.FromJsonFile("config.json"),
		loader.FromEnvFile(".env"),
		loader.FromSystemEnv(),
	)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to load config: %w", err)
	}

	return appConfig, &appConfig.Log, &appConfig.Http, nil
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

func Startup(container di.Container) error {
	// TODO(@eser) load config
	// TODO(@eser) load logger
	// TODO(@eser) load metrics
	// TODO(@eser) load database if any
	// TODO(@eser) load redis if any
	// TODO(@eser) load http router
	// TODO(@eser) load http service

	err := di.RegisterFn(
		container,
		configfx.RegisterDependencies,
		LoadConfig,

		logfx.RegisterDependencies,
		metricsfx.RegisterDependencies,
		httpfx.RegisterDependencies,
		datafx.RegisterDependencies,

		// RegisterMiddlewares,

		// healthcheck.RegisterRoutes,
		// openapi.RegisterRoutes,

		// home.RegisterIndexRoute,
		// protected.RegisterIndexRoute,
	)

	return err
}

func Run() {
	err := Startup(di.Default)
	if err != nil {
		panic(err)
	}
}
