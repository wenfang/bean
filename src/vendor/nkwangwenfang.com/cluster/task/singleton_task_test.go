package task

import (
	"testing"
	"time"

	"golang.org/x/net/context"
)

func TestSingleton(t *testing.T) {
	l := WithSingleton(&dumpTask{})
	cancel, done, err := l.Start(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	_, _, err = l.Start(context.Background())
	if err != ErrIsRunning {
		t.Fatal(err)
	}
	_, _, err = l.Start(context.Background())
	if err != ErrIsRunning {
		t.Fatal(err)
	}
	time.Sleep(5e9)
	cancel()
	<-done
}
