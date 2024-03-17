package gee

import (
	"log"
	"net/http"
)

type routers struct {
	handlers map[string]HandlerFunc
}

func NewRouter() *routers {
	return &routers{
		handlers: map[string]HandlerFunc{},
	}
}

func (r *routers) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s -%s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *routers) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c.Writer, c.Req)
	} else {
		c.String(http.StatusNotFound, "NOT FOUND - %s", c.Path)
	}
}
