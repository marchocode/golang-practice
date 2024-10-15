package db

import "fmt"

type SimpleDb struct {
	data map[string]string
}

func NewSimpleDb() *SimpleDb {

	// simple data
	d := make(map[string]string)

	d["a"] = "1"
	d["b"] = "2"

	return &SimpleDb{
		data: d,
	}
}

func (m *SimpleDb) Select(key string) string {
	return m.data[key]
}

func (m *SimpleDb) Update(key, val string) {
	m.data[key] = val
}

func (m *SimpleDb) SelectAll() {
	fmt.Println(m.data)
}
