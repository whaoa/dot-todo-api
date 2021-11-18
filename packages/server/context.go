package server

import "github.com/gin-gonic/gin"

type Context struct {
	*gin.Context
}

func (ctx *Context) DotRequestID() string {
	id := ctx.GetHeader(HeaderXRequestID)
	if id == "" {
		id = ctx.Writer.Header().Get(HeaderXRequestID)
	}
	return id
}

func createContext(context *gin.Context) *Context {
	return &Context{context}
}
