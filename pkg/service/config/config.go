package config

import (
	"github.com/eser/acik.io/pkg/bliss"
)

type AppConfig struct {
	bliss.BaseConfig

	AppName  string `conf:"NAME" default:"acik-service"`
	Postgres struct {
		Dsn string `conf:"DSN" default:"postgres://localhost:5432"`
	} `conf:"POSTGRES"`
}
