package main

import (
	"context"
	"errors"
	"log"

	"github.com/redis/go-redis/v9"
)

type Student struct {
	Id      int    `redis:"id"`
	Name    string `redis:"name"`
	Age     int    `redis:"age"`
	Address string `redis:"address"`
}

func getConnect() *redis.Client {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})

	return rdb
}

func main() {

	ctx := context.TODO()
	client := getConnect()

	defer client.Close()

	err := client.Set(ctx, "goredis", "v8", 0).Err()
	log.Printf("set a value\n")

	if err != nil {
		log.Fatalln(err)
	}

	re, err := client.Get(ctx, "goredis").Result()

	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("get a value = %s\n", re)

	// no value.
	err = client.Get(ctx, "demo").Err()

	if errors.Is(err, redis.Nil) {
		log.Println("key is not yet.")
	}

	// store a struct
	stu := Student{
		Id:      1,
		Name:    "mrc",
		Age:     15,
		Address: "usa",
	}

	err = client.HSet(ctx, "stu:1", stu).Err()

	if err != nil {
		panic(err)
	}

}
