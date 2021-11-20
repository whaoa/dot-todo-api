package ping

import (
	"github.com/whaoa/dot-todo-api/internal/response"
	"github.com/whaoa/dot-todo-api/packages/server"
)

func Handler(ctx *server.Context) {
	ctx.JSON(
		response.OK.Status(),
		response.Create(response.OK, ctx.Request.Method),
	)
}
