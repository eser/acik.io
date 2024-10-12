package httpfx

import (
	"github.com/eser/acik.io/pkg/bliss/metricsfx"
	"go.uber.org/fx"
)

var FxModule = fx.Module( //nolint:gochecknoglobals
	"httpservice",
	fx.Provide(
		FxNew,
	),
	fx.Invoke(
		registerHooks,
	),
)

type FxResult struct {
	fx.Out

	HttpService HttpService
	Routes      Router
}

func FxNew(config *Config, mp metricsfx.MetricsProvider) FxResult {
	routes := NewRouter("/")
	httpService := NewHttpService(config, routes, mp)

	return FxResult{
		Out: fx.Out{},

		HttpService: httpService,
		Routes:      routes,
	}
}

func registerHooks(lc fx.Lifecycle, hs HttpService) {
	lc.Append(fx.Hook{
		OnStart: hs.Start,
		OnStop:  hs.Stop,
	})
}
