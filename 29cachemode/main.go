package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"marcho.life/cachemode/mode"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	f := mode.NewRefreshAhead()

	// read
	r.GET("/refresh_a_head", func(c *gin.Context) {

		k := c.Query("k")
		val := f.Read(k)

		log.Printf("refresh_a_head read a key=%s val=%s", k, val)

		c.JSON(200, gin.H{
			k: val,
		})

	})

	// write
	r.POST("/refresh_a_head", func(c *gin.Context) {

		k := c.Query("k")
		val := c.Query("val")

		f.Write(k, val)

		c.JSON(200, gin.H{
			k: val,
		})
	})

	r.Run("0.0.0.0:9090")
}
