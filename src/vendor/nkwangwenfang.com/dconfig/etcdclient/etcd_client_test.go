package etcdclient

import (
	"testing"

	"golang.org/x/net/context"

	"nkwangwenfang.com/dconfig"
	"nkwangwenfang.com/log"
)

var cfg = Config{
	Servers: []string{"http://127.0.0.1:2379"},
	App:     "test",
}

func TestEtcdClient(t *testing.T) {
	client, err := New(cfg)
	if err != nil {
		t.Fatal(err)
	}

	content, err := dconfig.Load(context.Background(), client, "testkey")
	if err != nil {
		t.Fatal(err)
	}
	log.Info(content)
}
