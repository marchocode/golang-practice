package gee

import (
	"fmt"
	"log"
)

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		handlers: make(map[string]HandlerFunc),
	}
}

func (c *router) handler(ctx *Context) {

	key := fmt.Sprintf("%s-%s", ctx.Method, ctx.Req.URL.Path)

	if handler, ok := c.handlers[key]; ok {
		handler(ctx)
	} else {

	}

}

func (c *router) addRouter(method string, pattern string, handler HandlerFunc) {

	log.Printf("addRouter %s - %s", method, pattern)

	key := fmt.Sprintf("%s-%s", method, pattern)
	c.handlers[key] = handler
}
