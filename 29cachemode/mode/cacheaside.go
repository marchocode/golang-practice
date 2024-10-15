package mode

import (
	"marcho.life/cachemode/cache"
	"marcho.life/cachemode/db"
)

type CacheASide struct {
	database db.Database
	cache    cache.Cache
}

func NewCacheAside() *CacheASide {
	return &CacheASide{
		database: db.NewMysql(),
		cache:    cache.NewRedisClient(),
	}
}

// 优先从缓存加载，若没有，再去数据库查询
func (c *CacheASide) Read(key string) string {

	// 从缓存加载
	if v := c.cache.Get(key); v != "" {
		return v
	}

	v := c.database.Select(key)

	// 写回缓存
	c.cache.Set(key, v)

	return v
}

func (c *CacheASide) Write(key, val string) {

	// 写回数据库
	c.database.Update(key, val)

	// 写回缓存
	c.cache.Set(key, val)
}
