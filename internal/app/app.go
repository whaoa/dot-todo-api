package app

import (
	"github.com/rs/zerolog"
	"github.com/whaoa/dot-todo-api/packages/config"
	"github.com/whaoa/dot-todo-api/packages/server"
	"github.com/whaoa/dot-todo-api/packages/server/middlewares"
)

type Options struct {
	Config config.Config
	Logger zerolog.Logger
}

type Core struct {
	Config     config.Config
	BaseLogger zerolog.Logger
	AppLogger  zerolog.Logger
	HttpLogger zerolog.Logger
	HttpServer *server.App
}

func Boot(options Options) {
	core := createCore(options)
	core.AppLogger.Info().Interface("config", core.Config).Msg("app boot")

	startHttpServer(core)
}

func createCore(options Options) *Core {
	appLogger := options.Logger.With().Bool("app", true).Logger()
	httpLogger := options.Logger.With().Bool("http", true).Logger()
	httpServer := server.CreateApp(httpLogger, options.Config.HTTP.Address)
	return &Core{
		Config:     options.Config,
		BaseLogger: options.Logger,
		AppLogger:  appLogger,
		HttpLogger: httpLogger,
		HttpServer: httpServer,
	}
}

func startHttpServer(core *Core) {
	core.HttpServer.Use(
		middlewares.RequestId(middlewares.DefaultRequestIDConfig),
		middlewares.Logger(core.HttpLogger),
	)

	core.HttpLogger.Fatal().
		Err(core.HttpServer.Start()).
		Msg("http server start error")
}
