package gee

import (
	"net/http"
)

// 路由处理
type HandlerFunc func(c *Context)

type engine struct {
	router *router
}

func New() *engine {
	return &engine{router: newRouter()}
}

func (e *engine) addRoute(method string, pattern string, handler HandlerFunc) {
	e.router.addRoute(method, pattern, handler)
}

func (e *engine) GET(pattern string, handler HandlerFunc) {
	e.addRoute("GET", pattern, handler)
}

func (e *engine) POST(pattern string, handler HandlerFunc) {
	e.addRoute("POST", pattern, handler)
}

func (e *engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

func (e *engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// add router
	c := newContext(w, r)
	e.router.handle(c)
}
