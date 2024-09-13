package main

import (
	"net/http"

	"marcho.life/gee/gee"
)

func main() {

	engine := gee.New()

	engine.Get("/ping", func(ctx *gee.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	engine.Get("/hello/:name", func(ctx *gee.Context) {

		name := ctx.Param("name")

		ctx.Json(http.StatusOK, gee.H{
			"username": name,
			"password": "123456",
		})
	})

	engine.Run()
}
