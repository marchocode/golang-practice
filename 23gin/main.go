package main

import (
	"database/sql"
	"fmt"
	"log"
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

func initDb() {

	db, err := sql.Open("sqlite3", "db.sqlite")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	create := `
		CREATE TABLE users(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username varchar(32) not null,
			password varchar(32) not null
	)
	`
	_, err = db.Exec(create)

	if err != nil {
		log.Fatalln(err)
	}

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

	initDb()

	r := setupRouter()
	r.Run(":8080")
}
