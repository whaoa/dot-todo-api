package main

import (
	"github.com/whaoa/dot-todo-api/internal/app"
	"github.com/whaoa/dot-todo-api/packages/config"
	"github.com/whaoa/dot-todo-api/packages/logger"
	"os"
)

func main() {
	conf := config.Get()
	log := logger.Create(logger.Options{
		Level:  conf.Logger.Level,
		Writer: os.Stdout,
	})
	app.Boot(app.Options{Config: conf, Logger: log})
}
