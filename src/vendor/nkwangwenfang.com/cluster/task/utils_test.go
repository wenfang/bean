package task

import (
	"time"

	"golang.org/x/net/context"

	"nkwangwenfang.com/log"
)

type dumpTask struct{}

func (d *dumpTask) Start(c context.Context) (context.CancelFunc, chan struct{}, error) {
	done := make(chan struct{})
	ctx, cancel := context.WithCancel(c)

	go func() {
		defer close(done)

		tick := time.Tick(1e9)
		for {
			select {
			case <-tick:
				log.Info("dumptask run")
			case <-ctx.Done():
				return
			}
		}
	}()
	return cancel, done, nil
}

type dumpOnce struct{}

func (d *dumpOnce) Do(_ context.Context) error {
	log.Info("dumpOnce do")
	return nil
}
