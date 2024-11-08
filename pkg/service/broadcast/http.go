package broadcast

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/eser/acik.io/pkg/bliss/httpfx"
	"github.com/eser/acik.io/pkg/bliss/httpfx/middlewares"
	"github.com/eser/acik.io/pkg/proto/broadcast"
	"github.com/eser/acik.io/pkg/service"
)

func RegisterHttpRoutes(routes httpfx.Router, appConfig *service.AppConfig, logger *slog.Logger) {
	routes.
		Route("GET /protected", middlewares.AuthMiddleware(), func(ctx *httpfx.Context) httpfx.Result {
			// message := fmt.Sprintf("Hello from %s! this endpoint is protected!", appConfig.AppName)

			// return ctx.Results.PlainText(message)
			v := broadcast.Channel{
				Id:   "01JBREQEH27498TQRYBWA3GP81",
				Name: "eser.live",
			}

			return ctx.Results.Json(v.String())
		}).
		HasSummary("Protected page").
		HasDescription("A page protected with JWT auth.").
		HasResponse(http.StatusOK)

	routes.
		Route("POST /send", func(ctx *httpfx.Context) httpfx.Result {
			body, err := io.ReadAll(ctx.Request.Body)
			if err != nil {
				return ctx.Results.Error(http.StatusInternalServerError, err.Error())
			}

			var payload broadcast.SendRequest
			err = json.Unmarshal(body, &payload)
			if err != nil {
				return ctx.Results.Error(http.StatusBadRequest, err.Error())
			}

			logger.Info(
				"Send",
				slog.String("channelId", payload.GetChannelId()),
				slog.Any("message", payload.GetMessage()),
			)

			return ctx.Results.Ok()
		}).
		HasSummary("Send a message to a channel").
		HasDescription("Send a message to a channel.").
		HasResponse(http.StatusOK)
}
