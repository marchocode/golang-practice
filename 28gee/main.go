package main

import (
	"net/http"

	"marcho.life/gee/gee"
)

func main() {

	engine := gee.New()

	engine.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	engine.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})

	engine.Run()
}
