package handler

import (
	"log"

	"github.com/gin-gonic/gin"
)

var logger = log.Default()

func Logger() gin.HandlerFunc {

	return func(c *gin.Context) {

		logger.Println("Logger in")
		c.Next()
		logger.Println("Logger end")
	}
}

func WhitePaper() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("test2 start")
		c.Next()
		log.Println("test2 emd")
	}
}

func TokenVerify() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("test2 start")
		c.Next()
		log.Println("test2 emd")
	}
}
