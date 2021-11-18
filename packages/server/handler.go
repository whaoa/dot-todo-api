package server

import "github.com/gin-gonic/gin"

type HandlerFunc func(*Context)

func wrapHandler(handler HandlerFunc) gin.HandlerFunc {
	return func(context *gin.Context) {
		handler(createContext(context))
	}
}

func wrapHandlers(handlers ...HandlerFunc) (result []gin.HandlerFunc) {
	for _, handler := range handlers {
		result = append(result, wrapHandler(handler))
	}
	return
}
