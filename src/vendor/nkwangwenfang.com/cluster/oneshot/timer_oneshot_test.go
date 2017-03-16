package oneshot

import (
	"testing"
	"time"

	"golang.org/x/net/context"

	"nkwangwenfang.com/log"
)

type dumpOneshot struct{}

func (d *dumpOneshot) Do(ctx context.Context) error {
	log.Info("dumpTimer Do")
	oneshot := WithTimer(d, 1e9)
	return oneshot.Do(ctx)
}

func TestTimer(t *testing.T) {
	oneshot := WithTimer(&dumpOneshot{}, 2e9)
	oneshot.Do(context.Background())
	time.Sleep(10e9)
}
