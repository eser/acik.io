package httpfx

import (
	"github.com/eser/acik.io/pkg/bliss/di"
	"github.com/eser/acik.io/pkg/bliss/metricsfx"
)

// var FxModule = fx.Module( //nolint:gochecknoglobals
// 	"httpservice",
// 	fx.Provide(
// 		FxNew,
// 	),
// 	fx.Invoke(
// 		registerHooks,
// 	),
// )

// func registerHooks(lc fx.Lifecycle, hs HttpService) {
// 	lc.Append(fx.Hook{
// 		OnStart: hs.Start,
// 		OnStop:  hs.Stop,
// 	})
// }

func Startup(container di.Container, config *Config, mp metricsfx.MetricsProvider) {
	routes := NewRouter("/")
	httpService := NewHttpService(config, routes, mp)

	di.RegisterFor[Router](container, routes)
	di.RegisterFor[HttpService](container, httpService)
}
