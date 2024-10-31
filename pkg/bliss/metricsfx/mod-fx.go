package metricsfx

import "github.com/eser/acik.io/pkg/bliss/di"

// var FxModule = fx.Module( //nolint:gochecknoglobals
// 	"metrics",
// 	fx.Provide(
// 		FxNew,
// 	),
// )

func Startup(container di.Container) {
	mp := NewMetricsProvider()

	di.RegisterFor[MetricsProvider](container, mp)
}
