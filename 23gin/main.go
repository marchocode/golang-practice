package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	os.Remove("db.sqlite")
}

var db map[string]string = map[string]string{}

type User struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func openDb() {
	db, err := sql.Open("sqlite3", "sqlite.db")
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	auth := r.Group("/auth")

	{
		auth.POST("/register", func(ctx *gin.Context) {

			var user User
			ctx.BindJSON(&user)
			fmt.Println(user)

			db[user.Username] = user.Password

			ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
		})

		auth.POST("/login", func(ctx *gin.Context) {

		})

	}

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
