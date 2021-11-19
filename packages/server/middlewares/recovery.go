package middlewares

import (
	"github.com/rs/zerolog"
	"github.com/whaoa/dot-todo-api/packages/server"
	"net"
	"net/http"
	"os"
	"strings"
)

type RecoveryConfig struct {
	Logger  zerolog.Logger
	Handler func(*server.Context, error)
}

func DefaultRecoveryHandler(ctx *server.Context, err error) {
	ctx.AbortWithStatus(http.StatusInternalServerError)
}

func Recovery(option RecoveryConfig) server.HandlerFunc {
	if option.Handler == nil {
		option.Handler = DefaultRecoveryHandler
	}

	return func(ctx *server.Context) {
		defer func() {
			err := recover()
			if err == nil {
				return
			}

			option.Logger.Error().Err(err.(error)).Msg("panic recovered")

			// 检查连接是否断开
			var brokenPipe bool
			if ne, ok := err.(*net.OpError); ok {
				if se, ok := ne.Err.(*os.SyscallError); ok {
					errMsg := strings.ToLower(se.Error())
					isBroken := strings.Contains(errMsg, "broken pipe")
					isRest := strings.Contains(errMsg, "connection reset by peer")
					if isBroken || isRest {
						brokenPipe = true
					}
				}
			}

			// 如果连接已断开，则无法向其写入响应状态
			if brokenPipe {
				_ = ctx.Error(err.(error))
				ctx.Abort()
			} else {
				option.Handler(ctx, err.(error))
			}
		}()
		ctx.Next()
	}
}
