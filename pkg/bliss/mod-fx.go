package bliss

import (
	"github.com/eser/acik.io/pkg/bliss/configfx"
	"github.com/eser/acik.io/pkg/bliss/datafx"
	"github.com/eser/acik.io/pkg/bliss/httpfx"
	"github.com/eser/acik.io/pkg/bliss/logfx"
	"github.com/eser/acik.io/pkg/bliss/metricsfx"
	"go.uber.org/fx"
)

var FxModule = fx.Module( //nolint:gochecknoglobals
	"bliss",
	logfx.FxModule,
	configfx.FxModule,
	metricsfx.FxModule,
	httpfx.FxModule,
	datafx.FxModule,
)
