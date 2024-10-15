package cache

import "time"

type Cache interface {
	Get(key string) string

	Set(key, val string)

	SetWithTime(key, val string, time time.Duration)

	Keys() []string

	TTL(key string) time.Duration

	FlushAll()

	Show()
}
