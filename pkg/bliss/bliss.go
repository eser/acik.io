package bliss

import (
	"log/slog"

	"github.com/eser/acik.io/pkg/bliss/configfx"
	"github.com/eser/acik.io/pkg/bliss/di"
)

func Load() {
	// TODO(@eser) load config
	// TODO(@eser) load logger
	// TODO(@eser) load metrics
	// TODO(@eser) load database if any
	// TODO(@eser) load redis if any
	// TODO(@eser) load http router
	// TODO(@eser) load http service
	di.Register(di.Default, configfx.NewConfigLoader())
	di.Register(di.Default, slog.Default())

	loggingInvoker := di.Default.CreateInvoker(func(logger *slog.Logger) {
		logger.Info("Hello, World!")
	})

	loggingInvoker()
}
