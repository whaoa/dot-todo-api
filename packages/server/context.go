package server

import "github.com/gin-gonic/gin"

type Context struct {
	*gin.Context
}

func createContext(context *gin.Context) *Context {
	return &Context{context}
}
