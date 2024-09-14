package main

import (
	"net/http"

	"marcho.life/gee/gee"
)

func main() {

	engine := gee.New()
	engine.Use(gee.Logger())
	engine.Use(gee.Recovery())
	

	engine.Static("/assets", "./static")

	engine.Get("/panic", func(ctx *gee.Context) {
		arr := []string{"1", "2", "3"}
		ctx.String(http.StatusOK, arr[4])
	})

	v1 := engine.Group("/v1")

	{
		v1.Get("/ping", func(ctx *gee.Context) {
			ctx.String(http.StatusOK, "v1 pong")
		})
	}

	v2 := engine.Group("/v2")

	{
		v2.Get("/hello/:name", func(ctx *gee.Context) {

			name := ctx.Param("name")

			ctx.Json(http.StatusOK, gee.H{
				"username": name,
				"password": "123456",
			})
		})
	}

	engine.Run(":9090")
}
