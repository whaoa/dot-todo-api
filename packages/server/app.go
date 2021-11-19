package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type App struct {
	*gin.Engine
	Address string
	Logger  zerolog.Logger
}

func (app *App) Use(middleware ...HandlerFunc) {
	app.Engine.Use(wrapHandlers(middleware...)...)
}

func (app *App) Group(path string, middlewares ...HandlerFunc) *Group {
	group := app.Engine.Group(path, wrapHandlers(middlewares...)...)
	return createGroup(group)
}

func (app *App) Start() error {
	app.Logger.Info().Msgf("http server will start at %s", app.Address)
	return app.Run(app.Address)
}

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func Create(logger zerolog.Logger, address string) *App {
	return &App{
		Engine:  gin.New(),
		Logger:  logger,
		Address: address,
	}
}
