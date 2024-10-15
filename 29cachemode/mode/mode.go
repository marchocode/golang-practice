package mode

type Mode interface {
	
	Read(key string) string

	Write(key, val string)
}
