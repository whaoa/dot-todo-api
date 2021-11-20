package app

import (
	"github.com/rs/zerolog"
	"github.com/whaoa/dot-todo-api/internal/handler/api"
	"github.com/whaoa/dot-todo-api/packages/server"
	"github.com/whaoa/dot-todo-api/packages/server/middlewares"
)

func startHttpServer(server *server.App, logger zerolog.Logger) {
	server.Use(
		middlewares.Recovery(middlewares.RecoveryConfig{Logger: logger}),
		middlewares.RequestId(middlewares.DefaultRequestIDConfig),
		middlewares.Logger(logger),
	)

	api.Register(
		server.Group("/api"),
		logger.With().Bool("handler", true).Logger(),
	)

	logger.Fatal().Err(server.Start()).Msg("http server start error")
}
