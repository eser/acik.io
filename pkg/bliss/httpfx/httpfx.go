package httpfx

import (
	"github.com/eser/acik.io/pkg/bliss/di"
	"github.com/eser/acik.io/pkg/bliss/metricsfx"
)

func RegisterDependencies(container di.Container, config *Config, mp metricsfx.MetricsProvider) {
	routes := NewRouter("/")
	httpService := NewHttpService(config, routes, mp)

	di.RegisterFor[Router](container, routes)
	di.RegisterFor[HttpService](container, httpService)
}
