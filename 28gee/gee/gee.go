package gee

import (
	"log"
	"net/http"
	"path"
	"strings"
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

func (group *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	absolutePath := path.Join(group.prefix, relativePath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))

	return func(c *Context) {
		file := c.Param("filepath")
		// Check if file exists and/or if we have permission to access it
		if _, err := fs.Open(file); err != nil {
			c.Status(http.StatusNotFound)
			return
		}

		fileServer.ServeHTTP(c.Write, c.Req)
	}
}

// 将根目录文件映射到指定路由下
func (group *RouterGroup) Static(relativePath string, root string) {
	handler := group.createStaticHandler(relativePath, http.Dir(root))
	urlPattern := path.Join(relativePath, "/*filepath")
	// Register GET handlers
	group.Get(urlPattern, handler)
}

func (c *Engine) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	log.Printf("http request path=%s\n", req.URL.Path)
	ctx := newContext(res, req)

	// 给 ctx 绑定中间件
	middlewares := make([]HandlerFunc, 0)

	for _, group := range c.groups {

		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}

	ctx.middlewares = middlewares
	c.router.handler(ctx)
}

func (c *Engine) Use(middle ...HandlerFunc) {
	c.middlewares = append(c.middlewares, middle...)
}

func (c *Engine) Run(addr string) {
	http.ListenAndServe(addr, c)
}
