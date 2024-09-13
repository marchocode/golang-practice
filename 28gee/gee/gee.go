package gee

import (
	"log"
	"net/http"
)

type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup
}

type HandlerFunc func(*Context)

// 路由分组
type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	parent      *RouterGroup
	engine      *Engine
}

// group create a new RouterGroup
func (g *RouterGroup) Group(prefix string) *RouterGroup {

	engine := g.engine

	newGroup := &RouterGroup{
		prefix: g.prefix + prefix,
		parent: g,
		engine: engine,
	}

	engine.groups = append(engine.groups, newGroup)

	return newGroup
}

func New() *Engine {
	engine := &Engine{router: newRouter()}

	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}

	return engine
}

func (g *RouterGroup) addRouter(method string, path string, handler HandlerFunc) {

	pattern := g.prefix + path
	g.engine.router.addRouter(method, pattern, handler)
}

func (g *RouterGroup) Get(url string, f HandlerFunc) {
	g.addRouter("GET", url, f)
}

func (g *RouterGroup) Post(url string, f HandlerFunc) {
	g.addRouter("POST", url, f)
}

func (c *Engine) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	log.Printf("http request path=%s\n", req.URL.Path)
	ctx := newContext(res, req)
	c.router.handler(ctx)
}

func (c *Engine) Run(addr string) {
	http.ListenAndServe(addr, c)
}
