package middlewares

import (
	"github.com/google/uuid"
	"github.com/whaoa/dot-todo-api/packages/server"
)

type RequestIDConfig struct {
	Skipper   func(*server.Context) bool
	Generator func() string
	Handler   server.HandlerFunc
}

var DefaultRequestIDConfig = RequestIDConfig{
	Skipper:   DefaultSkipper,
	Generator: DefaultGenerator,
}

func DefaultSkipper(ctx *server.Context) bool {
	return false
}

func DefaultGenerator() string {
	return uuid.NewString()
}

func RequestId(config RequestIDConfig) server.HandlerFunc {
	return func(ctx *server.Context) {
		if !config.Skipper(ctx) {
			id := ctx.DotRequestID()
			if id == "" {
				id = config.Generator()
			}
			ctx.Header(server.HeaderXRequestID, id)
		}
		ctx.Next()
	}
}
