package healthcheck

import (
	"net/http"

	"github.com/eser/acik.io/pkg/bliss/httpfx"
)

// var FxModule = fx.Module( //nolint:gochecknoglobals
// 	"healthcheck",
// 	fx.Invoke(
// 		RegisterRoutes,
// 	),
// )

func RegisterRoutes(routes httpfx.Router) error {
	routes.
		Route("GET /health-check", func(ctx *httpfx.Context) httpfx.Result {
			return ctx.Results.Ok()
		}).
		HasSummary("Health Check").
		HasDescription("Health Check Endpoint").
		HasResponse(http.StatusNoContent)

	return nil
}
