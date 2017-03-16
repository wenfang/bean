package log

import (
	"testing"
	"time"
)

type dtest int

func (d dtest) String() string {
	return "dtest"
}

type dtests struct {
	Name string
	Age  int
}

func TestLogger(t *testing.T) {
	SetLevel("debug")
	Debug("debug message")
	Warning("can't find warning")
	Fatal("null")
	time.Sleep(1e6)
	SetLevel("debug")
	Warning(20, 45.367, "name", true, "testkey", dtest(20))
	Fatal("dtests", dtests{"p1", 30})
}

func BenchmarkLogger(b *testing.B) {
	SetLogFile("testdata")
	for i := 0; i < b.N; i++ {
		Info("test", "key1", "value1")
	}
}
