package controller

import "github.com/gin-gonic/gin"

type LoginUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(ctx *gin.Context) {

}

func Logout(ctx *gin.Context) {

}

func Register(ctx *gin.Context) {

}
