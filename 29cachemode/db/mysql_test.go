package db

import (
	"testing"
)

func TestMysql(t *testing.T) {

	mysql := NewMysql()
	defer mysql.Close()

	mysql.Update("1", "one")
	mysql.Update("2", "two")

	one := mysql.Select("1")
	t.Logf("one value = %s", one)
}
