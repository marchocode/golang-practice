package cache

import (
	"testing"
)

func TestGet(t *testing.T) {

	cache := NewRedisClient()
	t.Log(cache.TTL("a").Seconds())

}
