package task

import (
	"testing"
	"time"

	"golang.org/x/net/context"
)

func TestLoop(t *testing.T) {
	l := WithLoop(&dumpOnce{}, 1e9)
	cancel, done, err := l.Start(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(5e9)
	cancel()
	<-done
}
