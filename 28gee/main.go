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

	engine.Get("/hello", func(ctx *gee.Context) {

		name := ctx.Query("name")

		ctx.Json(http.StatusOK, gee.H{
			"username": name,
			"password": "123456",
		})
	})

	engine.Run()
}
