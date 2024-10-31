package home

import (
	"fmt"
	"net/http"

	"github.com/eser/acik.io/pkg/bliss/httpfx"
	"github.com/eser/acik.io/pkg/bliss/httpfx/middlewares"
	"github.com/eser/acik.io/pkg/service/config"
)

func RegisterIndexRoute(routes httpfx.Router, appConfig *config.AppConfig) error {
	routes.
		Route("GET /", func(ctx *httpfx.Context) httpfx.Result {
			message := fmt.Sprintf(
				"Hello %s (%s) from %s!",
				ctx.Request.Context().Value(middlewares.ClientAddr),
				ctx.Request.Context().Value(middlewares.ClientAddrOrigin),
				appConfig.AppName,
			)

			return ctx.Results.PlainText(message)
		}).
		HasSummary("Homepage").
		HasDescription("This is the homepage of the service.").
		HasResponse(http.StatusOK)

	return nil
}
