package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/whaoa/dot-todo-api/packages/server"
	"strconv"
	"time"
)

func Logger(logger zerolog.Logger) server.HandlerFunc {
	return func(ctx *server.Context) {
		start := time.Now()

		logger.Info().
			Bool("route", true).
			Str("direct", "in").
			Str("id", ctx.DotRequestID()).
			Str("method", ctx.Request.Method).
			Str("path", ctx.Request.URL.Path).
			Str("remote-ip", ctx.ClientIP()).
			Msg(ctx.Request.URL.RawQuery)

		ctx.Next()

		end := time.Now()

		logger.Info().
			Bool("route", true).
			Str("direct", "out").
			Str("id", ctx.DotRequestID()).
			Int("state", ctx.Writer.Status()).
			Str("usage", end.Sub(start).String()).
			Str("length", strconv.FormatInt(int64(ctx.Writer.Size()), 10)).
			Msg(ctx.Errors.ByType(gin.ErrorTypePrivate).String())
	}
}
