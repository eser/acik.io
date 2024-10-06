package service

import (
	"github.com/eser/acik.io/pkg/bliss"
	"github.com/eser/acik.io/pkg/bliss/logfx"
	"go.uber.org/fx"
)

func Run() {
	app := fx.New(
		fx.WithLogger(logfx.GetFxLogger),
		bliss.FxModule,
		FxModule,
	)

	app.Run()
}
