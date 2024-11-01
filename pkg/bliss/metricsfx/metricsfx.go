package metricsfx

import "github.com/eser/acik.io/pkg/bliss/di"

func RegisterDependencies(container di.Container) {
	mp := NewMetricsProvider()

	di.RegisterFor[MetricsProvider](container, mp)
}
