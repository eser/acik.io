package datafx

import (
	"github.com/eser/acik.io/pkg/bliss/di"
)

func RegisterDependencies(container di.Container) {
	dp := NewDataProvider()

	di.RegisterFor[DataProvider](container, dp)
}
