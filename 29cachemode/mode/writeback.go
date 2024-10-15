package mode

import (
	"time"

	"marcho.life/cachemode/cache"
	"marcho.life/cachemode/db"
)

type WriteBack struct {
	cache    cache.Cache
	database db.Database
}

func NewWriteBack() *WriteBack {
	// 开启一个 go routine 持续监听

	wb := &WriteBack{
		cache:    cache.NewRedisClient(),
		database: db.NewMysql(),
	}

	// listening...
	go wb.listen()

	return wb
}

func (w *WriteBack) listen() {

	for {

		for _, key := range w.cache.Keys() {

			lastTime := w.cache.TTL(key)

			// 快过期
			if lastTime.Seconds() <= 3 {

				val := w.cache.Get(key)
				w.database.Update(key, val)
			}

		}

		time.Sleep(1 * time.Second)
	}

}

// 直接从缓存读取
func (w *WriteBack) Read(key string) string {
	return w.cache.Get(key)
}

// 直接写入缓存
// 存储10秒
func (w *WriteBack) Write(key, val string) {
	w.cache.SetWithTime(key, val, 10*time.Second)
}
