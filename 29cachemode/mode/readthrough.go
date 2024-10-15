package mode

import (
	"marcho.life/cachemode/cache"
	"marcho.life/cachemode/db"
)

type ReadThrough struct {
	cache    cache.Cache
	database db.Database
}

// 读穿透，在缓存没有数据的时候，缓存主动帮我们读取
func (r *ReadThrough) Get(key string) string {

	if v := r.cache.Get(key); v != "" {
		return v
	}

	// 读取并回写到缓存
	v := r.database.Select(key)

	// 这里可以异步进行
	r.cache.Set(key, v)

	return v
}

// 写还是和CacheAside一样，没有区别
func (r *ReadThrough) Set(key, val string) {

}
