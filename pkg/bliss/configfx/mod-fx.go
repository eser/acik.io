package configfx

import (
	"github.com/eser/acik.io/pkg/bliss/di"
)

// var FxModule = fx.Module( //nolint:gochecknoglobals
// 	"config",
// 	fx.Provide(
// 		FxNew,
// 	),
// )

func Startup(container di.Container) {
	cl := NewConfigLoader()

	di.RegisterFor[ConfigLoader](container, cl)
}
