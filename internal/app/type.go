package app

import (
	"github.com/rs/zerolog"
	"github.com/whaoa/dot-todo-api/packages/config"
	"github.com/whaoa/dot-todo-api/packages/server"
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

func createCore(options Options) *Core {
	appLogger := options.Logger.With().Bool("app", true).Logger()
	httpLogger := options.Logger.With().Bool("http", true).Logger()
	httpServer := server.Create(httpLogger, options.Config.HTTP.Address)

	return &Core{
		Config:     options.Config,
		BaseLogger: options.Logger,
		AppLogger:  appLogger,
		HttpLogger: httpLogger,
		HttpServer: httpServer,
	}
}
