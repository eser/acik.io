package logfx

import (
	"github.com/eser/acik.io/pkg/bliss/di"
)

// var FxModule = fx.Module( //nolint:gochecknoglobals
// 	"log",
// 	fx.Provide(
// 		FxNew,
// 	),
// )

func Startup(container di.Container, config *Config) error {
	logger, err := NewLogger(config)
	if err != nil {
		return err
	}

	di.Register(container, logger)

	return nil
}
