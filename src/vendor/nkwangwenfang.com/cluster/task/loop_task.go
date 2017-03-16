package task

import (
	"time"

	"golang.org/x/net/context"

	"nkwangwenfang.com/cluster/oneshot"
	"nkwangwenfang.com/log"
)

type loopTask struct {
	oneshot  oneshot.Oneshot
	interval time.Duration
}

// WithLoop 使Oneshot具有Loop运行能力，成为一个Task
func WithLoop(oneshot oneshot.Oneshot, interval time.Duration) Task {
	return &loopTask{oneshot: oneshot, interval: interval}
}

func (lt *loopTask) Start(c context.Context) (context.CancelFunc, chan struct{}, error) {
	done := make(chan struct{})
	ctx, cancel := context.WithCancel(c)

	go func() {
		defer close(done)

		// 先执行一次
		if err := lt.oneshot.Do(ctx); err != nil {
			log.Error("loopTask oneshot do failed", "error", err)
		}

		tick := time.Tick(lt.interval)
		for {
			select {
			case <-tick:
				if err := lt.oneshot.Do(ctx); err != nil {
					log.Error("loopTask oneshot do failed", "error", err)
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return cancel, done, nil
}
