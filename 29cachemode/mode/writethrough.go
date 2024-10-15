package mode

import (
	"marcho.life/cachemode/cache"
	"marcho.life/cachemode/db"
)

type WriteThrough struct {
	database db.Database
	cache    cache.Cache
}

func (w *WriteThrough) Get(key string) string {
	return ""
}

// 写穿透，只需要向缓存写入数据，其他的无需关心
func (w *WriteThrough) Set(key, val string) {

	w.database.Update(key, val)
	// 回写
	// 这一步可以异步进行
	w.cache.Set(key, val)
}
