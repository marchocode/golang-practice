package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	ctx context.Context
	rdb *redis.Client
}

func NewRedisClient() *RedisClient {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return &RedisClient{
		rdb: rdb,
		ctx: context.Background(),
	}

}

func (r *RedisClient) Get(key string) string {
	return r.rdb.Get(r.ctx, key).Val()
}

func (r *RedisClient) Set(key, val string) {
	r.rdb.Set(r.ctx, key, val, -1)
}

func (r *RedisClient) SetWithTime(key, val string, time time.Duration) {
	r.rdb.Set(r.ctx, key, val, time)
}

func (r *RedisClient) FlushAll() {
	r.rdb.FlushAll(r.ctx)
}

func (r *RedisClient) Keys() []string {
	return r.rdb.Keys(r.ctx, "*").Val()
}

func (r *RedisClient) TTL(key string) time.Duration {
	return r.rdb.TTL(r.ctx, key).Val()
}

func (r *RedisClient) Show() {

	keys := r.rdb.Keys(r.ctx, "*").Val()

	for _, key := range keys {
		fmt.Printf("%s -> %s \n", key, r.rdb.Get(r.ctx, key).Val())
	}

}
