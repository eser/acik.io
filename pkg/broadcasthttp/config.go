package broadcasthttp

import (
	"github.com/eser/acik.io/pkg/bliss"
)

type AppConfig struct {
	AppName string `conf:"NAME" default:"broadcasthttp"`
	bliss.BaseConfig
}
