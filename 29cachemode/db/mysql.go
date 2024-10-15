package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func init() {

}

type Mysql struct {
	db *sql.DB
}

func NewMysql() *Mysql {

	database, err := sql.Open("mysql", "root:root@/test")

	if err != nil {
		panic(err)
	}

	return &Mysql{
		db: database,
	}
}

func (m *Mysql) Select(key string) string {

	qstat, err := m.db.Prepare("Select val from test where k = ?")

	if err != nil {
		panic(err)
	}

	defer qstat.Close()

	if err != nil {
		panic(err)
	}

	var val string
	qstat.QueryRow(key).Scan(&val)

	return val
}

func (m *Mysql) Update(key, val string) {

	stat, err := m.db.Prepare("REPLACE INTO test values (?,?)")

	if err != nil {
		panic(err)
	}

	defer stat.Close()
	stat.Exec(key, val)
}

func (m *Mysql) SelectAll() {
}

func (m *Mysql) Close() {
	m.db.Close()
}
