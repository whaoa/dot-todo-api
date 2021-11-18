package main

import (
	"github.com/whaoa/dot-todo-api/package/config"
	"github.com/whaoa/dot-todo-api/package/logger"
	"os"
)

func main() {
	conf := config.Get()
	log := logger.Create(logger.Options{
		Level:  conf.Logger.Level,
		Writer: os.Stdout,
	})
	log.Debug().Interface("config", conf).Msg("")
}
