package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Group struct {
	G *gin.RouterGroup
}

func (g *Group) Use(middlewares ...HandlerFunc) {
	g.G.Use(wrapHandlers(middlewares...)...)
}

func (g *Group) Group(path string, middlewares ...HandlerFunc) *Group {
	group := g.G.Group(path, wrapHandlers(middlewares...)...)
	return &Group{group}
}

func (g *Group) Handle(method, path string, middlewares ...HandlerFunc) {
	g.G.Handle(method, path, wrapHandlers(middlewares...)...)
}

func (g *Group) Match(methods []string, path string, handlers ...HandlerFunc) {
	for _, method := range methods {
		if method != "" {
			g.Handle(method, path, handlers...)
		}
	}
}

func (g *Group) Any(path string, handlers ...HandlerFunc) {
	g.G.Any(path, wrapHandlers(handlers...)...)
}

func (g *Group) GET(path string, handlers ...HandlerFunc) {
	g.G.GET(path, wrapHandlers(handlers...)...)
}

func (g *Group) POST(path string, handlers ...HandlerFunc) {
	g.G.POST(path, wrapHandlers(handlers...)...)
}

func (g *Group) DELETE(path string, handlers ...HandlerFunc) {
	g.G.DELETE(path, wrapHandlers(handlers...)...)
}

func (g *Group) PATCH(path string, handlers ...HandlerFunc) {
	g.G.PATCH(path, wrapHandlers(handlers...)...)
}

func (g *Group) PUT(path string, handlers ...HandlerFunc) {
	g.G.PUT(path, wrapHandlers(handlers...)...)
}

func (g *Group) OPTIONS(path string, handlers ...HandlerFunc) {
	g.G.OPTIONS(path, wrapHandlers(handlers...)...)
}

func (g *Group) HEAD(path string, handlers ...HandlerFunc) {
	g.G.HEAD(path, wrapHandlers(handlers...)...)
}

func (g *Group) StaticFile(path, filepath string) {
	g.G.StaticFile(path, filepath)
}

func (g *Group) Static(path, root string) {
	g.G.Static(path, root)
}

func (g *Group) StaticFS(path string, fs http.FileSystem) {
	g.G.StaticFS(path, fs)
}

func createGroup(group *gin.RouterGroup) *Group {
	return &Group{group}
}
