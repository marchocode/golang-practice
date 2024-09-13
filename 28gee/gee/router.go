package gee

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

// 切割pattern
func (c *router) parsePattern(pattern string) []string {

	parts := make([]string, 0)

	for _, item := range strings.Split(pattern, "/") {

		if item == "" {
			continue
		}

		parts = append(parts, item)

		if item[0] == '*' {
			break
		}

	}

	return parts
}

func (c *router) handler(ctx *Context) {

	n, params := c.getRouter(ctx.Method, ctx.Path)

	if n == nil {
		ctx.String(http.StatusNotFound, "404 not found \n")
	}

	// 参数赋值
	ctx.Params = params
	key := fmt.Sprintf("%s-%s", ctx.Method, n.pattern)

	ctx.middlewares = append(ctx.middlewares, c.handlers[key])
	ctx.Next()
}

func (c *router) addRouter(method string, pattern string, handler HandlerFunc) {

	log.Printf("addRouter %s - %s", method, pattern)
	key := fmt.Sprintf("%s-%s", method, pattern)

	parts := c.parsePattern(pattern)

	// 空节点初始化
	if _, ok := c.roots[method]; !ok {
		c.roots[method] = &node{}
	}

	// 将当前路由初始化到前缀树中
	c.roots[method].insert(pattern, parts, 0)

	c.handlers[key] = handler
}

// 解析路由
func (c *router) getRouter(method string, path string) (*node, map[string]string) {

	searchParts := c.parsePattern(path)
	params := make(map[string]string)

	// 按照method 获得根节点
	root, ok := c.roots[method]

	if !ok {
		return nil, nil
	}

	// 寻找适配的路由
	n := root.search(searchParts, 0)

	if n == nil {
		return nil, nil
	}

	// 解析原有路由，进行替换
	parts := c.parsePattern(n.pattern)

	// 替换router里面的变量
	for index, part := range parts {

		// 替换变量
		if part[0] == ':' {
			params[part[1:]] = searchParts[index]
		}

		// 替换通配符
		if part[0] == '*' && len(part) > 1 {
			params[part[1:]] = strings.Join(searchParts[index:], "/")
			break
		}
	}

	return n, params
}
