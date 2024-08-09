package router

import (
	"github.com/gin-gonic/gin"
	"marcho.life/gorms/controller"
	"marcho.life/gorms/handler"
)

func Router() *gin.Engine {

	r := gin.New()
	r.Use(handler.Logger())

	// auth
	auth := r.Group("/auth")
	{
		auth.POST("/login", controller.Login)
		auth.PUT("/logout", controller.Logout)
		auth.POST("/register", controller.Register)
	}

	return r
}
