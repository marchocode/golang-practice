package mode

import (
	"testing"
	"time"
)

func TestRead(t *testing.T) {

	m := NewWriteBack()

	m.Write("version", "0.0")
	m.Write("demo", "TestRead")

	// change
	m.Write("version", "1.1")

	v := m.Read("version")
	t.Logf("v = %s", v)

	// wait
	time.Sleep(15 * time.Second)

	// 查询是否入库
	m.database.SelectAll()

}
