package identitysvc

import (
	"context"
	"fmt"

	"github.com/eser/acik.io/pkg/bliss/configfx"
	"github.com/eser/acik.io/pkg/bliss/datafx"
	"github.com/eser/acik.io/pkg/bliss/di"
	"github.com/eser/acik.io/pkg/bliss/grpcfx"
	"github.com/eser/acik.io/pkg/bliss/lib"
	"github.com/eser/acik.io/pkg/bliss/logfx"
	"github.com/eser/acik.io/pkg/bliss/metricsfx"
)

func LoadConfig(loader configfx.ConfigLoader) (*AppConfig, *logfx.Config, *grpcfx.Config, *datafx.Config, error) { //nolint:lll
	appConfig := &AppConfig{} //nolint:exhaustruct

	err := loader.LoadDefaults(appConfig)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("failed to load config: %w", err)
	}

	return appConfig, &appConfig.Log, &appConfig.Grpc, &appConfig.Data, nil
}

func Run() error {
	err := di.RegisterFn(
		di.Default,
		configfx.RegisterDependencies,
		LoadConfig,

		logfx.RegisterDependencies,
		metricsfx.RegisterDependencies,
		grpcfx.RegisterDependencies,
		datafx.RegisterDependencies,

		RegisterGrpcService,
	)
	if err != nil {
		panic(err)
	}

	run := di.CreateInvoker(
		di.Default,
		func(
			grpcService grpcfx.GrpcService,
		) error {
			ctx := context.Background()

			cleanup, err := grpcService.Start(ctx)
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

	return err
}
