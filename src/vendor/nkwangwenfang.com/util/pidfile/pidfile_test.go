package pidfile

import (
	"testing"
)

func TestPid(t *testing.T) {
	filename := "testdata/test.pid"

	if err := Save(filename); err != nil {
		t.Fatal(err)
	}
	Remove(filename)
}
