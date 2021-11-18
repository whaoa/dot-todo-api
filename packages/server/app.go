package server

import "github.com/gin-gonic/gin"

type App struct {
	*gin.Engine
}

func (app *App) Use(middleware ...HandlerFunc) {
	app.Engine.Use(wrapHandlers(middleware...)...)
}

func (app *App) Group(path string, middlewares ...HandlerFunc) *Group {
	group := app.Engine.Group(path, wrapHandlers(middlewares...)...)
	return createGroup(group)
}

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func Create() *App {
	app := gin.New()
	return &App{app}
}

func CreateApp() *App {
	app := Create()
	app.Engine.Use(gin.Recovery())
	return app
}
