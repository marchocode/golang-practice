package mode

import (
	"testing"
	"time"
)

func TestRefresh(t *testing.T) {
	r := NewRefreshAhead()

	r.Write("3", "three")

	time.Sleep(5 * time.Second)
}
