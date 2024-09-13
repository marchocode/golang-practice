package gee

import (
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(ctx *Context) {
		t := time.Now()
		ctx.Next()
		log.Printf("[Logger] %s [%v] \n", ctx.Path, time.Since(t))
	}
}
