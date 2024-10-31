package service

import (
	"github.com/eser/acik.io/pkg/bliss/di"
)

func Run() {
	err := Startup(di.Default)
	if err != nil {
		panic(err)
	}

	// app := fx.New(
	// 	fx.WithLogger(logfx.GetFxLogger),
	// 	bliss.FxModule,
	// 	FxModule,
	// )

	// app.Run()
}
