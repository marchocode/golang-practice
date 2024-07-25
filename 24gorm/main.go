package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"createAt"`
	UpdatedAt time.Time      `json:"updateAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type User struct {
	gorm.Model
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func getDb() *gorm.DB {

	db, err := gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}

func update(ctx *gin.Context) {

	var u User
	err := ctx.BindJSON(&u)

	if err != nil {
		log.Println(err)
	}

	db := getDb()

	db.Save(&u)
}

func delete(ctx *gin.Context) {

	id := ctx.Params.ByName("id")
	db := getDb()
	db.Delete(&User{}, id)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})

}

func findUser(ctx *gin.Context) {

	var u User
	id := ctx.Params.ByName("id")

	db := getDb()
	tx := db.First(&u, id)

	if tx.Error != nil {
		log.Println(tx.Error)
	}

	ctx.JSON(http.StatusOK, u)
}

func addUser(ctx *gin.Context) {

	var u User
	ctx.BindJSON(&u)

	db := getDb()
	db.Create(&u)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func test1() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("test1 start")
		c.Next()
		log.Println("test1 emd")
	}
}

func test2() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("test2 start")
		c.Next()
		log.Println("test2 emd")
	}
}

func readPublicKey() (*rsa.PublicKey, error) {

	publicKeyData, err := os.ReadFile("pub.pem")

	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(publicKeyData)

	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing public key")
	}

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)

	if err != nil {
		return nil, err
	}

	switch pub := publicKey.(type) {
	case *rsa.PublicKey:
		return pub, nil
	default:
		return nil, fmt.Errorf("not an RSA public key")
	}

}

func readPrivateKey() (*rsa.PrivateKey, error) {

	privateKeyData, err := os.ReadFile("privkey.pem")

	if err != nil {
		log.Fatalln(err)
	}

	privateKeyBlock, _ := pem.Decode(privateKeyData)

	if privateKeyBlock == nil {
		return nil, errors.New("私钥文件格式错误")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("无法解析私钥")
	}

	return privateKey, nil
}

func verifyToken(token string) (bool, error) {

	pub, err := readPublicKey()

	if err != nil {
		return false, nil
	}

	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return pub, nil
	})

	if err != nil {
		return false, nil
	}

	return t.Valid, nil

}

func genToken() string {

	key, err := readPrivateKey()

	if err != nil {
		log.Fatalln(err)
	}

	t := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.MapClaims{
		"id": "123",
	})

	s, err := t.SignedString(key)

	if err != nil {
		fmt.Println(err)
	}

	return s
}

func login(ctx *gin.Context) {

}

func router() *gin.Engine {

	r := gin.New()
	r.Use(test1(), test2())

	r.POST("/login", login)

	user := r.Group("/user")
	{
		user.POST("/", addUser)
		user.PUT("/", update)
		user.GET("/:id", findUser)
		user.DELETE("/:id", delete)
	}

	return r
}

func main() {

	// db := getDb()
	// db.AutoMigrate(&User{})

	// r := router()

	// r.Run(":8080")

	s := genToken()
	fmt.Println(s)

	b, err := verifyToken(s)
	fmt.Println(b)
	fmt.Println(err)
}
