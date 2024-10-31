package datafx

import (
	"github.com/eser/acik.io/pkg/bliss/di"
)

// var FxModule = fx.Module( //nolint:gochecknoglobals
// 	"data",
// 	fx.Provide(
// 		FxNew,
// 	),
// )

// TODO(@eser) multiple db support.
func Startup(container di.Container) {
	dp := NewDataProvider()

	di.RegisterFor[DataProvider](container, dp)
}
