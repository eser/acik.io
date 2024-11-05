package logfx

import (
	"github.com/eser/acik.io/pkg/bliss/di"
)

func RegisterDependencies(container di.Container, config *Config) error {
	logger, err := NewLoggerAsDefault(config)
	if err != nil {
		return err
	}

	di.Register(container, logger)

	return nil
}
