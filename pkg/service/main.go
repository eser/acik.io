package service

import (
	"github.com/eser/acik.io/pkg/bliss"
	"github.com/eser/acik.io/pkg/bliss/logfx"
	"go.uber.org/fx"
)

func Run() {
	bliss.Load()

	app := fx.New(
		fx.WithLogger(logfx.GetFxLogger),
		bliss.FxModule,
		FxModule,
	)

	app.Run()
}
