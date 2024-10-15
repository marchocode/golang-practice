package db

type Database interface {
	
	Select(key string) string

	Update(key, val string)

	SelectAll()

	Close();
}
