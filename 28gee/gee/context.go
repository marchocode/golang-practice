package gee

import (
	"encoding/json"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	Write http.ResponseWriter
	Req   *http.Request

	Method     string
	RequestURI string

	Code int
}

const (
	ContentType = "Content-Type"
)

func newContext(res http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Write:      res,
		Req:        req,
		Method:     req.Method,
		RequestURI: req.RequestURI,
	}
}

func (c *Context) Status(code int) {
	c.Code = code
	c.Write.WriteHeader(code)
}

func (c *Context) SetHeader(key string, value string) {
	c.Write.Header().Add(key, value)
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) Json(code int, data interface{}) {
	c.SetHeader(ContentType, "application/json")
	c.Status(code)
	
	encoder := json.NewEncoder(c.Write)

	if err := encoder.Encode(data); err != nil {
		http.Error(c.Write,err.Error(),500)
	}

}

func (c *Context) String(code int, data string) {
	c.SetHeader(ContentType, "text/plain")
	c.Status(code)
	c.Write.Write([]byte(data))
}
