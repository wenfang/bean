package zklock

import (
	"testing"

	"github.com/samuel/go-zookeeper/zk"
)

func TestZkLock(t *testing.T) {
	conn, _, err := zk.Connect([]string{"127.0.0.1:2181"}, 10e9)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	l := New(conn, "/testkeyaa")
	_, err = l.Lock()
	if err != nil {
		t.Fatal(err)
	}

	if _, err = l.Lock(); err != nil {
		t.Fatal(err)
	}
}
