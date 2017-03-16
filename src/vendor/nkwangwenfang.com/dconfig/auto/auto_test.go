package auto

import (
	"fmt"
	"testing"
	"time"

	"golang.org/x/net/context"

	"nkwangwenfang.com/cluster/task"
	"nkwangwenfang.com/dconfig/etcdclient"
)

type myObserver struct{}

func (o *myObserver) ConfigUpdate(value string) {
	fmt.Println("config update", value)
}

var cfg = etcdclient.Config{
	Servers: []string{"http://127.0.0.1:2379"},
	App:     "test",
}

func TestAuto(t *testing.T) {
	client, err := etcdclient.New(cfg)
	if err != nil {
		t.Fatal(err)
	}

	ac := New(client)
	ac.Register("testkey", "testvalue", &myObserver{})
	task := task.WithLoop(ac, 3e9)
	cancel, done, err := task.Start(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(10e9)
	cancel()
	<-done
}
