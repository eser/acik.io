package home

import (
	"net/http"

	"github.com/eser/acik.io/pkg/bliss/httpfx"
	"github.com/eser/acik.io/pkg/proto/vehicle"
	"github.com/eser/acik.io/pkg/service/config"
)

func TestRoutes(routes httpfx.Router, appConfig *config.AppConfig) {
	routes.
		Route("GET /test", func(ctx *httpfx.Context) httpfx.Result {
			v := vehicle.Vehicle{
				Id:    1,
				Name:  "Car",
				Brand: "Toyota",
				Type:  vehicle.VehicleType_SEDAN,
			}

			return ctx.Results.Json(v.String())
		}).
		HasSummary("Test").
		HasDescription("This is a test endpoint.").
		HasResponse(http.StatusOK)
}
