package api

import (
	"github.com/rs/zerolog"
	"github.com/whaoa/dot-todo-api/internal/handler/api/ping"
	"github.com/whaoa/dot-todo-api/packages/server"
)

func Register(route *server.Group, logger zerolog.Logger) {
	// ping
	route.Any("/ping", ping.Handler)
}
