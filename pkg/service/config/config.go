package config

import (
	"github.com/eser/acik.io/pkg/bliss"
)

type AppConfig struct {
	AppName  string `conf:"NAME" default:"acik.io"`
	Postgres struct {
		Dsn string `conf:"DSN" default:"postgres://localhost:5432"`
	} `conf:"POSTGRES"`
	bliss.BaseConfig
}
