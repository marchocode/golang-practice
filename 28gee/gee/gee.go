package gee

import (
	"fmt"
	"net/http"
)

type Engine struct {
	router map[string]http.HandlerFunc
}

func New() *Engine {
	return &Engine{
		router: make(map[string]http.HandlerFunc),
	}
}

func (c *Engine) Get(url string, handler http.HandlerFunc) {
	c.addRouter("GET", url, handler)
}

func (c *Engine) addRouter(method string, pattern string, handler http.HandlerFunc) {
	key := fmt.Sprintf("%s-%s", method, pattern)
	c.router[key] = handler
}

func (c *Engine) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	key := fmt.Sprintf("%s-%s", req.Method, req.RequestURI)

	if handler, ok := c.router[key]; ok {
		handler(res, req)
	} else {
		res.Write([]byte("404 not fouund"))
	}

}

func (c *Engine) Run() {
	http.ListenAndServe(":9090", c)
}
