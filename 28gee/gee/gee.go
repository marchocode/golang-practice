package gee

import (
	"log"
	"net/http"
)

type Engine struct {
	router *router
}

type HandlerFunc func(*Context)

func New() *Engine {
	return &Engine{
		router: newRouter(),
	}
}

func (c *Engine) addRouter(method string, pattern string, handler HandlerFunc) {
	c.router.addRouter(method, pattern, handler)
}

func (c *Engine) Get(url string, f HandlerFunc) {
	c.addRouter("GET", url, f)
}

func (c *Engine) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	log.Printf("http request path=%s\n", req.URL.Path)
	ctx := newContext(res, req)
	c.router.handler(ctx)
}

func (c *Engine) Run() {
	http.ListenAndServe(":9090", c)
}
