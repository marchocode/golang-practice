package gee

import (
	"testing"
)

func AddRouter(t *testing.T) *router {

	r := newRouter()

	r.addRouter("GET", "/", nil)
	r.addRouter("GET", "/hello/:name", nil)
	r.addRouter("GET", "/hello/b/c", nil)
	r.addRouter("GET", "/hi/:name", nil)
	r.addRouter("GET", "/assets/*filepath", nil)

	return r
}

func TestParseSearch(t *testing.T) {

	r := AddRouter(t)

	re := r.parsePattern("/hello/:name")
	t.Log(re)
}

func TestParsePattern(t *testing.T) {

	r := AddRouter(t)

	re := r.parsePattern("/hello/:name")
	t.Log(re)
}

func TestGetRouter(t *testing.T) {

	r := AddRouter(t)

	node, m := r.getRouter("GET", "/assets/css/jquery.js")

	t.Log(node)
	t.Log(m)
}
