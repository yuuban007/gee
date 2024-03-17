package gee

import (
	"net/http"
)

// HandlerFunc
type HandlerFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	router routers
}

// New is constructor of gee.Engine
func New() *Engine {
	return &Engine{
		router: routers{
			handlers: map[string]HandlerFunc{},
		},
	}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

// GET define the method to add GET request
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST define the method to add POST request
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// RUN
func (engine *Engine) RUN(addr string) error {
	return http.ListenAndServe(addr, engine)
}

// ServeHTTP implements http.Handler.
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := NewContext(w, req)
	engine.router.handle(c)
}
