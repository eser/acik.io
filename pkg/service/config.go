package service

import (
	"github.com/eser/acik.io/pkg/bliss"
)

type AppConfig struct {
	AppName string `conf:"NAME" default:"acik.io"`
	bliss.BaseConfig
}
