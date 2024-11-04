package grpcfx

import (
	"github.com/eser/acik.io/pkg/bliss/di"
	"github.com/eser/acik.io/pkg/bliss/metricsfx"
)

func RegisterDependencies(container di.Container, config *Config, mp metricsfx.MetricsProvider) {
	service := NewGrpcService(config, mp)

	di.RegisterFor[GrpcService](container, service)
	di.Register(container, service.InnerMetrics)
}
