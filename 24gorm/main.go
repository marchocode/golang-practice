package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"marcho.life/gorms/router"
)

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

func main() {

	r := router.Router()
	r.Run(":8080")
}
