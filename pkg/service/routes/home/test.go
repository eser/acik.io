package home

import (
	"net/http"

	"github.com/eser/acik.io/pkg/bliss/httpfx"
	"github.com/eser/acik.io/pkg/proto/broadcast"
	"github.com/eser/acik.io/pkg/service/config"
)

func TestRoutes(routes httpfx.Router, appConfig *config.AppConfig) {
	routes.
		Route("GET /test", func(ctx *httpfx.Context) httpfx.Result {
			v := broadcast.Channel{
				Id:   "01JBREQEH27498TQRYBWA3GP81",
				Name: "eser.live",
			}

			return ctx.Results.Json(v.String())
		}).
		HasSummary("Test").
		HasDescription("This is a test endpoint.").
		HasResponse(http.StatusOK)
}
