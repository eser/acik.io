package middlewares

import (
	"github.com/eser/acik.io/pkg/bliss/httpfx"
	"github.com/eser/acik.io/pkg/bliss/lib"
)

const CorrelationIdHeader = "X-Correlation-Id"

func CorrelationIdMiddleware() httpfx.Handler {
	return func(ctx *httpfx.Context) httpfx.Result {
		// FIXME(@eser): no need to check if the header is specified
		correlationId := ctx.Request.Header.Get(CorrelationIdHeader)
		if correlationId == "" {
			correlationId = lib.IdsGenerateUnique()
		}

		result := ctx.Next()

		ctx.ResponseWriter.Header().Set(CorrelationIdHeader, correlationId)

		return result
	}
}