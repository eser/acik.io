package middlewares

import "github.com/eser/acik.io/pkg/bliss/httpfx"

func ErrorHandlerMiddleware() httpfx.Handler {
	return func(ctx *httpfx.Context) httpfx.Result {
		result := ctx.Next()

		return result
	}
}